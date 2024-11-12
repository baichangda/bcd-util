package server

import (
	"testing"
)

func TestCmd(t *testing.T) {
	cmd := Cmd()
	err := cmd.ParseFlags([]string{
		"-e0/15 *",
		"-uroot:bcd@tcp(192.168.23.129:3306)/bcd?multiStatements=true&charset=utf8mb4&parseTime=True&loc=Local",
		"-a192.168.23.129:6379",
		"-pbcd",
		//"-c",
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
