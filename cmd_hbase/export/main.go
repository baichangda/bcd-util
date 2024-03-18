package export

import (
	"bcd-util/util"
	"bufio"
	"context"
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/tsuna/gohbase"
	"github.com/tsuna/gohbase/hrpc"
	"io"
	"io/fs"
	"os"
)

var addr string
var table string
var startRowKey string
var endRowKey string
var filePath string
var needRowKey bool

func Cmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "export",
		Short: "导出hbase表中指定数据、每一条数据格式为(rowKey,value)、多条数据之间换行分割",
		Run: func(cmd *cobra.Command, args []string) {
			open, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, fs.ModePerm)
			defer open.Close()
			writer := bufio.NewWriter(open)

			client := gohbase.NewClient(addr)
			defer client.Close()
			scan, err := hrpc.NewScanRangeStr(context.Background(), table, startRowKey, endRowKey)
			if err != nil {
				util.Log.Errorf("%+v", err)
				return
			}
			scanner := client.Scan(scan)
			defer scanner.Close()

			count := 0
			for {
				next, err := scanner.Next()
				if err != nil {
					if err == io.EOF {
						break
					} else {
						util.Log.Errorf("%+v", err)
						return
					}
				}

				if count > 0 {
					_, err = writer.WriteRune('\n')
					if err != nil {
						util.Log.Errorf("%+v", err)
						return
					}
				}
				count++

				cells := next.Cells
				cell0 := cells[0]

				if needRowKey {
					//rowKey := string(cell.Row)
					_, err = writer.Write(cell0.Row)
					if err != nil {
						util.Log.Errorf("%+v", err)
						return
					}
					_, err = writer.WriteRune(',')
					if err != nil {
						util.Log.Errorf("%+v", err)
						return
					}
				}

				var value []byte
				if len(cells) == 1 {
					value = cell0.Value
				} else {
					temp := make(map[string]any)
					for _, e := range cells {
						temp[string(e.Qualifier)] = string(e.Value)
					}
					value, err = json.Marshal(temp)
					if err != nil {
						util.Log.Errorf("%+v", err)
						return
					}
				}
				_, err = writer.Write(value)
				if err != nil {
					util.Log.Errorf("%+v", err)
					return
				}
			}

			err = writer.Flush()
			if err != nil {
				util.Log.Errorf("%+v", err)
				return
			}

			util.Log.Infof("fetch count[%d] table[%s] startRowKey[%s] endRowKey[%s]", count, table, startRowKey, endRowKey)
		},
	}
	cmd.Flags().StringVarP(&addr, "addr", "a", "incar-dn-1,incar-nn-1,incar-nn-2", "hbase地址")
	cmd.Flags().StringVarP(&table, "table", "t", "immotors:json_12m", "hbase表名")
	cmd.Flags().StringVarP(&startRowKey, "startRowKey", "s", "2312000LS5A33LR3FB35697620231219163903####", "开始rowKey(包含)")
	cmd.Flags().StringVarP(&endRowKey, "endRowKey", "e", "2312000LS5A33LR3FB35697620231219164044####", "结束rowKey(不包含)")
	cmd.Flags().StringVarP(&filePath, "filePath", "f", "res.txt", "结果输出文件路径")
	cmd.Flags().BoolVarP(&needRowKey, "needRowKey", "r", false, "导出结果是否需要rowKey")
	return &cmd
}
