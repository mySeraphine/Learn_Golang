//自定义Handler

package main

import "net/http"

type helloHandler struct{}

func (m *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

type aboutHandler struct{}

func (a *aboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About Page."))
}

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome!"))
}

func main() {
	mh := helloHandler{}
	ab := aboutHandler{}
	server := http.Server{
		Addr:    "localhost:2333",
		Handler: nil, //use a DefaultserveMux
	}
	http.Handle("/hello", &mh)
	http.Handle("/about", &ab)
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Home Page!")) //use HandleFunc
	})
	http.HandleFunc("/welcome", welcome)
	server.ListenAndServe()
}
