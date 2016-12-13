package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi, this is the %q endpoint", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/bye", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is the %q endpoint, good bye", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":9090", nil))
}
