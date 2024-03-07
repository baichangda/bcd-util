package immotors

import (
	"bcd-util/support_parse/parse"
	"bcd-util/util"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/xuri/excelize/v2"
	"io/fs"
	"os"
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

const hexStr = "000100006466E9B3000300006583EC6E000413C339C87BDE00051DD636E16BEC00060000050000070007F024FFF8FFBC000801CC0001000000094300C4100033000AB70400620000000B000000000000000C0017AE000000000D000992400000000E082C00000000000F5208ABE0000008008169000000000801000000000000080200000000000008030250AB000000D006003E8D3A8987B4A0365D34F9A820FFFC0080180100060699FC9700008000A0001FF90050010FF915B2001B48006D011F016F26B7E0C95F722400B4F800000000D008004600000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000D00900120000000000000000000FE000010005648000D00A00554C534A4533363039364D533134303439354C53323143303141353239363030303638393836303932313735303030383935323034333839303836303334323032323030303031303231303030303333353932343437D00B00D96C7DF87DC87DF87DE87DF87DD87DF87DE87E007DE07E007DD87DD07DD87DD87DD87DE87DC87DD87DE07DD87DE87DE07DF07DD07DF87DF07E007DE87DF07DE07E087DF07DE87DD87DF87DF07DC87DF07DE07DF07DE07DC87DC87DF07DE87DF87DE87DF07DE07E007DE07DF87DE87DE87DF87DE07DF07DF07DF07DE87E007DF07DE87DE87DF87DF87E007E007E007DF87E087DF07DE07DF87DF07DF87DE07DF87DE87E087DF07DF07DD87DF87DF07DF87DE87E087DE07DE07DF87DE07E007DF87DF07DE87DF07DF87DF87DE87E087DF87DF87DF07E007DE07E08D00C00190C3E003C003C003C003E003C003C003C003D003C003C003C00D00D00190C3F003D003D003E003E003D003D003E003E003C003D003E00D00E00191830354C50454A334333353234304142434230313032303135D00F000C007D00007831800019007E00D01000380000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000D01100230000000000000000000000000000000000000000000000000000000000000000000000D012003F000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000D013004D0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000D01400230000000000000000000000000000000000000000000000000000000000000000000000D015000700000000000000D016003100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000D017002A000000000000000000000000000000000000000000000000000000000000000000000000000000000000D0180014CF639C87BB0EEB1B720000323374D920FA040300D0190010CE10053DB55580000000000000000000D01A002D000000000000000000000000000000000031323334353637383930313233343536373839303132333435363738D01B003F000000313233343536373839303132333435363738000031323334353637383930313233343536373800003132333435363738393031323334353637380000D01C002E31323334353637383930313233343536373831323334353637383930313233343536370000000000000000000000D01D000C000000000000000000000000D01F001000000000000000000000000000000000"

func TestPacket_ToJson(t *testing.T) {
	decodeString, err := hex.DecodeString(hexStr)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	byteBuf := parse.ToByteBuf(decodeString)
	packet := To_Packet(byteBuf)
	jsonObj := packet.ToJson(packet.F_evt_D00A.F_VIN)
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

func TestJson_ToBytes_fromHex(t *testing.T) {
	decodeString, err := hex.DecodeString(hexStr)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	byteBuf := parse.ToByteBuf(decodeString)
	packet := To_Packet(byteBuf)
	jsonObj := packet.ToJson(packet.F_evt_D00A.F_VIN)
	bytes, err := jsonObj.ToBytes(time.Now().UnixMilli(), 30)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	//t.Log(string(bytes))
	err = os.WriteFile("res.txt", bytes, fs.ModePerm)
	if err != nil {
		t.Fatalf("%+v", err)
	}
}

const base64Str = "H4sIAAAAAAAAA+WYXUgUURTHz7pqk666hMjsg7Qk2T5Y3hl21xkKGY1ISmhBwiDdNj/mSkj10nMTRBQoFPQilBgERWkJWfRoFtHXw/RSK/iwj0tQLCGxkrrdO7PkuO2sE7vjR93LztyPc8+d3fvj/M8sOAD65w6fhOLqA+LbY6NQUvup4c7QGJQClAC5bgO9MI4PADMAZU8UCMIQVJAxx8BFtRSah4WPzXtCDTcCjvuQWgSFdUA99C4tppYViIHXAefh5c+FPqg5AVD0uaW8/8oFt+cZe/mp5ljdDjsAqvRdqmJAHmiZjJbD7CAOyUmM5Bhuk2NyFDNYwgJGuA37yBgZlxO0R66CHMdeYiuR0QR2y1HSR5jFLLETMJCWoI1IZNaLQ8QDIv68xE9UWyERC69m6yX3KJl363N0v/Q+dA6RZ/DJceI/Qb2Se0zbg9okiV+BtMlKzdpNKoMZ+g20lfpaaseqLvC49gGtew0fvaoVZLLxd9dYG0GtBA+LAu2h1jb+YNAfCKCWQ1wr4hApnFoFrppuAR6AAh6IKCoL1Yzz1u2Zyi+zCeX19/1d8wzU9UZUD7i/ofKJ3aNq+nAlRa0FFwBbT3uOoa8uBi6d0489ffxF6bsTnIO0UVmc1E8s7ULnqMsaRy+mc3HkX7g3beBoCYBytHPByBF0rnA0lczOUcm7FY4SW4ajhIEjSpJGUQZHyHaOdoXd8NCMo3A8g6Prheao2z6OVscjKxw5xfziEWOZI28WjtgMjnwGjnx5xiP7OaoLR8w5iszZHY/CBeDomjWOpPdrcWTUtY2PR74C6pr9HFFdGzfjqEe1m6NThYhHGfnRn7o2sJojs/zIqq7R0/SlOYrZxJFU0PxoHXSt2w0TZhz1vbKbo4hFjqRc8chKnm2FI6OubXw82lp5dl13xJwj+XkGR4lCc3TaIkcKlE1azY/SHGXNs10jU/nr2nrEo9Ucbf78iOraIzOOBsbt5qjH+nu/ZY5y5dmuEYvva7ny7Gjeefa/lx9RXXtsxtGZMbt1rbcQHGXPj8zf17aGrq3NUXwTcUR1zZSjwZtGjoZtiEd9FjmS/j4e/f3/R5qulcHx9o4jnUf9SGwKdYgiz/FN7R28n+O4Fo7n6C/nF0QhiAQkcBzPi4jzizwKwholH0D/50BHBXPSDNCzV20LdKkUwI8383d/AeyQfMgxFwAA"

func TestJson_ToBytes_fromBase64(t *testing.T) {
	res1, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	res2, err := util.UnGzip(res1)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	res3 := parse.ToByteBuf(res2)
	packets := To_Packets(res3)
	packet := packets[0]
	jsonObj := packet.ToJson(packets[9].F_evt_D00A.F_VIN)
	//marshal, err := json.Marshal(jsonObj)
	//if err != nil {
	//	t.Fatalf("%+v", err)
	//}
	//t.Log(string(marshal))
	bytes, err := jsonObj.ToBytes(time.Now().UnixMilli(), 10)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	err = os.WriteFile("res.txt", bytes, fs.ModePerm)
	if err != nil {
		t.Fatalf("%+v", err)
	}
}
