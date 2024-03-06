package immotors_json

import (
	"bcd-util/support_parse/immotors"
	"bcd-util/util"
	"context"
	"encoding/json"
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
		Use:   "immotors_json",
		Short: "智己飞凡压测模拟器(json)、车辆vin以TEST000000开头、后面按照顺序生成序号",
		Run: func(cmd *cobra.Command, args []string) {
			Start()
		},
	}
	cmd.Flags().IntVarP(&period, "period", "p", 30, "报文上报间隔(秒)")
	cmd.Flags().IntVarP(&startIndex, "startIndex", "s", 0, "车辆开始索引(从0开始)")
	cmd.Flags().IntVarP(&num, "num", "n", 1, "压测车辆数")
	cmd.Flags().StringSliceVarP(&kafkaAddress, "kafkaAddress", "a", []string{"127.0.0.1:9092"}, "kafka地址")
	cmd.Flags().StringVarP(&topic, "topic", "t", "gw-test", "kafka topic")
	cmd.Flags().StringVarP(&filePath, "filePath", "f", "sample.txt", "存储样例报文的路径(必须且仅存储一条报文在文件中、格式为包含至少10条json报文)")
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
	sample = []byte(content)

	jsonObj := immotors.Json{}
	err = json.Unmarshal(sample, &jsonObj)
	if err != nil {
		util.Log.Errorf("sample Unmarshal error:\n%+v", err)
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
	if num <= period {
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
	jsonObj := immotors.Json{}
	_ = json.Unmarshal(sample, &jsonObj)

	jsonObj.Tboxinfo.VIN = vin
	for _, channel := range jsonObj.Channels {
		if channel.ID == 1 {
			data := channel.Data
			for _, e := range data {
				e["VIN"] = vin
			}
			break
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
		startTss := sendTss - 29
		jsonObj.FileCreationTime = startTss
		for i, channel := range jsonObj.Channels {
			jsonObj.Channels[i].Starttime = startTss
			if channel.ID == 1 {
				data := channel.Data
				for j, e := range data {
					e["TBOXSysTim"] = startTss + int64(j)
				}
			}
		}
		marshal, err := json.Marshal(jsonObj)
		if err != nil {
			util.Log.Errorf("%+v", err)
			return
		}
		gzip, err := util.Gzip(marshal)
		if err != nil {
			util.Log.Errorf("%+v", err)
			return
		}
		atomic.AddUint32(&sendNum, 1)
		err = w.WriteMessages(ctx, kafka.Message{
			Value: gzip,
		})
		if err != nil {
			util.Log.Errorf("%+v", err)
			return
		}
		sendTs += int64(period) * 1000
	}
}
