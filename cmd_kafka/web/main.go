package web

import (
	"bcd-util/util"
	"context"
	"embed"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/segmentio/kafka-go"
	"github.com/spf13/cobra"
	"io"
	"io/fs"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type Param struct {
	KafkaAddrs   []string `json:"kafkaAddrs"`
	KafkaTopic   string   `json:"kafkaTopic"`
	Data         string   `json:"data"`
	MsgSplit     string   `json:"msgSplit"`
	MsgSplitType string   `json:"msgSplitType"`
}

var port int

//go:embed resource
var FS embed.FS

func Cmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "web",
		Short: "启动web服务器、可以发送数据到kafka",
		Run: func(cmd *cobra.Command, args []string) {
			//gin.SetMode(gin.ReleaseMode)
			engine := gin.New()

			sub, err2 := fs.Sub(FS, "resource")
			if err2 != nil {
				util.Log.Errorf("%+v", err2)
				return
			}
			engine.StaticFS("/resource", http.FS(sub))

			//engine.Static("/resource", "cmd_kafka/web/resource")

			engine.GET("/", func(c *gin.Context) {
				c.Redirect(http.StatusMovedPermanently, "/resource/producer.html")
			})

			engine.GET("/consumer", func(ctx *gin.Context) {
				//升级websocket
				wsConn, _, _, err := ws.UpgradeHTTP(ctx.Request, ctx.Writer)
				if err != nil {
					util.Log.Errorf("%+v", err)
					return
				}
				go func() {
					defer wsConn.Close()
					//获取参数
					wsClient := WsClient{}

					var kafkaAddrs = strings.Split(ctx.Query("kafkaAddrs"), ",")
					var kafkaTopic = ctx.Query("kafkaTopic")
					var kafkaGroupId = ctx.Query("kafkaGroupId")

					//初始化ws客户端对象
					err = wsClient.init(kafkaAddrs, kafkaTopic, kafkaGroupId, wsConn, wsClient.onWsMsg)
					if err != nil {
						util.Log.Errorf("%+v", err)
						return
					}

					util.Log.Infof("-----exit ws-----")
				}()
			})

			engine.POST("/producer", func(ctx *gin.Context) {
				res := make(map[string]any)
				all, err := io.ReadAll(ctx.Request.Body)
				if err != nil {
					util.Log.Errorf("%+v", err)
					return
				}
				ctx.Header("content-type", "application/json;charset=utf-8")

				param := Param{}
				err = json.Unmarshal(all, &param)
				util.Log.Infof("param: %+v", param)

				if err != nil {
					util.Log.Errorf("%+v", err)
					res["msg"] = "反序列化参数失败"
					res["succeed"] = false
					ctx.JSON(200, res)
					return
				}

				//连接kafka
				kafkaWriter := &kafka.Writer{
					Addr:         kafka.TCP(param.KafkaAddrs...),
					Topic:        param.KafkaTopic,
					Balancer:     &kafka.LeastBytes{},
					BatchTimeout: 10 * time.Millisecond,
					BatchSize:    1,
					//Async:                  true,
					AllowAutoTopicCreation: true,
				}

				var messages []kafka.Message
				switch param.MsgSplitType {
				case "":
					messages = []kafka.Message{{Value: []byte(param.Data)}}
				case "1":

				}
				if param.MsgSplitType == "" {
					messages = []kafka.Message{{Value: []byte(param.Data)}}
				} else {
					var split []string
					if param.MsgSplitType == "1" {
						split = strings.Split(param.Data, "\n")
					} else {
						split = strings.Split(param.Data, param.MsgSplit)
					}
					messages = make([]kafka.Message, len(split))
					for i := range len(split) {
						messages[i] = kafka.Message{
							Value: []byte(split[i]),
						}
					}
				}

				err = kafkaWriter.WriteMessages(context.Background(), messages...)

				if err != nil {
					util.Log.Errorf("%+v", err)
					res["msg"] = "发送数据到kafka失败、错误原因:[" + err.Error() + "]"
					res["succeed"] = false
					ctx.JSON(200, res)
					return
				}

				res["msg"] = "发送" + strconv.Itoa(len(messages)) + "条数据到kafka成功"
				res["succeed"] = true
				ctx.JSON(200, res)

				util.Log.Infof("send to kafka[%s] topic[%s] num[%d] succeed", strings.Join(param.KafkaAddrs, ","), param.KafkaTopic, len(messages))
			})

			err := engine.Run(":" + strconv.Itoa(port))
			if err != nil {
				util.Log.Errorf("%+v", err)
			}
		},
	}
	cmd.Flags().IntVarP(&port, "port", "p", 23456, "web服务端口")
	return &cmd
}

type WsClient struct {
	conn    net.Conn
	stopCtx context.Context
	stopFn  context.CancelFunc
}

func (e *WsClient) init(kafkaAddrs []string, kafkaTopic string, kafkaGroupId string, conn net.Conn, onMsg func(msg *InMsg)) error {
	stopCtx, stopFn := context.WithCancel(context.Background())
	e.conn = conn
	e.stopCtx = stopCtx
	e.stopFn = stopFn

	go func() {
		reader := kafka.NewReader(kafka.ReaderConfig{
			Brokers:     kafkaAddrs,
			GroupID:     kafkaGroupId,
			Topic:       kafkaTopic,
			MaxAttempts: 1,
		})
		defer reader.Close()

		util.Log.Infof("init kafkaAddrs[%s] kafkaTopic[%s] kafkaGroupId[%s]", strings.Join(kafkaAddrs, ","), kafkaTopic, kafkaGroupId)

		for {
			message, err := reader.ReadMessage(e.stopCtx)
			if err != nil {
				if errors.Is(err, context.Canceled) {
					return
				}
				util.Log.Errorf("%+v", err)
				e.send(&OutMsg{
					Flag:    102,
					Data:    err.Error(),
					Succeed: false,
				})
				e.stopFn()
				return
			}
			e.send(&OutMsg{
				Flag:    101,
				Data:    string(message.Value),
				Succeed: true,
			})
		}
	}()

	//启动协程监听ws信息
	for {
		select {
		case <-stopCtx.Done():
			return nil
		default:
			err := conn.SetReadDeadline(time.Now().Add(100 * time.Millisecond))
			if err != nil {
				util.Log.Errorf("%+v", err)
				return nil
			}
			text, err := wsutil.ReadClientText(conn)
			if err != nil {
				if os.IsTimeout(err) {
					continue
				} else {
					util.Log.Errorf("%+v", err)
					return nil
				}
			}
			inMsg := InMsg{}
			err = json.Unmarshal(text, &inMsg)
			if err != nil {
				util.Log.Errorf("%+v", err)
				continue
			}
			onMsg(&inMsg)
		}
	}

}

func (e *WsClient) send(msg *OutMsg) {
	marshal, err := json.Marshal(msg)
	if err != nil {
		e.stopFn()
		util.Log.Errorf("%+v", err)
	}
	err = wsutil.WriteServerText(e.conn, marshal)
	if err != nil {
		e.stopFn()
		util.Log.Errorf("%+v", err)
	}
}

type InMsg struct {
	/**
	1、连接kafka开始消费
	*/
	Flag int    `json:"flag"`
	Data string `json:"data"`
}

type OutMsg struct {
	/**
	1、连接kafka开始消费
	101、接收到kafka数据
	102、kafka断开通知
	*/
	Flag    int    `json:"flag"`
	Data    string `json:"data"`
	Succeed bool   `json:"succeed"`
}

type KafkaConfig struct {
	kafkaAddrs   []string
	kafkaTopic   string
	kafkaGroupId string
}

func (e *WsClient) onWsMsg(msg *InMsg) {
	switch msg.Flag {
	default:
		util.Log.Warnf("flag[%d] not support", msg.Flag)
	}
}

func (e *WsClient) HandleConsume(data string) {
	kafkaConfig := KafkaConfig{}
	err := json.Unmarshal([]byte(data), &kafkaConfig)
	if err != nil {
		util.Log.Errorf("%+v", e)
		e.stopFn()
		return
	}

}
