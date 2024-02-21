package immotors

import (
	"bcd-util/support_parse/parse"
	"bcd-util/util"
	"encoding/base64"
	"encoding/json"
	"github.com/pkg/errors"
	"strconv"
	"time"
)

type Bin struct {
	FileName    string `json:"fileName"`
	FileContent string `json:"fileContent"`
	Timestamp   int64  `json:"timestamp"`
	MessageId   string `json:"messageId"`
	Ext         Ext    `json:"ext"`
}

type Ext struct {
	VehicleModel string `json:"vehicleModel"`
}

func ToBin(vin string, vehicleType string, ts int64, packets []Packet) (*Bin, error) {
	ts = ts - 9
	dateStr := time.Unix(ts, 0).Format("20060102150405")
	buf_empty := parse.ToByteBuf_empty()
	Write_Packets(packets, buf_empty)
	toBytes := buf_empty.ToBytes()
	//util.Log.Infof("--------------\n%s", hex.EncodeToString(toBytes))
	r, err := util.Gzip(toBytes)
	if err != nil {
		return nil, err
	}
	return &Bin{
		FileName:    vin + "_" + dateStr[0:8] + "_" + dateStr[8:] + "_E_V2.0.6.8.bl.gz",
		FileContent: base64.StdEncoding.EncodeToString(r),
		Timestamp:   ts,
		MessageId:   vin + strconv.FormatInt(ts, 10),
		Ext:         Ext{VehicleModel: vehicleType},
	}, nil
}

func (e *Bin) ToBytes() ([]byte, error) {
	marshal, err := json.Marshal(e)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return marshal, nil
}
