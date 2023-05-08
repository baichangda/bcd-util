package main

import (
	"github.com/spf13/cobra"
	"gmmc-tool/redis"
)

var rootCmd = &cobra.Command{}

func main() {
	rootCmd.AddCommand(redis.Cmd())
	rootCmd.Execute()
}
