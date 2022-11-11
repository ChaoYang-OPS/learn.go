package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		fmt.Println(k, v)
	}
	io.Copy(os.Stdout, r.Body)
	defer r.Body.Close()
	fmt.Fprintf(w, "Hello World\n")
}
func main() {
	http.HandleFunc("/", HelloHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Panic(err)
	}
}
