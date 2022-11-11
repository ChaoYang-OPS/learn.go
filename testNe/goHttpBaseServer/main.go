package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
func main() {
	http.HandleFunc("/", HelloHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Panic(err)
	}
}
