package cmd

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

func init() {
	goCmd.AddCommand(headCmd)
}

var headCmd = &cobra.Command{
	Use:                   "head [OPTIONS] [ADDRESS]",
	Short:                 "http head request",
	SilenceUsage:          true,
	SilenceErrors:         true,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		head(args)
	},
}

func head(args []string) {

	if len(args) == 0 {

		fmt.Println("Need Input IP address")
		return
	}

	headExecute(args[0])
}

func headExecute(ip string) {

	resp, err := http.Head(ip)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {

		fmt.Printf("StatusCode= %v\n", resp.StatusCode)
		fmt.Printf("Status= %v\n", resp.Status)
		return
	}

	for k, v := range resp.Header {
		fmt.Printf("%v: %v\n", k, v[0])
	}
}
