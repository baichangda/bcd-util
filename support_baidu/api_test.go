package support_baidu

import "testing"

func TestOcrDocConvert(t *testing.T) {
	res, err := OcrDocConvert("", "", "d:/test.pdf", "1")
	if err != nil {
		t.Errorf("%+v", err)
	}
	t.Log(res)
}

func TestOcrDocConvertResult(t *testing.T) {
	res, err := OcrDocConvertResult("sKCl8EKr")
	if err != nil {
		t.Errorf("%+v", err)
	}
	t.Log(res)
}

func TestInitBaiduFanYiJs(t *testing.T) {
	InitBaiduFanYiJs()
}
