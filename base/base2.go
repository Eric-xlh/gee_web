package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	engine := NewEngine()
	log.Fatal(http.ListenAndServe(":9999", engine))
}

type Engine struct {
}

func NewEngine() *Engine {
	return &Engine{}
}

func (eg *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.path=%q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "head[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
