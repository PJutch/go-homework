package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "0.0.1")
	})

	http.ListenAndServe(":8080", nil)
}
