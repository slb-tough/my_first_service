package main

import (
	"net/http"
	"io"
	"log"
	"im/tcp_"
)

func TestHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "test info......."+req.Method)
}

func main() {
	tcp_.Init_tcp()
	http.HandleFunc("/test", TestHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
