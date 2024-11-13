package gb32960

import (
	"bcd-util/support_parse/gb32960"
	"bcd-util/support_parse/parse"
	"bcd-util/util"
	"context"
	"embed"
	"encoding/hex"
	"encoding/json"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/pkg/errors"
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

var port int
var period int

func Cmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "gb32960",
		Short: "gb32960模拟器",
		Run: func(cmd *cobra.Command, args []string) {
			start()
		},
	}
	cmd.Flags().IntVarP(&port, "port", "o", 13579, "ws服务监听端口")
	cmd.Flags().IntVarP(&period, "period", "p", 10, "报文上报间隔(秒)")
	return &cmd
}

const sample = "232302FE4C534A4533363039364D53313430343935010141170608100A10010103010040000003520F2827811C012E2000000002010101594FDB4E2F4A0F3227100500073944E501DD620A0601090E1B01370E14010145010444070300021387000000000801010F282781006C00016C0E180E190E1A0E190E190E180E180E1A0E1B0E180E190E1A0E180E180E190E1A0E1A0E190E180E1A0E180E1A0E1A0E180E170E190E170E190E170E190E1B0E190E190E190E180E180E170E170E180E170E170E170E190E170E180E170E190E170E170E170E180E180E190E190E140E180E180E170E170E150E160E160E180E190E170E180E170E180E170E180E170E160E190E150E180E160E180E170E160E160E170E150E170E170E140E170E160E160E170E170E170E170E160E170E160E170E140E170E170E160E160E170E170E170E160E160E160E16090101000C454545444544444445444544F5"

type InMsg struct {
	/**
	1、连接tcp网关
	2、更新运行数据
	3、发送登陆报文
	4、发送登出报文
	*/
	Flag int    `json:"flag"`
	Data string `json:"data"`
}

type OutMsg struct {
	/**
	1、连接tcp网关结果
	2、更新运行数据结果
	3、发送登陆报文结果
	4、发送登出报文结果
	101、同步服务器运行数据到客户端
	102、发送数据到网关成功通知
	103、接收到网关的响应数据
	104、tcp网关断开通知
	*/
	Flag    int    `json:"flag"`
	Data    string `json:"data"`
	Succeed bool   `json:"succeed"`
}

type TcpClient struct {
	conn    net.Conn
	stopCtx context.Context
	stopFn  context.CancelFunc
}

func (e *TcpClient) init(address string, wsClient *WsClient, onConnect func(error)) chan bool {
	resChan := make(chan bool)
	//启动协程接收tcp网关数据
	go func() {
		conn, err := net.Dial("tcp", address)
		if err != nil {
			onConnect(err)
			util.Log.Errorf("%+v", err)
			return
		}
		e.conn = conn
		defer conn.Close()

		stopCtx, stopFn := context.WithCancel(wsClient.stopCtx)
		e.stopCtx = stopCtx
		e.stopFn = stopFn
		defer stopFn()

		onConnect(nil)

		buf := make([]byte, 8192)
		for {
			select {
			case <-stopCtx.Done():
				return
			case <-wsClient.stopCtx.Done():
				return
			default:
				err := conn.SetReadDeadline(time.Now().Add(100 * time.Millisecond))
				if err != nil {
					util.Log.Errorf("%+v", err)
					return
				}
				n, err := conn.Read(buf)
				if err != nil {
					if os.IsTimeout(err) {
						continue
					} else {
						wsClient.send(&OutMsg{
							Flag:    104,
							Data:    "",
							Succeed: true,
						})
						util.Log.Errorf("%+v", err)
						return
					}
				}
				wsClient.onTcpMsg(buf[0:n])
			}
		}
	}()
	return resChan
}

func (e *TcpClient) send(b []byte) {
	_, err := e.conn.Write(b)
	if err != nil {
		e.stopFn()
		util.Log.Errorf("%+v", err)
	}
}

type WsClient struct {
	vin       string
	conn      net.Conn
	tcpClient *TcpClient
	packet    *gb32960.Packet
	stopCtx   context.Context
	stopFn    context.CancelFunc
}

func (e *WsClient) init(vin string, conn net.Conn, onMsg func(msg *InMsg)) error {
	defer conn.Close()
	e.vin = vin
	e.conn = conn

	//定义停止通知ctx
	stopCtx, stopFn := context.WithCancel(context.Background())
	e.stopCtx = stopCtx
	e.stopFn = stopFn
	defer stopFn()

	//初始化样本
	decodeString, err := hex.DecodeString(sample)
	if err != nil {
		return errors.WithStack(err)
	}
	byteBuf := parse.ToByteBuf(decodeString)

	temp := gb32960.To_Packet(byteBuf)
	temp.F_vin = vin
	e.packet = temp

	//更新客户端运行数据
	marshal, err := json.MarshalIndent(e.packet.F_data, "", "   ")
	if err != nil {
		return errors.WithStack(err)
	}
	e.send(&OutMsg{
		Flag:    101,
		Data:    string(marshal),
		Succeed: true,
	})

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

func (e *WsClient) startSendRunData() {
	//启动协程执行循环发送
	go func() {
		var sendTs = time.Now().UnixMilli()
		for {
			waitMills := sendTs - time.Now().UnixMilli()
			if waitMills > 0 {
				select {
				case <-e.tcpClient.stopCtx.Done():
					return
				case <-time.After(time.Duration(waitMills) * time.Millisecond):
					e.sendRunData(sendTs)
				}
			} else {
				select {
				case <-e.tcpClient.stopCtx.Done():
					return
				default:
					e.sendRunData(sendTs)
				}
			}
			sendTs += int64(period) * 1000
		}
	}()
}

//go:embed resource
var FS embed.FS

func start() {
	//gin.SetMode(gin.ReleaseMode)
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

	//engine.Static("/resource", "cmd_simlator/gb32960/resource")

	engine.POST("/parse", func(ctx *gin.Context) {
		res := make(map[string]any)
		all, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			util.Log.Errorf("%+v", err)
			return
		}
		ctx.Header("content-type", "application/json;charset=utf-8")
		decodeString, err := hex.DecodeString(string(all))
		if err != nil {
			util.Log.Errorf("%+v", err)
			res["msg"] = "解析失败、数据不是16进制"
			res["succeed"] = false
			ctx.JSON(200, res)
			return
		}
		buf := parse.ToByteBuf(decodeString)

		var packet *gb32960.Packet
		func() {
			defer func() {
				if err := recover(); err != nil {
					util.Log.Errorf("%+v", err)
					res["msg"] = "解析失败、报文不符合32960协议格式"
					res["succeed"] = false
					ctx.JSON(200, res)
				}
			}()
			packet = gb32960.To_Packet(buf)
		}()

		if packet != nil {
			packetJson, err := json.Marshal(packet)
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

		var p gb32960.Packet
		err = json.Unmarshal(all, &p)
		if err != nil {
			util.Log.Errorf("%+v", err)
			res["msg"] = "反解析失败、数据不是国标报文的json格式、错误原因:\n" + err.Error()
			res["succeed"] = false
			ctx.JSON(200, res)
			return
		}
		var bs []byte
		func() {
			defer func() {
				if err := recover(); err != nil {
					util.Log.Errorf("%+v", err)
					res["msg"] = "解析失败、报文不符合32960协议格式"
					res["succeed"] = false
					ctx.JSON(200, res)
				}
			}()
			buf := parse.ToByteBuf_empty()
			p.Write(buf)
			bs = buf.ToBytes()
		}()

		if bs != nil {
			res["data"] = strings.ToUpper(hex.EncodeToString(bs))
			res["succeed"] = true
			ctx.JSON(200, res)
		}
	})

	engine.GET("/ws", func(ctx *gin.Context) {
		//升级websocket
		wsConn, _, _, err := ws.UpgradeHTTP(ctx.Request, ctx.Writer)
		if err != nil {
			util.Log.Errorf("%+v", err)
			return
		}
		go func() {
			defer wsConn.Close()
			//获取参数
			vin := ctx.Query("vin")

			//定义ws对象
			wsClient := &WsClient{}

			//初始化ws客户端对象
			err = wsClient.init(vin, wsConn, wsClient.onWsMsg)
			if err != nil {
				util.Log.Errorf("%+v", err)
				return
			}

			util.Log.Infof("-----exit ws-----")
		}()
	})
	err := engine.Run(":" + strconv.Itoa(port))
	if err != nil {
		util.Log.Errorf("%+v", err)
	}
}

func (e *WsClient) sendRunData(sendTs int64) {
	vehicleRunData := e.packet.F_data.(*gb32960.VehicleRunData)
	vehicleRunData.F_collectTime = sendTs
	buf := parse.ToByteBuf_empty()

	e.packet.Write(buf)
	b := buf.ToBytes()
	e.tcpClient.send(b)
	util.Log.Infof("sendRunData vin[%s] time[%s] succeed", e.vin, time.UnixMilli(sendTs).Format("20060102150405"))
	e.send(&OutMsg{
		Flag:    102,
		Data:    hex.EncodeToString(b),
		Succeed: true,
	})
}

func (e *WsClient) onTcpMsg(b []byte) {
	e.send(&OutMsg{
		Flag:    103,
		Data:    hex.EncodeToString(b),
		Succeed: true,
	})
}

func (e *WsClient) onWsMsg(msg *InMsg) {
	switch msg.Flag {
	case 1:
		e.HandleConnectTcp(msg.Data)
	case 2:
		e.HandleUpdateRunData(msg.Data)
	case 3:
		e.HandleSendLogin(msg.Data)
	case 4:
		e.HandleSendLogout(msg.Data)
	default:
		util.Log.Warnf("flag[%d] not support", msg.Flag)
	}
}

func (e *WsClient) HandleConnectTcp(data string) {
	//初始化tcp连接
	e.tcpClient = &TcpClient{}
	e.tcpClient.init(data, e, func(err error) {
		if err == nil {
			e.startSendRunData()
			e.send(&OutMsg{
				Flag:    1,
				Data:    "",
				Succeed: true,
			})
		} else {
			e.send(&OutMsg{
				Flag:    1,
				Data:    err.Error(),
				Succeed: false,
			})
		}
	})
}

func (e *WsClient) HandleUpdateRunData(data string) {
	vehicleRunData := gb32960.VehicleRunData{}
	err := json.Unmarshal([]byte(data), &vehicleRunData)
	if err != nil {
		util.Log.Errorf("%+v", err)
		e.send(&OutMsg{
			Flag:    2,
			Data:    "",
			Succeed: false,
		})
		return
	}
	e.packet.F_data = &vehicleRunData

	buf := parse.ToByteBuf_empty()
	e.packet.Write(buf)

	actualLen := uint16(buf.ReadableBytes() - 25)
	exceptLen := e.packet.F_contentLength
	if actualLen != exceptLen {
		e.packet.F_contentLength = actualLen
	}

	e.send(&OutMsg{
		Flag:    2,
		Data:    "",
		Succeed: true,
	})
	util.Log.Infof("HandleUpdateRunData vin[%s]", e.vin)
}

func (e *WsClient) HandleSendLogin(data string) {
	vehicleLoginData := gb32960.VehicleLoginData{}
	err := json.Unmarshal([]byte(data), &vehicleLoginData)
	if err != nil {
		util.Log.Errorf("%+v", err)
		e.send(&OutMsg{
			Flag:    3,
			Data:    "",
			Succeed: false,
		})
		return
	}
	bytes := gb32960.ToPacketBytes(1, 0xFE, e.vin, &vehicleLoginData)
	e.tcpClient.send(bytes)

	e.send(&OutMsg{
		Flag:    3,
		Data:    hex.EncodeToString(bytes),
		Succeed: true,
	})
	util.Log.Infof("HandleSendLogin vin[%s]", e.vin)
}

func (e *WsClient) HandleSendLogout(data string) {
	vehicleLogoutData := gb32960.VehicleLogoutData{}
	err := json.Unmarshal([]byte(data), &vehicleLogoutData)
	if err != nil {
		util.Log.Errorf("%+v", err)
		e.send(&OutMsg{
			Flag:    4,
			Data:    "",
			Succeed: false,
		})
	}
	bytes := gb32960.ToPacketBytes(4, 0xFE, e.vin, &vehicleLogoutData)
	e.tcpClient.send(bytes)

	e.send(&OutMsg{
		Flag:    4,
		Data:    hex.EncodeToString(bytes),
		Succeed: true,
	})
	util.Log.Infof("HandleSendLogout vin[%s]", e.vin)
}

func Main() {
	cobra.MousetrapHelpText = ""
	_ = Cmd().Execute()
}
