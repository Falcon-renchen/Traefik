package main

import "net/http"

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("user api1"))
	})
	http.ListenAndServe(":8001",nil)
}
