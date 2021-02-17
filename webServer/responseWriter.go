package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)
func writeExample(write http.ResponseWriter,resquest *http.Request)  {
	//go中声明原始字符串需要用顿号`，而不是使用单引号‘
	str:= `<html>			
<head><title>Go web</title></head>
<body><h1>Hello_World</h1></body>
</html> `

	write.Write([]byte(str))
}

func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w,"No such server,try another one")
}

func headerExample(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Location","http://google.com")
	w.WriteHeader(302)
}

type Post struct {
	User string
	Thread []string
}

func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	post:=&Post{
		User:"Jeff",
		Thread:[]string{"first","second","third"},
	}
	json,_:=json.Marshal(post)
	w.Write(json)
}

func main() {
	server:=http.Server{
		Addr: "localhost:2333",
	}
	http.HandleFunc("/write",writeExample)
	http.HandleFunc("/writeheader",writeHeaderExample)
	http.HandleFunc("/redirect",headerExample)
	http.HandleFunc("/json",jsonExample)
	server.ListenAndServe()
}
