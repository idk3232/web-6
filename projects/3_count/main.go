package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	counter := 0
	http.HandleFunc("/count", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			w.Write([]byte(strconv.Itoa(counter)))
		}
		if r.Method == http.MethodPost {
			r.ParseForm()
			numberString := r.Form.Get("count")
			a, err := strconv.Atoi(numberString)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintln(w, "это не число")
				return
			}
			counter += a
		}
		if r.Method != http.MethodPost && r.Method != http.MethodGet {
			http.Error(w, "method is not allowed", http.StatusMethodNotAllowed)

		}

	})
	http.ListenAndServe(":3333", nil)
}
