package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/home", func(writer http.ResponseWriter, request *http.Request) {
		url:=request.URL
		query:=url.Query()

		id:=query["id"]
		log.Println(id)

		name:=query.Get("name")
		log.Println(name)
	})

	http.ListenAndServe("localhost:2333",nil)
}
