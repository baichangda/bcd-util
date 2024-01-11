package web

import (
	"bcd-util/util"
	"context"
	"embed"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
	"github.com/spf13/cobra"
	"io"
	"io/fs"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Param struct {
	KafkaAddrs []string `json:"kafkaAddrs"`
	KafkaTopic string   `json:"kafkaTopic"`
	Data       string   `json:"data"`
}

var port int

//go:embed resource
var FS embed.FS

func Cmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "web",
		Short: "启动web服务器、可以发送数据到kafka",
		Run: func(cmd *cobra.Command, args []string) {
			gin.SetMode(gin.ReleaseMode)
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

				err = kafkaWriter.WriteMessages(context.Background(), kafka.Message{
					Value: []byte(param.Data),
				})

				if err != nil {
					util.Log.Errorf("%+v", err)
					res["msg"] = "发送数据到kafka失败、错误原因:[" + err.Error() + "]"
					res["succeed"] = false
					ctx.JSON(200, res)
					return
				}

				res["msg"] = "发送成功"
				res["succeed"] = true
				ctx.JSON(200, res)

				util.Log.Infof("send to kafka[%s] topic[%s] succeed", strings.Join(param.KafkaAddrs, ","), param.KafkaTopic)
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
