package support_gnet

import (
	"github.com/panjf2000/gnet/v2"
	"log"
	"testing"
)

type echoServer struct {
	gnet.BuiltinEventEngine

	eng       gnet.Engine
	addr      string
	multicore bool
}

func (es *echoServer) OnBoot(eng gnet.Engine) gnet.Action {
	es.eng = eng
	log.Printf("echo server with multi-core=%t is listening on %s\n", es.multicore, es.addr)
	return gnet.None
}

func (es *echoServer) OnTraffic(c gnet.Conn) gnet.Action {
	buf, _ := c.Next(-1)
	println(string(buf))
	_ = c.AsyncWrite(buf, nil)
	//_, _ = c.Write(buf)
	return gnet.None
}

func Test1(t *testing.T) {
	echo := &echoServer{addr: "tcp://:9000", multicore: true}
	log.Fatal(gnet.Run(echo, echo.addr, gnet.WithMulticore(true)))
}
