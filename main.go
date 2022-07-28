package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gosmartwizard/Venkata_Karthikeya/controllers"
	"net/http"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Frequently Asked Questions</h1><p>Here is a list of questions that our users commonly ask.</p>")
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>Sorry, but we couldn't find the page you were looking for</h1>")
}

func main() {
	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers()

	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(notFound)
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")

	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")

	fmt.Println("Web Server started")
	http.ListenAndServe(":4949", r)
}
