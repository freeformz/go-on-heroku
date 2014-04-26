package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	http.HandleFunc("/", hello)
	fmt.Println("listening...")
	fmt.Println(http.ListenAndServe(":"+port, nil))
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "hello, world")
	fmt.Println("Hello From the Go App: " + req.URL.String())
}
