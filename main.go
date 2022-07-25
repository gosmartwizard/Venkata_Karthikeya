package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1> Welcome to Awesome site !!! <h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("Web Server started")
	http.ListenAndServe(":4949", nil)
}
