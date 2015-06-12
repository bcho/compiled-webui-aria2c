package main

import (
	"flag"
	"fmt"
	"net/http"
	"os/exec"
)

var (
	host   = "127.0.0.1"
	port   = 8080
	aria2c = "aria2c"
)

func serveHTTP() {
	http.Handle("/", http.FileServer(assetFS()))

	serve := fmt.Sprintf("%s:%d", host, port)
	if err := http.ListenAndServe(serve, nil); err != nil {
		panic(err)
	}
}

func serveAria2RPC() {
	// TODO customize args
	aria2cArgs := []string{
		"--enable-rpc",
		"--rpc-listen-all",
	}
	// TODO catch stderr output
	cmd := exec.Command(aria2c, aria2cArgs...)
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

func main() {
	flag.Parse()

	go serveAria2RPC()
	serveHTTP()
}

func init() {
	flag.StringVar(&host, "host", host, "serve host")
	flag.IntVar(&port, "port", port, "serve port")
	flag.StringVar(&aria2c, "aria2c", aria2c, "aria2c binary")
}
