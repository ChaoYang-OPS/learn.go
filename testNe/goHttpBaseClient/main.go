package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
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

func complexHttpRequest() {
	reader := strings.NewReader(`{"code","success"}`)
	if req, err := http.NewRequest("POST", "http://localhost:8080/", reader); err != nil {
		panic(err)
	} else {
		// 自定义请求头
		req.Header.Add("server", "Engine")
		req.Header.Add("KYES", "KESSS")
		// 自定义cookie
		req.AddCookie(&http.Cookie{Name: "auth", Value: "123456", Expires: time.Now().Add(time.Hour)})
		client := &http.Client{
			Transport:     nil,
			CheckRedirect: nil,
			Jar:           nil,
			Timeout:       100 * time.Millisecond,
		}
		resp, err := client.Do(req) // 提交http请求
		defer resp.Body.Close()     // must be close
		fmt.Println(resp, err)
	}

}
func main() {
	get()
	post()
	complexHttpRequest()
}
