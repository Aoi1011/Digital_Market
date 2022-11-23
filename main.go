package main

import (
	"fmt"
	"net/http"
)

type Router struct{}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		http.Error(w, "Page not found", http.StatusNotFound)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>Welcome to my awesome site!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"a@gmail.com\">a@gmail.com</a></p>")
}

func main() {
	fmt.Printf("Starting the server on :3000...")
	http.ListenAndServe(":3000", http.HandlerFunc(pathHandler)) // http.HandlerFunc is a type, so when we write the following code we are converting pathHandler to the http.HandlerFunc type
}
