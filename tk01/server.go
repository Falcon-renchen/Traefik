package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	go (func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
			writer.Write([]byte("v1: user api1\r\n"))
		})
		server := &http.Server{
			Addr:    ":8001",
			Handler: mux,
		}
		server.ListenAndServe()
	})()

	go (func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
			writer.Write([]byte("v1: user api2\r\n"))
		})
		server := &http.Server{
			Addr:    ":8002",
			Handler: mux,
		}
		server.ListenAndServe()
	})()

	go (func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
			writer.Write([]byte("v2: user api\r\n"))
		})
		server := &http.Server{
			Addr:    ":8003",
			Handler: mux,
		}
		server.ListenAndServe()
	})()

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	getErr := <-c
	log.Println(getErr)
}
