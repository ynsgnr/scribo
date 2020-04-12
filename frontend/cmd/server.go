package main

import (
	"net/http"
	"strings"
)

func main() {
	fs := http.FileServer(http.Dir("."))
	if err := http.ListenAndServe(":80", http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Add("Cache-Control", "no-cache")
		if strings.HasSuffix(req.URL.Path, ".wasm") {
			resp.Header().Set("content-type", "application/wasm")
		}
		fs.ServeHTTP(resp, req)
	})); err != nil {
		panic(err)
	}
}
