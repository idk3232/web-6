package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, web!")
	})

	http.ListenAndServe(":8080", nil)
}
