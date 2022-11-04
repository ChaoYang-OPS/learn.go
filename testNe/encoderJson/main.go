package main

import (
	"compress/zlib"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/bytedance/sonic"
	"io"
	"os"
)

type Student struct {
	Name string
	Age  int
}

func base64Demo() {
	bs := []byte{1, 4, 2, 3}
	str := base64.StdEncoding.EncodeToString(bs)
	fmt.Println(str)

	if cont, err := base64.StdEncoding.DecodeString(str); err == nil {
		fmt.Println(cont)
	} else {
		fmt.Println(err)
	}
}

func compressDemo() {
	fin, err := os.Open("main.go")
	if err != nil {
		return
	}
	defer fin.Close()
	stat, _ := fin.Stat()
	fmt.Printf("压缩前文件的大小 %d B\n", stat.Size())
	fount, err := os.OpenFile("time.zlib", os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return
	}
	defer fount.Close()
	bs := make([]byte, 1024)
	writer := zlib.NewWriter(fount)
	for {
		n, err := fin.Read(bs)
		if err == nil {
			writer.Write(bs[:n])
		} else {
			if err == io.EOF {
				return
			} else {
				fmt.Println(err)
				return
			}
		}
	}
	writer.Flush()
	founts2, err := os.Open("time.zlib")
	if err != nil {
		return
	}
	readers, err := zlib.NewReader(founts2)
	if err != nil {
		return
	}
	io.Copy(os.Stdout, readers)
	readers.Close()
}
func main() {
	stu := Student{
		Name: "TF",
		Age:  19,
	}
	bs, _ := json.Marshal(stu)
	fmt.Printf(string(bs))
	bs2, _ := sonic.Marshal(stu)
	fmt.Println(string(bs2))
	base64Demo()
	compressDemo()
}
