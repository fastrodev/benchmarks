package main

import (
	"fmt"
	"net/http"
)

func main() {
	port := ":9001"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	fmt.Println("server started at localhost:9001")
	http.ListenAndServe(port, nil)
}
