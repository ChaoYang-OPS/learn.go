package main

import (
	"log"
	"os"
)

func logger() {
	log.Printf("%d\n", 5)

	foots, err := os.OpenFile("my.log", os.O_APPEND|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return
	}
	defer foots.Close()
	logWriter := log.New(foots, "CHINA", log.Ldate|log.Lmicroseconds)
	logWriter.Printf("%s\n", "hello world")
	logWriter.Println("Hee")
}

func main() {
	logger()
}
