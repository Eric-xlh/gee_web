package main

import (
	"fmt"
	"gee_web/gee"
	"log"
	"net/http"
)

func main() {
	gee := gee.New()
	gee.GET("/", indexHandle)
	gee.POST("/hello", helloHandle)
	log.Fatal(gee.Run(":9999"))
}

func indexHandle(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.path=%q\n", req.URL.Path)
}

func helloHandle(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "head[%q]=%q\n", k, v)
	}
}
