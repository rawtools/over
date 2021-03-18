package main

import (
	"fmt"
	"os"

	"raw.tools/over/cmd"
)

var overCmd = cmd.RootCommand("over")

func init() {
	overCmd.AddCommand(cmd.ListCmd)
	overCmd.AddCommand(cmd.NewCmd)
}

func main() {
	if err := overCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
