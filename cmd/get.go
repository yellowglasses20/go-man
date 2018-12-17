package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hokaccha/go-prettyjson"
	"github.com/spf13/cobra"
)

func init() {
	goCmd.AddCommand(getCmd)
	goCmd.PersistentFlags().BoolP("pretty", "p", false, "Pretty print")
}

var getCmd = &cobra.Command{
	Use:                   "get [OPTIONS] [ADDRESS]",
	Short:                 "http get request",
	SilenceUsage:          true,
	SilenceErrors:         true,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		get(args)
	},
}

func get(args []string) {

	if len(args) == 0 {

		fmt.Println("Need Input IP address")
		return
	}

	execute(args[0])
}

func execute(ip string) {

	prettyFlag, err := goCmd.PersistentFlags().GetBool("pretty")
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := http.Get(ip)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	respByte, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {

		fmt.Println("StatusCode= %v\n", resp.StatusCode)
		fmt.Printf("Status= %v\n", resp.Status)
		fmt.Println(string(respByte))
		return
	}

	if prettyFlag == true && resp.Header.Get("Content-Type") == "application/json" {

		var tmp interface{}

		err := json.Unmarshal(respByte, &tmp)
		if err != nil {
			return
		}
		s, _ := prettyjson.Marshal(tmp)
		fmt.Println(string(s))
	} else {

		fmt.Println(string(respByte))
	}
}
