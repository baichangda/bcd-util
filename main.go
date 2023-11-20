package main

import (
	"bcd-util/cmd_monitor"
	"bcd-util/cmd_ocr"
	"bcd-util/cmd_pressTest"
	"bcd-util/cmd_redis"
	"bcd-util/cmd_simlator"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{}

func main() {
	rootCmd.AddCommand(cmd_redis.Cmd())
	rootCmd.AddCommand(cmd_monitor.Cmd())
	rootCmd.AddCommand(cmd_pressTest.Cmd())
	rootCmd.AddCommand(cmd_simlator.Cmd())
	rootCmd.AddCommand(cmd_ocr.Cmd())
	rootCmd.Execute()
	//cmd_ocr.Main()
}
