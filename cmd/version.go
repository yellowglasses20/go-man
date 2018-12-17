package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	goCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of go-man",
	Long:  `All software has versions.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v0.1 -- HEAD")
	},
}
