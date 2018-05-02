package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	// 设置路由
	http.HandleFunc("/", SayHello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func SayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World, this is v1 with net/http.")
}
