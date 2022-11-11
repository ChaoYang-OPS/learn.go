package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func get() {
	resp, err := http.Get("http://localhost:8080/")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", resp)
	io.Copy(os.Stdout, resp.Body)
	defer resp.Body.Close() // 一定要调用关闭,否则会协程泄漏
}
func main() {
	get()
}
