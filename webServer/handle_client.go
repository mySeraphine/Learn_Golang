package main

import "net/http"

type myHandler struct {}

func (m *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {	//一定要跟源码一样是ServeHTTP才能实现
	w.Write([]byte("Hello World"))
}

func main() {
	mh := myHandler{}
	server := http.Server{
		Addr:    "localhost:2333",
		Handler: &mh,
	}
	server.ListenAndServe()

	//http.ListenAndServe("localhost:2333",nil)	//simple Http handler
}
