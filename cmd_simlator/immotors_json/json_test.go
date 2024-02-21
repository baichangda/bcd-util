package immotors_json

import (
	"bcd-util/support_parse/immotors"
	"bcd-util/support_parse/parse"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/xuri/excelize/v2"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestGenerateCode(t *testing.T) {
	excelPath := "D:\\files\\20231219-智己远程监控-新车型适配-v1.70-check(1).xlsx"
	sheetIndex := 3
	rowLineNo_start := 117
	rowLineNo_end := 456
	colName_groupName := "A"
	colName_evtId := "B"
	colName_signalName := "E"
	openFile, err := excelize.OpenFile(excelPath)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	defer openFile.Close()
	sheetName := openFile.GetSheetName(sheetIndex)
	fmt.Printf("read sheetName[%s]\n", sheetName)
	groupName_signals := make(map[string][][3]string)
	for i := rowLineNo_start; i <= rowLineNo_end; i++ {
		itoa := strconv.Itoa(i)
		v, err := openFile.GetCellValue(sheetName, colName_groupName+itoa)
		if err != nil {
			t.Fatalf("%+v", err)
		}
		groupName := strings.TrimSpace(v)
		v, err = openFile.GetCellValue(sheetName, colName_evtId+itoa)
		if err != nil {
			t.Fatalf("%+v", err)
		}
		evtId := strings.TrimSpace(v)
		v, err = openFile.GetCellValue(sheetName, colName_signalName+itoa)
		if err != nil {
			t.Fatalf("%+v", err)
		}
		signalName := strings.TrimSpace(v)
		index := strings.Index(signalName, "[")
		if index != -1 {
			signalName = signalName[:index]
		}
		//fmt.Printf("%s,%s,%s\n", groupName, evtId, signalName)

		arr, ok := groupName_signals[groupName]
		if ok {
			arr = append(arr, [3]string{groupName, evtId, signalName})
		} else {
			arr = [][3]string{{groupName, evtId, signalName}}
		}
		groupName_signals[groupName] = arr
	}

	for k, v := range groupName_signals {
		fmt.Printf("  //group %s\n", k)
		fmt.Printf("  data_%s:=make(map[string]any)\n", k)
		for _, e := range v {
			fmt.Printf("  if p.F_evt_%s!=nil{\n", e[1][2:])
			fmt.Printf("  	data_%s[\"%s\"]=p.F_evt_%s.F_%s\n", k, e[2], e[1][2:], e[2])
			fmt.Printf("  }\n")
		}
		fmt.Printf(`  channels = append(channels, Channel{
	ID:                  %d,
	Starttime:           ts / 1000,
	CollectiofrequecyHz: 1,
	Data:                []map[string]any{data_%s},
  })

`, GroupName_groupId[k], k)
	}
}

func TestBinToJson(t *testing.T) {
	decodeString, err := hex.DecodeString(hexStr)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	byteBuf := parse.ToByteBuf(decodeString)
	packet := immotors.To_Packet(byteBuf)
	jsonObj := BinToJson(packet)
	marshal, err := json.MarshalIndent(jsonObj, "", "  ")
	if err != nil {
		t.Fatalf("%+v", err)
	}
	t.Log(string(marshal))

	newJsonObj := Json{}
	err = json.Unmarshal(marshal, &newJsonObj)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	t.Log(newJsonObj)
}

func TestJson_ToBytes(t *testing.T) {
	decodeString, err := hex.DecodeString(hexStr)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	byteBuf := parse.ToByteBuf(decodeString)
	packet := immotors.To_Packet(byteBuf)
	jsonObj := BinToJson(packet)
	bytes, err := jsonObj.ToBytes(time.Now().UnixMilli(), 30)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	t.Log(string(bytes))
}
