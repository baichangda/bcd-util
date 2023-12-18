package immotors

import (
	"bcd-util/support_parse/immotors"
	"bcd-util/support_parse/parse"
	"bcd-util/util"
	"context"
	"embed"
	"encoding/base64"
	"encoding/json"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
	"github.com/spf13/cobra"
	"io/fs"
	"net"
	"net/http"
	"strconv"
	"time"
)

type Json struct {
	FileName    string  `json:"fileName"`
	FileContent string  `json:"fileContent"`
	Timestamp   int64   `json:"timestamp"`
	MessageId   string  `json:"messageId"`
	Ext         JsonExt `json:"ext"`
}

func ToJson(vin string, vehicleType string, ts int64, packets []immotors.Packet) (*Json, error) {
	ts = ts - 9
	dateStr := time.Unix(ts, 0).Format("20060102150405")
	buf_empty := parse.ToByteBuf_empty()
	immotors.Write_Packets(packets, buf_empty)
	toBytes := buf_empty.ToBytes()
	//util.Log.Infof("--------------\n%s", hex.EncodeToString(toBytes))
	r, err := util.Gzip(toBytes)
	if err != nil {
		return nil, err
	}
	return &Json{
		FileName:    vin + "_" + dateStr[0:8] + "_" + dateStr[8:] + "_E_V2.0.6.8.bl.gz",
		FileContent: base64.StdEncoding.EncodeToString(r),
		Timestamp:   ts,
		MessageId:   vin + strconv.FormatInt(ts, 10),
		Ext:         JsonExt{VehicleModel: vehicleType},
	}, nil
}

func (e *Json) ToBytes() ([]byte, error) {
	marshal, err := json.Marshal(e)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return marshal, nil
}

type JsonExt struct {
	VehicleModel string `json:"vehicleModel"`
}

var kafkaAddress []string
var topic string

func Cmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "immotors",
		Short: "智己模拟器",
		Run: func(cmd *cobra.Command, args []string) {
			start()
		},
	}
	cmd.Flags().StringSliceVarP(&kafkaAddress, "kafkaAddress", "a", []string{"10.0.11.50:39003"}, "kafka地址")
	cmd.Flags().StringVarP(&topic, "topic", "t", "gw-test", "发送的topic")
	return &cmd
}

const base64Str = "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"

type InMsg struct {
	/**
	1、更新运行数据
	*/
	Flag int    `json:"flag"`
	Data string `json:"data"`
}

type OutMsg struct {
	/**
	1、更新运行数据结果
	101、同步服务器运行到客户端
	102、发送数据到kafka成功通知
	*/
	Flag    int    `json:"flag"`
	Data    string `json:"data"`
	Succeed bool   `json:"succeed"`
}

type WsClient struct {
	vin         string
	vehicleType string
	packet      *immotors.Packet
	conn        net.Conn
}

func (e *WsClient) init(vin string, vehicleType string, conn net.Conn) error {
	e.vin = vin
	e.vehicleType = vehicleType
	e.conn = conn

	//初始化样本
	decodeString, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return errors.WithStack(err)
	}
	unGzip, err := util.UnGzip(decodeString)
	if err != nil {
		return err
	}
	byteBuf := parse.ToByteBuf(unGzip)

	temp := immotors.To_Packets(byteBuf)
	e.packet = &temp[0]
	e.packet.F_evt_D00A.F_VIN = vin
	return nil
}

func (e *WsClient) HandleUpdatePacket(cancelCtx context.Context, data string) {
	packet := immotors.Packet{}
	err := json.Unmarshal([]byte(data), &packet)
	if err != nil {
		util.Log.Errorf("%+v", err)
		err = e.response(cancelCtx, OutMsg{
			Flag:    1,
			Data:    "",
			Succeed: false,
		})
		if err != nil {
			util.Log.Errorf("%+v", err)
		}
		return
	}
	e.packet = &packet
	err = e.response(cancelCtx, OutMsg{
		Flag:    1,
		Data:    "",
		Succeed: true,
	})
	if err != nil {
		util.Log.Errorf("%+v", err)
	}
	util.Log.Infof("HandleUpdatePacket vin[%s] vehicleType[%s]", e.vin, e.vehicleType)
}

//go:embed resource
var FS embed.FS

func start() {
	//连接kafka
	kafkaWriter := &kafka.Writer{
		Addr:         kafka.TCP(kafkaAddress...),
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		BatchTimeout: 100 * time.Millisecond,
		BatchSize:    1000,
		//Async:                  true,
		AllowAutoTopicCreation: true,
	}
	defer kafkaWriter.Close()

	engine := gin.Default()
	engine.Use(gzip.Gzip(gzip.DefaultCompression))

	sub, err2 := fs.Sub(FS, "resource")
	if err2 != nil {
		util.Log.Errorf("%+v", err2)
		return
	}
	engine.StaticFS("/immotors/resource", http.FS(sub))

	//engine.Static("/immotors/resource", "cmd_simlator/immotors/resource")

	engine.GET("/immotors/ws", func(ctx *gin.Context) {
		//升级websocket
		conn, _, _, err := ws.UpgradeHTTP(ctx.Request, ctx.Writer)
		if err != nil {
			util.Log.Errorf("%+v", err)
			return
		}
		go func() {
			defer conn.Close()
			//获取参数
			vin := ctx.Query("vin")
			vehicleType := ctx.Query("vehicleType")
			//创建客户端
			client := &WsClient{}
			err = client.init(vin, vehicleType, conn)
			if err != nil {
				util.Log.Errorf("%+v", err)
				return
			}

			//定义ctx
			cancelCtx, cancelFn := context.WithCancel(context.Background())
			defer cancelFn()

			//更新客户端运行数据
			marshal1, err := json.MarshalIndent(client.packet, "", "   ")
			if err != nil {
				util.Log.Errorf("%+v", err)
				return
			}
			err = client.response(cancelCtx, OutMsg{
				Flag:    101,
				Data:    string(marshal1),
				Succeed: true,
			})
			if err != nil {
				util.Log.Errorf("%+v", err)
				return
			}

			//启动协程执行循环发送
			go func() {
				prevTs := time.Now().UnixMilli()
				for {
					select {
					case <-cancelCtx.Done():
						return
					default:
						err := client.send(cancelCtx, kafkaWriter)
						if err != nil {
							util.Log.Errorf("%+v", err)
							return
						}
					}
					diff := 10000 + prevTs - time.Now().UnixMilli()

					if diff > 0 {
						time.Sleep(time.Duration(diff) * time.Millisecond)
					}

					//sleep后设置下次发送时间
					prevTs = prevTs + 10000
				}
			}()

			for {
				text, err := wsutil.ReadClientText(conn)
				if err != nil {
					util.Log.Errorf("%+v", err)
					return
				}
				inMsg := InMsg{}
				err = json.Unmarshal(text, &inMsg)
				if err != nil {
					util.Log.Errorf("%+v", err)
					return
				}
				switch inMsg.Flag {
				case 1:
					client.HandleUpdatePacket(cancelCtx, inMsg.Data)
				default:
					util.Log.Warnf("flag[%d] not support", inMsg.Flag)
				}
			}
		}()

	})
	err := engine.Run(":13579")
	if err != nil {
		util.Log.Errorf("%+v", err)
	}
}

func (e *WsClient) response(cancelCtx context.Context, msg OutMsg) error {
	marshal, err := json.Marshal(msg)
	if err != nil {
		return errors.WithStack(err)
	}
	err = wsutil.WriteServerText(e.conn, marshal)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (e *WsClient) send(cancelCtx context.Context, w *kafka.Writer) error {
	now := time.Now()
	ts := now.UnixMilli() / 1000
	temp := e.packet
	packets := make([]immotors.Packet, 10)
	for i := 0; i < 10; i++ {
		cur := *temp
		cur.F_evt_0001 = &immotors.Evt_0001{
			F_evtId:      0x0001,
			F_TBOXSysTim: ts - int64(10-i),
		}
		packets[i] = cur
	}

	dateStr := now.Format("20060102150405")

	toJson, err := ToJson(e.vin, e.vehicleType, ts, packets)
	if err != nil {
		return err
	}
	toBytes, err := toJson.ToBytes()
	if err != nil {
		return err
	}
	msg := kafka.Message{
		Key:   []byte(toJson.MessageId),
		Value: toBytes,
	}

	err = w.WriteMessages(cancelCtx, msg)
	if err != nil {
		return errors.WithStack(err)
	}

	util.Log.Infof("send vin[%s] vehicleType[%s] time[%s] topic[%s] succeed", e.vin, e.vehicleType, dateStr, topic)

	err = e.response(cancelCtx, OutMsg{
		Flag:    102,
		Data:    "",
		Succeed: true,
	})
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
