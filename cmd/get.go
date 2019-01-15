package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/hokaccha/go-prettyjson"
	"github.com/spf13/cobra"
)

func init() {
	goCmd.AddCommand(getCmd)
	goCmd.PersistentFlags().BoolP("pretty", "p", false, "Pretty print")
	goCmd.PersistentFlags().StringArrayP("Headers", "H", nil, "Any HTTP headers(-H \"Authorization:Bearer token\")")
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

	getExecute(args[0])
}

func getExecute(ip string) {

	prettyFlag, err := goCmd.PersistentFlags().GetBool("pretty")
	if err != nil {
		fmt.Println(err)
		return
	}
	headers, err := goCmd.PersistentFlags().GetStringArray("Headers")
	if err != nil {
		fmt.Println(err)
		return
	}

	req, err := http.NewRequest("GET", ip, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, d := range headers {
		values := strings.Split(d, ":")
		if len(values) == 2 {
			req.Header.Set(values[0], values[1])
		}
	}

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	respByte, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {

		fmt.Printf("StatusCode= %v\n", resp.StatusCode)
		fmt.Printf("Status= %v\n", resp.Status)
		fmt.Println(string(respByte))
		return
	}

	contentType := resp.Header.Get("Content-Type")
	if prettyFlag == true && strings.Contains(contentType, "application/json") {

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
