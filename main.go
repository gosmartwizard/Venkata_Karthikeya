package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1> Welcome to Awesome site !!! </h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "To get in touch, please send an email to <a href=\"mailto:support@karthikeya.com\">support@karthikeya.com</a>.")
	}

}

func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("Web Server started")
	http.ListenAndServe(":4949", nil)
}
