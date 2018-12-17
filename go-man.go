package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/yellowglasses20/go-man/cmd"
	"github.com/yellowglasses20/go-man/term"
)

func main() {

	_, stdout, _ := term.StdStreams()
	logrus.SetOutput(stdout)

	goman := cmd.NewGoManCommand()

	if err := goman.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
