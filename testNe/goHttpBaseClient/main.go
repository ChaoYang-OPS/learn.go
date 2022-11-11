package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func get() {
	resp, err := http.Get("http://localhost:8080/")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", resp)
	io.Copy(os.Stdout, resp.Body)
	defer resp.Body.Close() // 一定要调用关闭,否则会协程泄漏
	for k, v := range resp.Header {
		fmt.Println(k, v)
	}
	fmt.Println(resp.Status)
}

func post() {
	reader := strings.NewReader(`{"code","success"}`)
	resp, err := http.Post("http://localhost:8080/", "application/json", reader)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Status)
}
func main() {
	get()
	post()
}
