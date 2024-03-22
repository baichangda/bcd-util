package support_tencent

import (
	"encoding/base64"
	ocr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ocr/v20181119"
	"io/fs"
	"os"
	"testing"
)

func TestOcr(t *testing.T) {
	file, err := os.ReadFile("d:/test.png")
	if err != nil {
		t.Errorf("%+v", err)
		return
	}
	tableOCRRequest := ocr.NewRecognizeTableAccurateOCRRequest()
	base64Str := base64.StdEncoding.EncodeToString(file)
	tableOCRRequest.ImageBase64 = &base64Str
	tableOCRResponse, err := Client_ocr.RecognizeTableAccurateOCR(tableOCRRequest)
	if err != nil {
		t.Errorf("%+v", err)
		return
	}
	decodeString, err := base64.StdEncoding.DecodeString(*tableOCRResponse.Response.Data)
	if err != nil {
		t.Errorf("%+v", err)
		return
	}
	err = os.WriteFile("test.xlsx", decodeString, fs.ModePerm)
	if err != nil {
		t.Errorf("%+v", err)
		return
	}
}
