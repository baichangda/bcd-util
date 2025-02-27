package immotors_bin

import (
	"bcd-util/support_parse/immotors"
	"bcd-util/support_parse/parse"
	"bcd-util/util"
	"context"
	"embed"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
	"github.com/spf13/cobra"
	"io"
	"io/fs"
	"net"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var kafkaAddress []string
var topic string
var port int
var period int

func Cmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "immotors_bin",
		Short: "智己模拟器(二进制)",
		Run: func(cmd *cobra.Command, args []string) {
			start()
		},
	}
	cmd.Flags().StringSliceVarP(&kafkaAddress, "kafkaAddress", "a", []string{"10.0.11.50:39003"}, "kafka地址")
	cmd.Flags().StringVarP(&topic, "topic", "t", "gw-test", "发送的topic")
	cmd.Flags().IntVarP(&port, "port", "p", 13579, "http端口")
	cmd.Flags().IntVarP(&period, "period", "r", 10, "上报频率")
	return &cmd
}

const hexStr = "000100006466E9B3000300006583EC6E000413C339C87BDE00051DD636E16BEC00060000050000070007F024FFF8FFBC000801CC0001000000094300C4100033000AB70400620000000B000000000000000C0017AE000000000D000992400000000E082C00000000000F5208ABE0000000104E934FE20000001100000000000008008169000000000801000000000000080200000000000008030250AB000000D006003E8D3A8987B4A0365D34F9A820FFFC0080180100060699FC9700008000A0001FF90050010FF915B2001B48006D011F016F26B7E0C95F722400B4F800000000D008004600000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000D00900120000000000000000000FE000010005648000D00A00554C534A4533363039364D533134303439354C53323143303141353239363030303638393836303932313735303030383935323034333839303836303334323032323030303031303231303030303333353932343437D00B00D96C7DF87DC87DF87DE87DF87DD87DF87DE87E007DE07E007DD87DD07DD87DD87DD87DE87DC87DD87DE07DD87DE87DE07DF07DD07DF87DF07E007DE87DF07DE07E087DF07DE87DD87DF87DF07DC87DF07DE07DF07DE07DC87DC87DF07DE87DF87DE87DF07DE07E007DE07DF87DE87DE87DF87DE07DF07DF07DF07DE87E007DF07DE87DE87DF87DF87E007E007E007DF87E087DF07DE07DF87DF07DF87DE07DF87DE87E087DF07DF07DD87DF87DF07DF87DE87E087DE07DE07DF87DE07E007DF87DF07DE87DF07DF87DF87DE87E087DF87DF87DF07E007DE07E08D00C00190C3E003C003C003C003E003C003C003C003D003C003C003C00D00D00190C3F003D003D003E003E003D003D003E003E003C003D003E00D00E00191830354C50454A334333353234304142434230313032303135D00F000C007D00007831800019007E00D01000380000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000D01100230000000000000000000000000000000000000000000000000000000000000000000000D012003F000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000D013004D0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000D01400230000000000000000000000000000000000000000000000000000000000000000000000D015000700000000000000D016003100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000D017002A000000000000000000000000000000000000000000000000000000000000000000000000000000000000D0180014CF639C87BB0EEB1B720000323374D920FA040300D0190010CE10053DB55580000000000000000000D01A002D000000000000000000000000000000000031323334353637383930313233343536373839303132333435363738D01B003F000000313233343536373839303132333435363738000031323334353637383930313233343536373800003132333435363738393031323334353637380000D01C002E31323334353637383930313233343536373831323334353637383930313233343536370000000000000000000000D01D000C000000000000000000000000D01F001000000000000000000000000000000000D020000E0000000000000000000000000000"

type InMsg struct {
	/**
	1、更新运行数据
	2、开始发送运行数据
	*/
	Flag int    `json:"flag"`
	Data string `json:"data"`
}

type OutMsg struct {
	/**
	1、更新运行数据结果
	2、开始发送运行数据
	101、同步服务器运行到客户端
	102、发送数据到kafka成功通知
	*/
	Flag    int    `json:"flag"`
	Data    string `json:"data"`
	Succeed bool   `json:"succeed"`
}

type WsClient struct {
	kafkaWriter  *kafka.Writer
	vin          string
	vehicleType  string
	packet       *immotors.Packet
	conn         net.Conn
	cancelCtx    context.Context
	cancelFn     context.CancelFunc
	lock         sync.Mutex
	startSendCtx context.Context
	startSendFn  context.CancelFunc
}

func (e *WsClient) init(vin string, vehicleType string, conn net.Conn) error {
	e.vin = vin
	e.vehicleType = vehicleType
	e.conn = conn

	//初始化样本
	bs, err := hex.DecodeString(hexStr)
	if err != nil {
		return errors.WithStack(err)
	}
	byteBuf := parse.ToByteBuf(bs)

	packet := immotors.To_Packet(byteBuf)
	packet.F_evt_D00A.F_VIN = vin
	e.packet = packet

	//更新客户端运行数据
	marshal1, err := json.MarshalIndent(packet, "", "   ")
	if err != nil {
		return errors.WithStack(err)
	}
	err = e.response(OutMsg{
		Flag:    101,
		Data:    string(marshal1),
		Succeed: true,
	})
	if err != nil {
		return err
	}
	return nil
}

func (e *WsClient) HandleUpdatePacket(data string) {
	packet := immotors.Packet{}
	err := json.Unmarshal([]byte(data), &packet)
	if err != nil {
		util.Log.Errorf("%+v", err)
		err = e.response(OutMsg{
			Flag:    1,
			Data:    "",
			Succeed: false,
		})
		if err != nil {
			util.Log.Errorf("%+v", err)
		}
		return
	}
	if packet.F_evt_D00A != nil {
		packet.F_evt_D00A.F_VIN = e.vin
	}

	//校验是否能解析
	ok := true
	func() {
		defer func() {
			if err := recover(); err != nil {
				util.Log.Errorf("%+v", err)
				util.Log.Infof("%s", data)
				ok = false
			}
		}()
		empty := parse.ToByteBuf_empty()
		packet.Write(empty)
	}()
	if ok {
		e.packet = &packet
		err = e.response(OutMsg{
			Flag:    1,
			Data:    "",
			Succeed: true,
		})
		if err != nil {
			util.Log.Errorf("%+v", err)
		}
		util.Log.Infof("HandleUpdatePacket vin[%s] vehicleType[%s]", e.vin, e.vehicleType)
	} else {
		err = e.response(OutMsg{
			Flag:    1,
			Data:    "json数据错误、无法反解析为二进制",
			Succeed: false,
		})
		if err != nil {
			util.Log.Errorf("%+v", err)
		}
	}
}

func (e *WsClient) HandleStartSend(data string) {
	e.lock.Lock()
	defer e.lock.Unlock()
	if data == "1" {
		if e.startSendFn != nil {
			e.startSendFn()
			e.startSendCtx = nil
			e.startSendFn = nil
		}
		ctx, cancelFunc := context.WithCancel(e.cancelCtx)
		e.startSendCtx = ctx
		e.startSendFn = cancelFunc
		//启动协程执行循环发送
		go func() {
			nextTs := time.Now().UnixMilli()
			for {
				diff := nextTs - time.Now().UnixMilli()
				if diff <= 0 {
					select {
					case <-e.startSendCtx.Done():
						return
					default:
					}
				} else {
					select {
					case <-e.startSendCtx.Done():
						return
					case <-time.After(time.Duration(diff) * time.Millisecond):
					}
				}
				err := e.send(nextTs)
				if err != nil {
					util.Log.Errorf("%+v", err)
					return
				}
				//sleep后设置下次发送时间
				nextTs = nextTs + int64(period*1000)
			}
		}()
		err := e.response(OutMsg{
			Flag:    2,
			Data:    "1",
			Succeed: true,
		})
		if err != nil {
			util.Log.Errorf("%+v", err)
		}
	} else {
		if e.startSendFn != nil {
			e.startSendFn()
			e.startSendCtx = nil
			e.startSendFn = nil
		}
		err := e.response(OutMsg{
			Flag:    2,
			Data:    "2",
			Succeed: true,
		})
		if err != nil {
			util.Log.Errorf("%+v", err)
		}
	}

}

//go:embed resource
var FS embed.FS

func start() {
	//gin.SetMode(gin.ReleaseMode)
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

	engine.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/resource/index.html")
	})

	sub, err2 := fs.Sub(FS, "resource")
	if err2 != nil {
		util.Log.Errorf("%+v", err2)
		return
	}
	engine.StaticFS("/resource", http.FS(sub))

	//engine.Static("/resource", "cmd_simlator/immotors_bin/resource")

	engine.POST("/parse", func(ctx *gin.Context) {
		res := make(map[string]any)
		all, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			util.Log.Errorf("%+v", err)
			return
		}
		ctx.Header("content-type", "application/json;charset=utf-8")

		bytes, err := base64.StdEncoding.DecodeString(string(all))
		if err != nil {
			util.Log.Errorf("%+v", err)
			res["msg"] = "解析失败、数据不是base64格式"
			res["succeed"] = false
			ctx.JSON(200, res)
			return
		}

		unGzip, err := util.UnGzip(bytes)
		if err != nil {
			util.Log.Errorf("%+v", err)
			res["msg"] = "解析失败、数据不是gzip格式"
			res["succeed"] = false
			ctx.JSON(200, res)
			return
		}

		buf := parse.ToByteBuf(unGzip)

		var packets []immotors.Packet
		func() {
			defer func() {
				if err := recover(); err != nil {
					util.Log.Errorf("%+v", err)
					res["msg"] = "解析失败、报文不符合智己协议格式"
					res["succeed"] = false
					ctx.JSON(200, res)
				}
			}()
			packets = immotors.To_Packets(buf)
		}()

		if packets != nil {
			packetJson, err := json.Marshal(packets)
			if err != nil {
				util.Log.Errorf("%+v", err)
				return
			}
			res["data"] = string(packetJson)
			res["succeed"] = true
			ctx.JSON(200, res)
		}
	})

	engine.POST("/deParse", func(ctx *gin.Context) {
		res := make(map[string]any)
		all, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			util.Log.Errorf("%+v", err)
			return
		}
		ctx.Header("content-type", "application/json;charset=utf-8")

		var ps []immotors.Packet
		err = json.Unmarshal(all, &ps)
		if err != nil {
			util.Log.Errorf("%+v", err)
			res["msg"] = "反解析失败、反序列化失败、原因:\n" + err.Error()
			res["succeed"] = false
			ctx.JSON(200, res)
			return
		}

		buf := parse.ToByteBuf_empty()
		deParseSucceed := true
		func() {
			defer func() {
				if err := recover(); err != nil {
					util.Log.Errorf("%+v", err)
					res["msg"] = "反解析失败、报文不符合智己协议格式"
					res["succeed"] = false
					deParseSucceed = true
					ctx.JSON(200, res)
				}
			}()
			immotors.Write_Packets(ps, buf)
		}()

		if deParseSucceed {
			bytes, err := util.Gzip(buf.ToBytes())
			if err != nil {
				util.Log.Errorf("%+v", err)
				res["msg"] = "反解析失败、数据gzip失败"
				res["succeed"] = false
				ctx.JSON(200, res)
				return
			}

			res["data"] = base64.StdEncoding.EncodeToString(bytes)
			res["succeed"] = true
			ctx.JSON(200, res)
		}
	})

	engine.GET("/ws", func(ctx *gin.Context) {
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
			//定义ctx
			cancelCtx, cancelFn := context.WithCancel(context.Background())
			defer cancelFn()
			//创建客户端
			client := &WsClient{
				cancelCtx:   cancelCtx,
				cancelFn:    cancelFn,
				lock:        sync.Mutex{},
				kafkaWriter: kafkaWriter,
			}

			err = client.init(vin, vehicleType, conn)
			if err != nil {
				util.Log.Errorf("%+v", err)
				return
			}

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
					client.HandleUpdatePacket(inMsg.Data)
				case 2:
					client.HandleStartSend(inMsg.Data)
				default:
					util.Log.Warnf("flag[%d] not support", inMsg.Flag)
				}
			}
		}()

	})
	err := engine.Run(":" + strconv.Itoa(port))
	if err != nil {
		util.Log.Errorf("%+v", err)
	}
}

func (e *WsClient) response(msg OutMsg) error {
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

func (e *WsClient) send(ts int64) error {
	temp := e.packet
	tss := ts / 1000
	packets := make([]immotors.Packet, period)
	for i := 0; i < period; i++ {
		cur := *temp
		cur.F_evt_0001 = &immotors.Evt_0001{
			F_evtId:      0x0001,
			F_TBOXSysTim: tss - int64(period-i),
		}
		packets[i] = cur
	}
	packets[period-1].F_evt_FFFF = &immotors.Evt_FFFF{
		F_evtId: 0xFFF,
		F_CRC32: 0,
	}

	toBin, err := immotors.ToBin(e.vin, e.vehicleType, tss, packets)
	if err != nil {
		return err
	}
	toBytes, err := toBin.ToBytes()
	if err != nil {
		return err
	}
	msg := kafka.Message{
		Key:   []byte(toBin.MessageId),
		Value: toBytes,
	}

	err = e.kafkaWriter.WriteMessages(e.cancelCtx, msg)
	if err != nil {
		return errors.WithStack(err)
	}

	dateStr := time.UnixMilli(ts).Format("20060102150405")
	util.Log.Infof("send vin[%s] vehicleType[%s] time[%s] topic[%s] succeed", e.vin, e.vehicleType, dateStr, topic)

	err = e.response(OutMsg{
		Flag:    102,
		Data:    "",
		Succeed: true,
	})
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
