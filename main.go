package main

import (
	"embed"
	"flag"
	"io/fs"
	"net/http"

	"github.com/sioodmy/lyricsapi/pkg/get"
)

//go:embed static
var static embed.FS

func serve_static() http.Handler {
	sub, err := fs.Sub(static, "static")
	if err != nil {
		panic(err)
	}
	file_server := http.FileServer(http.FS(sub))
	return file_server
}

func main() {
	listen := flag.String("l", "0.0.0.0:3000", "IP and port to listen on")
	mux := http.NewServeMux()

	mux.Handle("/", serve_static())
	mux.HandleFunc("/api/{query}", get.GetHandle)

	flag.Parse()
	http.ListenAndServe(*listen, mux)
}
