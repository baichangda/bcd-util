package immotors_bin

import (
	"bcd-util/support_parse/immotors"
	"bcd-util/support_parse/parse"
	"bcd-util/util"
	"context"
	"encoding/base64"
	"github.com/segmentio/kafka-go"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

var sample []byte

var period int
var startIndex int
var num int
var kafkaAddress []string
var topic string
var filePath string

func Cmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "immotors_bin",
		Short: "智己飞凡压测模拟器(二进制)、车辆vin以TEST000000开头、后面按照顺序生成序号",
		Run: func(cmd *cobra.Command, args []string) {
			Start()
		},
	}
	cmd.Flags().IntVarP(&period, "period", "p", 10, "报文上报间隔(秒)")
	cmd.Flags().IntVarP(&startIndex, "startIndex", "s", 0, "车辆开始索引(从0开始)")
	cmd.Flags().IntVarP(&num, "num", "n", 1, "压测车辆数")
	cmd.Flags().StringSliceVarP(&kafkaAddress, "kafkaAddress", "a", []string{"127.0.0.1:9092"}, "kafka地址")
	cmd.Flags().StringVarP(&topic, "topic", "t", "gw-test", "kafka topic")
	cmd.Flags().StringVarP(&filePath, "filePath", "f", "sample.txt", "存储样例报文的路径(必须且仅存储一条报文在文件中、格式为包含10条的原始报文base64的)")
	_ = cmd.MarkFlagRequired("address")
	return &cmd
}

var clientNum uint32 = 0
var sendNum uint32 = 0

func GetVins(num int, startIndex int) []string {
	vinPrefix := "TEST000000"
	vins := make([]string, num)
	for i := 0; i < num; i++ {
		no := i + startIndex
		itoa := strconv.Itoa(no)
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
	content := strings.TrimSpace(string(file))
	if content == "" {
		util.Log.Errorf("sample file[%s] content is empty", filePath)
		return
	}
	util.Log.Infof("load sample:\n%s", content)

	bytes, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		util.Log.Errorf("%+v", err)
		return
	}
	unGzip, err := util.UnGzip(bytes)
	if err != nil {
		util.Log.Errorf("%+v", err)
		return
	}
	sample = unGzip
	byteBuf := parse.ToByteBuf(unGzip)
	packets := immotors.To_Packets(byteBuf)
	if len(packets) != 10 {
		util.Log.Infof("sample packets len[%d],must be 10", len(packets))
		return
	}

	w := &kafka.Writer{
		Addr:                   kafka.TCP(kafkaAddress...),
		Topic:                  topic,
		Balancer:               &kafka.LeastBytes{},
		BatchTimeout:           100 * time.Millisecond,
		BatchSize:              1000,
		Async:                  true,
		AllowAutoTopicCreation: true,
	}
	defer w.Close()

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

	vins := GetVins(num, startIndex)
	if num < period {
		for _, e := range vins {
			go startClient(ctx, e, w)
			time.Sleep(1 * time.Second)
		}
	} else {
		//分组批次
		batchNum := num / period
		for i := 0; i < period; i++ {
			if i == period-1 {
				for _, e := range vins[i*batchNum:] {
					go startClient(ctx, e, w)
				}
			} else {
				for _, e := range vins[i*batchNum : (i+1)*batchNum] {
					go startClient(ctx, e, w)
				}
			}
			time.Sleep(1 * time.Second)
		}
	}

	for {
		time.Sleep(time.Hour)
	}
}

func startClient(ctx context.Context, vin string, w *kafka.Writer) {
	byteBuf := parse.ToByteBuf(sample)
	packets := immotors.To_Packets(byteBuf)
	for _, packet := range packets {
		if packet.F_evt_D00A != nil {
			packet.F_evt_D00A.F_VIN = vin
		}
	}
	atomic.AddUint32(&clientNum, 1)

	var sendTs = time.Now().UnixMilli()
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

		sendTss := sendTs / 1000
		for i, packet := range packets {
			if packet.F_evt_0001 != nil {
				packet.F_evt_0001.F_TBOXSysTim = sendTss - int64(10-i)
			}
		}
		res, err := immotors.ToBin(vin, "EP33", sendTss, packets)

		if err != nil {
			util.Log.Errorf("%+v", err)
			return
		}
		bytes, err := res.ToBytes()
		if err != nil {
			util.Log.Errorf("%+v", err)
			return
		}
		atomic.AddUint32(&sendNum, 1)
		err = w.WriteMessages(ctx, kafka.Message{
			Value: bytes,
		})
		if err != nil {
			util.Log.Errorf("%+v", err)
			return
		}
		sendTs += int64(period) * 1000
	}
}
