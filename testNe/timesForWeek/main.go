package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func weeKdays() {
	currentlyDay := time.Now()
	fmt.Println(currentlyDay.Weekday().String())
	fmt.Println(currentlyDay.Hour())
}

func mergeFile(dir string) {
	fout, err := os.OpenFile("big.txt", os.O_CREATE|os.O_APPEND|os.O_TRUNC, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		err := fout.Close()
		if err != nil {
			return
		}
		return
	}()
	writer := bufio.NewWriter(fout)
	if files, err := ioutil.ReadDir(dir); err != nil {
		fmt.Println(err)
		return
	} else {
		for _, file := range files {
			if file.IsDir() {
				continue
			}
			fileBaseName := file.Name()
			if strings.HasSuffix(fileBaseName, ".txt") {
				if fileData, err := os.Open(filepath.Join(dir, fileBaseName)); err != nil {
					fmt.Println(err)
					continue
				} else {
					defer fileData.Close()
					reader := bufio.NewReader(fileData)
					for {
						if line, err := reader.ReadString('\n'); err != nil {
							if err == io.EOF {
								if len(line) > 0 {
									writer.WriteString(line)
								}
							}
							break
						} else {
							writer.WriteString(line)
						}
					}
				}
			}
		}
	}
}
func main() {
	weeKdays()
}
