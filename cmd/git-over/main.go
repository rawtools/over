package main

import (
	"fmt"
	"os"

	"raw.tools/over/cmd"
)

var gitOverCmd = cmd.RootCommand("git-over")

// var gitOverCmd = &cobra.Command{
// 	Use:     "git-over",
// 	Short:   "git-based overlays",
// 	Long:    `over allows you to version your configuration files and workspaces settings`,
// 	Version: over.Version,
// }

func init() {
	gitOverCmd.AddCommand(cmd.ListCmd)
	// gitOverCmd.AddCommand(cmd.AddCmd)
	// gitOverCmd.AddCommand(cmd.ReleaseCmd)
	// gitOverCmd.AddCommand(cmd.ApplyCmd)
	// gitOverCmd.AddCommand(cmd.StatusCmd)
}

func main() {
	if err := gitOverCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
