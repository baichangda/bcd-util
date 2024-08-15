package client

import (
	"testing"
)

func TestCmd(t *testing.T) {
	cmd := Cmd()
	err := cmd.ParseFlags([]string{
		"-a192.168.23.129:6379",
		"-wbcd",
		"-itest",
	})
	if err != nil {
		t.Errorf("%+v", err)
		t.Fail()
	}
	err = cmd.Execute()
	if err != nil {
		t.Errorf("%+v", err)
		t.Fail()
	}
}
