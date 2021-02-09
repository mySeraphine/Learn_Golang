package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello_World!"))
	})

	http.ListenAndServe("localhost:2333",nil)

}
