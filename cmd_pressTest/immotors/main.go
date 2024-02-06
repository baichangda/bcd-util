package immotors

import (
	"bcd-util/support_parse/immotors"
	"bcd-util/support_parse/parse"
	"bcd-util/util"
	"encoding/base64"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"net"
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
var address string
var filePath string

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

func Cmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "immotors",
		Short: "智己飞凡、车辆vin以TEST000000开头、后面按照顺序生成序号",
		Run: func(cmd *cobra.Command, args []string) {
			Start()
		},
	}
	cmd.Flags().IntVarP(&period, "period", "p", 10, "报文上报间隔(秒)")
	cmd.Flags().IntVarP(&startIndex, "startIndex", "s", 0, "车辆开始索引(从0开始)")
	cmd.Flags().IntVarP(&num, "num", "n", 1, "压测车辆数")
	cmd.Flags().StringVarP(&address, "address", "a", "127.0.0.1:6666", "网关tcp地址")
	cmd.Flags().StringVarP(&filePath, "filePath", "f", "sample.txt", "存储样例报文的路径(必须且仅存储一条报文在文件中、格式为包含10条的原始报文base64的)")
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

	go func() {
		for {
			util.Log.Infof("client[%d] sendSpeed[%d/s]", atomic.LoadUint32(&clientNum), atomic.SwapUint32(&sendNum, 0)/3)
			time.Sleep(3 * time.Second)
		}
	}()
	vins := getVins()
	if num < period {
		for _, e := range vins {
			go startClient(e)
			time.Sleep(1 * time.Second)
		}
	} else {
		//分组批次
		batchNum := num / period
		for i := 0; i < period; i++ {
			if i == period-1 {
				for _, e := range vins[i*batchNum:] {
					go startClient(e)
				}
			} else {
				for _, e := range vins[i*batchNum : (i+1)*batchNum] {
					go startClient(e)
				}
			}
			time.Sleep(1 * time.Second)
		}
	}

	for {
		time.Sleep(time.Hour)
	}
}

func startClient(vin string) {
	byteBuf := parse.ToByteBuf(sample)
	packets := immotors.To_Packets(byteBuf)
	for _, packet := range packets {
		packet.F_evt_D00A.F_VIN = vin
	}
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
	for {
		waitMills := sendTs - time.Now().UnixMilli()
		if waitMills > 0 {
			time.Sleep(time.Duration(waitMills) * time.Millisecond)
		}
		res, err := ToJson(vin, "EP33", sendTs, packets)
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
		_, err = dial.Write(bytes)
		if err != nil {
			util.Log.Errorf("%+v", err)
			return
		}
		sendTs += int64(period) * 1000
	}
}
