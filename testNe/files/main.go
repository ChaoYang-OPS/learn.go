package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func WriteTextFile() {
	if fout, err := os.OpenFile("go.mod3", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 777); err != nil {
		fmt.Println(err)
		return
	} else {
		defer fout.Close()
		writers := bufio.NewWriter(fout)
		writers.WriteString("第一行\n")
		writers.WriteString("第二行\n")
		writers.Flush() // 强制清空缓存，并把缓存里面的内容写入磁盘
	}
}

func walk(path string) error {
	if subFiles, err := ioutil.ReadDir(path); err != nil {
		return err
	} else {
		for _, file := range subFiles {
			fmt.Println(file.Name())
			if file.IsDir() {
				if err := walk(filepath.Join(path, file.Name())); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
func main() {

	//if fin, err := os.Open("../times/times.go"); err != nil {
	//	fmt.Println(err)
	//	return
	//} else {
	//	defer fin.Close()
	//	reader := bufio.NewReader(fin)
	//	for {
	//		if line, err := reader.ReadString('\n'); err == nil {
	//			line = strings.TrimRight(line, "\n")
	//			fmt.Println(line)
	//		} else {
	//			if err == io.EOF {
	//				if len(line) > 0 {
	//					fmt.Println(line)
	//				} else {
	//					fmt.Println(err)
	//				}
	//			}
	//			break
	//		}
	//	}
	//}
	//WriteTextFile()
	walk(".")
}
