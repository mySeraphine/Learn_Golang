package main

import (
	"html/template"
	"net/http"
)

func indexprocess(w http.ResponseWriter, r *http.Request)  {
	t, _:=template.ParseFiles("src/tmpl.html")		//在指定模板文件时，从项目根目录开始寻找，需要指定到html文件所在目录
	t.Execute(w,"hello_world")
}

func main() {
	server:=http.Server{
		Addr: "127.0.0.1:2333",
	}
	http.HandleFunc("/process",indexprocess)
	server.ListenAndServe()
}
