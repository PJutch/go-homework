package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type JsonInput struct {
	InputString string `json:"inputString"`
}

type JsonOutput struct {
	OutputString string `json:"outputString"`
}

func main() {
	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			fmt.Fprint(w, "0.0.1")
		}
	})
	http.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			body, err := io.ReadAll(r.Body)
			if err != nil {
				fmt.Println(err)
				return
			}

			var input JsonInput
			err2 := json.Unmarshal(body, &input)
			if err2 != nil {
				fmt.Println(err2)
				return
			}

			response, err := json.Marshal(JsonOutput{input.InputString})
			if err != nil {
				fmt.Println(err)
				return
			}

			w.Write(response)
		}
	})

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println(err)
	}
}
