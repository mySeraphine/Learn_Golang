package main

import (
	"fmt"
	"net/http"
)

func main() {
	server:=http.Server{
		Addr: "localhost:2333",
	}
	http.HandleFunc("/url", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer,request.URL.Fragment)
	})
	server.ListenAndServe()
}
