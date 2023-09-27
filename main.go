package main

import "net/http"

func main() {
	http.HandleFunc("/ping",
		func(w http.ResponseWriter, req *http.Request) {
			w.Write([]byte(`{"message":"pong"}`))
		})
	http.ListenAndServe(":8080", nil)
}
