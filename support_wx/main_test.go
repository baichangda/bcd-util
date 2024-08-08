package support_wx

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"testing"
	"time"
)

func Test1(t *testing.T) {
	engine := gin.Default()
	engine.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	engine.GET("/wx/handle", func(ctx *gin.Context) {
		t.Logf("")
		ctx.String(200, "wxToken")
	})
	err := engine.Run(":11111")
	if err != nil {
		t.Errorf("%+v", err)
		t.Fail()
	}
}
