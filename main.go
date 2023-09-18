package main

import (
	"bcd-util/monitor"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{}

func main() {
	//rootCmd.AddCommand(redis.Cmd())
	rootCmd.AddCommand(monitor.Cmd())
	rootCmd.Execute()
}
