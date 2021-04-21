package main

import (
	"fmt"
	"net/http"
)

type Engine struct{}

func main() {
	engine := new(Engine)
	http.ListenAndServe(":9999", engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path=%q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q]=%q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 not found :%s \n", req.URL)

	}
}
