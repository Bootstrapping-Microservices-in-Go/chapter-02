package main

import (
	"fmt"
	"net/http"
)

const port = 3000

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!")
	})

	http.ListenAndServe(fmt.Sprint(":", port), mux)
}
