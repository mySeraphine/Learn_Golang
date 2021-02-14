package main

import (
	"fmt"
	"net/http"
)

func main() {
	server:=http.Server{
		Addr: "localhost:2333",
	}

	http.HandleFunc("/process", func(writer http.ResponseWriter, request *http.Request) {
		request.ParseForm()
		fmt.Fprintln(writer,request.PostForm)
	})

	server.ListenAndServe()
}
