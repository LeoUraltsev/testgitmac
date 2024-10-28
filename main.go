package main

import (
	"encoding/json"
	"net/http"
)

type Result struct {
	Status string
	Text   string
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		answer := Result{
			Status: "OK",
			Text:   "Hello world",
		}
		res, _ := json.Marshal(answer)

		w.Write(res)
	})

	s := http.Server{
		Addr: "localhost:8080",
	}

	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}

}
