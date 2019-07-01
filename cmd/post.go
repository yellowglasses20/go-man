package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/hokaccha/go-prettyjson"
	"github.com/spf13/cobra"
)

func init() {
	goCmd.AddCommand(postCmd)
	goCmd.PersistentFlags().StringP("data", "d", "", "HTTP post body parameter (-d \"name=go-man&key=value\")")
}

var postCmd = &cobra.Command{
	Use:                   "post [OPTIONS] [ADDRESS]",
	Short:                 "http post request",
	SilenceUsage:          true,
	SilenceErrors:         true,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		post(args)
	},
}

func post(args []string) {

	if len(args) == 0 {

		fmt.Println("Need Input IP address")
		return
	}

	postExecute(args[0])
}

func postExecute(ip string) {

	//pretty
	prettyFlag, err := goCmd.PersistentFlags().GetBool("pretty")
	if err != nil {
		fmt.Println(err)
		return
	}

	//data
	data, err := goCmd.PersistentFlags().GetString("data")
	if err != nil {
		fmt.Println(err)
		return
	}
	parseData := strings.Split(data, "&")

	values := url.Values{}
	for _, v := range parseData {
		val := strings.Split(v, "=")
		values.Add(val[0], val[1])
	}

	req, err := http.NewRequest("POST", ip, strings.NewReader(values.Encode()))
	if err != nil {
		fmt.Println(err)
		return
	}

	//header
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

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
