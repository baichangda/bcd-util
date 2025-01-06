package gb32960

import (
	"bcd-util/support_parse/gb32960"
	"bcd-util/support_parse/parse"
	"bcd-util/util"
	"context"
	"encoding/hex"
	"net"
	"os"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/spf13/cobra"
)

var sample string

var period int
var startIndex int
var num int
var address string
var filePath string

func Cmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "gb32960",
		Short: "gb32960压力测试、车辆vin以TEST000000开头、后面按照顺序生成序号",
		Run: func(cmd *cobra.Command, args []string) {
			Start()
		},
	}
	cmd.Flags().IntVarP(&period, "period", "p", 10, "报文上报间隔(秒)")
	cmd.Flags().IntVarP(&startIndex, "startIndex", "s", 0, "车辆开始索引(从0开始)")
	cmd.Flags().IntVarP(&num, "num", "n", 1, "压测车辆数")
	cmd.Flags().StringVarP(&address, "address", "a", "127.0.0.1:6666", "网关tcp地址")
	cmd.Flags().StringVarP(&filePath, "filePath", "f", "sample.txt", "存储样例报文的路径(必须且仅存储一条报文在文件中)")
	_ = cmd.MarkFlagRequired("address")
	return &cmd
}

var clientNum uint32 = 0
var sendNum uint32 = 0

func getVins() []string {
	vinPrefix := "TEST000000"
	vins := make([]string, num)
	for i := startIndex; i < startIndex+num; i++ {
		itoa := strconv.Itoa(i)
		vins[i] = vinPrefix + strings.Repeat("0", 7-len(itoa)) + itoa
	}
	return vins
}

func Start() {
	//加载样例数据
	file, err := os.ReadFile(filePath)
	if err != nil {
		util.Log.Errorf("%+v", err)
		return
	}
	sample = strings.TrimSpace(string(file))
	if sample == "" {
		util.Log.Errorf("sample file[%s] content is empty", filePath)
		return
	}
	util.Log.Infof("load sample:\n%s", sample)

	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-time.After(3 * time.Second):
			}
			util.Log.Infof("client[%d] sendSpeed[%d/s]", atomic.LoadUint32(&clientNum), atomic.SwapUint32(&sendNum, 0)/3)
		}
	}()
	vins := getVins()
	if num < period {
		for _, e := range vins {
			go startClient(ctx, e)
			time.Sleep(1 * time.Second)
		}
	} else {
		//分组批次
		batchNum := num / period
		for i := 0; i < period; i++ {
			if i == period-1 {
				for _, e := range vins[i*batchNum:] {
					go startClient(ctx, e)
				}
			} else {
				for _, e := range vins[i*batchNum : (i+1)*batchNum] {
					go startClient(ctx, e)
				}
			}
			time.Sleep(1 * time.Second)
		}
	}

	for {
		time.Sleep(time.Hour)
	}
}

func startClient(ctx context.Context, vin string) {
	//初始化报文
	decodeString, err := hex.DecodeString(sample)
	if err != nil {
		util.Log.Errorf("%+v", err)
		return
	}
	packet := gb32960.To_Packet(parse.ToByteBuf(decodeString))
	packet.F_vin = vin

	dial, err := net.Dial("tcp", address)
	if err != nil {
		util.Log.Errorf("%+v", err)
		os.Exit(0)
		return
	}
	defer dial.Close()
	atomic.AddUint32(&clientNum, 1)

	//启动一个协程一直读数据
	go func() {
		for {
			buf := make([]byte, 1024)
			_, err := dial.Read(buf)
			if err != nil {
				util.Log.Errorf("%+v", err)
				os.Exit(0)
				return
			}
		}
	}()

	var sendTs = time.Now().UnixMilli()
	var sendBuf = parse.ToByteBuf_empty()
A:
	for {
		waitMills := sendTs - time.Now().UnixMilli()
		if waitMills > 0 {
			select {
			case <-ctx.Done():
				break A
			case <-time.After(time.Duration(waitMills) * time.Millisecond):
			}
		} else {
			select {
			case <-ctx.Done():
				break A
			default:
			}
		}
		doBeforeSend(packet, sendTs)
		sendBuf.Clear()
		packet.Write(sendBuf)
		atomic.AddUint32(&sendNum, 1)
		_, err := dial.Write(sendBuf.ToBytes())
		if err != nil {
			util.Log.Errorf("%+v", err)
			return
		}
		sendTs += int64(period) * 1000
	}
}

func doBeforeSend(packet *gb32960.Packet, ts int64) {
	vehicleRunData := packet.F_data.(*gb32960.VehicleRunData)
	vehicleRunData.F_collectTime = ts
}
