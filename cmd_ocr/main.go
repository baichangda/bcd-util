package cmd_ocr

import (
	"bcd-util/support_baidu"
	"bcd-util/util"
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"io"
	"io/fs"
	"net/http"
	"strconv"
	"strings"
)

//go:embed resource
var FS embed.FS

var port int

func Cmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "ocr",
		Short: `启动ocr web服务`,
		Run: func(cmd *cobra.Command, args []string) {
			gin.SetMode(gin.ReleaseMode)
			engine := gin.New()
			engine.GET("/", func(c *gin.Context) {
				c.Redirect(http.StatusMovedPermanently, "/resource/index.html")
			})

			sub, err2 := fs.Sub(FS, "resource")
			if err2 != nil {
				util.Log.Errorf("%+v", err2)
				return
			}
			engine.StaticFS("/resource", http.FS(sub))
			//engine.Static("/resource", "cmd_ocr/resource")
			engine.POST("/ocr", func(ctx *gin.Context) {
				all, err := io.ReadAll(ctx.Request.Body)
				if err != nil {
					_ = ctx.Error(errors.WithStack(err))
				}
				split := strings.Split(string(all), ",")
				json, err := support_baidu.OcrAccurate(split[1], "", "", "", split[0], "", "", "")
				if err != nil {
					_ = ctx.Error(errors.WithStack(err))
					return
				}
				words_result := json.Get("words_result")
				if words_result.Exists() {
					sb := strings.Builder{}
					for _, cur := range words_result.Array() {
						sb.WriteString(cur.Get("words").Str)
						sb.WriteString("\n")
					}
					_, err := ctx.Writer.WriteString("0" + sb.String())
					if err != nil {
						_ = ctx.Error(errors.WithStack(err))
						return
					}
				} else {
					_, err := ctx.Writer.WriteString("1" + fmt.Sprintf("失败、错误信息:\n%s", json.Raw))
					if err != nil {
						_ = ctx.Error(errors.WithStack(err))
						return
					}
				}
			})
			err := engine.Run(":" + strconv.Itoa(port))
			if err != nil {
				util.Log.Errorf("%+v", err)
			}
		},
	}
	cmd.Flags().IntVarP(&port, "port", "p", 23456, "端口")

	return &cmd
}

func Main() {
	cobra.MousetrapHelpText = ""
	_ = Cmd().Execute()
}
