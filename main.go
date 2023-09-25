package main

import (
	"fmt"
	"net/http"
	"golang.org/x/example/hello/reverse"
)

func reverseHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("string")
	fmt.Fprintf(w, "%s", rev(query))
}

func rev(s string) string {
	return reverse.String(s)
}

func main() {
	http.HandleFunc("/reverse", reverseHandler)
	http.ListenAndServe(":8080", nil)
}
