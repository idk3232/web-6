package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api/user", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		w.Write([]byte(fmt.Sprintf("Hello,%s!", name)))
	})
	http.ListenAndServe(":9000", nil)
}
