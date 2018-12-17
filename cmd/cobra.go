package cmd

import (
	"github.com/spf13/cobra"
)

var goCmd = &cobra.Command{
	Use:   "go-man [OPTIONS]",
	Short: "API test tool",
	Long:  `go-man is simple API test tool.`,
}

func NewGoManCommand() *cobra.Command {
	return goCmd
}
