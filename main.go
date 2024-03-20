package main

import (
	"bcd-util/cmd_simlator"
	"bcd-util/util"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{}

func main() {
	util.StartWeb_pprof()
	//rootCmd.AddCommand(cmd_redis.Cmd())
	//rootCmd.AddCommand(cmd_kafka.Cmd())
	//rootCmd.AddCommand(cmd_monitor.Cmd())
	//rootCmd.AddCommand(cmd_pressTest.Cmd())
	rootCmd.AddCommand(cmd_simlator.Cmd())
	//rootCmd.AddCommand(cmd_ocr.Cmd())
	//rootCmd.AddCommand(cmd_hbase.Cmd())
	_ = rootCmd.Execute()
	//cmd_ocr.Main()
	//gb32960.Main()
}
