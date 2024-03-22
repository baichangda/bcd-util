package table

import (
	"bcd-util/support_tencent"
	"bcd-util/util"
	"github.com/spf13/cobra"
	"io/fs"
	"os"
	"strings"
)

func Cmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "table",
		Short: `读取当前文件夹下面所有的jpg、jpeg、png文件识别成对应的xlsx文件结果`,
		Run: func(cmd *cobra.Command, args []string) {
			dir, err := os.ReadDir("./")
			if err != nil {
				util.Log.Errorf("%+v", err)
				return
			}
			for _, f := range dir {
				if f.IsDir() {
					continue
				}
				name := f.Name()
				upper := strings.ToLower(name)
				if strings.HasSuffix(upper, ".jpg") || strings.HasSuffix(upper, ".jpeg") || strings.HasSuffix(upper, ".png") {
					namePre := name[:strings.LastIndex(name, ".")]
					res, err := support_tencent.RecognizeTableAccurateOCR(name)
					if err == nil {
						nameXlsx := namePre + ".xlsx"
						err := os.WriteFile(nameXlsx, res, fs.ModePerm)
						if err != nil {
							util.Log.Errorf("%+v", err)
						}
						util.Log.Infof("ocr [%s] -> [%s]", name, nameXlsx)
					} else {
						util.Log.Errorf("%+v", err)
					}
				}
			}
		},
	}
	return &cmd
}

func Main() {
	cobra.MousetrapHelpText = ""
	_ = Cmd().Execute()
}
