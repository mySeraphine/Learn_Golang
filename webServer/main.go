package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("pkg/hello.html")
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err)
		return
	}

	t.Execute(w, "Hello World")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)

	server := &http.Server {
		Addr:    ":8080",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
