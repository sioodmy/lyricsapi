package main

import (
	"flag"
	"net/http"

	"github.com/sioodmy/lyricsapi/internal/get"
)

func main() {
	listen := flag.String("l", "0.0.0.0:3000", "IP and port to listen on")
	mux := http.NewServeMux()

	mux.HandleFunc("/api/{query}", get.GetHandle)

	flag.Parse()
	http.ListenAndServe(*listen, mux)
}
