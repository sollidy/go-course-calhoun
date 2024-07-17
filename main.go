package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Mark Pavlov</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Send email</h1><p>to send email click <a href=\"mailto:sample@sample.ru\">sample@sample.ru</a></p>")
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<p>Page not found</p>")
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		notFoundHandler(w, r)
	}
}

func main() {
	http.HandleFunc("/", pathHandler)
	fmt.Println("Starting the server on: 3001 ...")
	http.ListenAndServe(":3001", nil)
}
