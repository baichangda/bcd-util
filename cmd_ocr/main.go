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

var https bool
var port int
var certFile string
var keyFile string

func Cmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "ocr",
		Short: `启动ocr https服务`,
		Run: func(cmd *cobra.Command, args []string) {
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
			if https {
				err := engine.RunTLS(":"+strconv.Itoa(port), certFile, keyFile)
				if err != nil {
					util.Log.Errorf("%+v", err)
				}
			} else {
				if port == 443 {
					err := engine.Run(":" + strconv.Itoa(80))
					if err != nil {
						util.Log.Errorf("%+v", err)
					}
				} else {
					err := engine.Run(":" + strconv.Itoa(port))
					if err != nil {
						util.Log.Errorf("%+v", err)
					}
				}
			}
		},
	}
	cmd.Flags().BoolVarP(&https, "https", "t", true, "是否https服务")
	cmd.Flags().IntVarP(&port, "port", "p", 443, "https默认443、http默认是80、如果手动指定了其他端口则以参数为准")
	cmd.Flags().StringVarP(&certFile, "certFile", "c", "./crt.pem", "证书crt文件地址")
	cmd.Flags().StringVarP(&keyFile, "keyFile", "k", "./key.pem", "证书key文件地址")

	return &cmd
}

func Main() {
	cobra.MousetrapHelpText = ""
	_ = Cmd().Execute()
}
