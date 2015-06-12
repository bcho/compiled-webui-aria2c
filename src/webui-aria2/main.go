package main

import (
	"flag"
	"fmt"
	"net/http"
)

var (
	host = "127.0.0.1"
	port = 8080
)

func main() {
	flag.Parse()

	http.Handle("/", http.FileServer(assetFS()))

	serve := fmt.Sprintf("%s:%d", host, port)
	if err := http.ListenAndServe(serve, nil); err != nil {
		panic(err)
	}
}

func init() {
	flag.StringVar(&host, "host", host, "serve host")
	flag.IntVar(&port, "port", port, "serve port")
}
