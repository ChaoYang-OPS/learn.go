package main

import (
	"fmt"
	"learn.go/testNe/rest-demo-host/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
