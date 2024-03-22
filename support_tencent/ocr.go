package support_tencent

import (
	"bcd-util/util"
	"encoding/base64"
	"github.com/pkg/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	ocr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ocr/v20181119"
	"os"
)

var Client_ocr *ocr.Client

func init() {
	credential := common.NewCredential("xxx", "xxx")
	clientProfile := profile.NewClientProfile()
	clientProfile.HttpProfile.ReqTimeout = 30
	var err error
	Client_ocr, err = ocr.NewClient(credential, regions.Beijing, clientProfile)
	if err != nil {
		util.Log.Errorf("%+v", err)
		return
	}
}

func RecognizeTableAccurateOCR(filePath string) ([]byte, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	tableOCRRequest := ocr.NewRecognizeTableAccurateOCRRequest()
	base64Str := base64.StdEncoding.EncodeToString(file)
	tableOCRRequest.ImageBase64 = &base64Str
	tableOCRResponse, err := Client_ocr.RecognizeTableAccurateOCR(tableOCRRequest)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	decodeString, err := base64.StdEncoding.DecodeString(*tableOCRResponse.Response.Data)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return decodeString, nil
}
