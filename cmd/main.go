package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Listen and serve: localhost:8080")
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "Hello, SRE!") })
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("logs: ok")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "OK")
	})
	http.ListenAndServe(":8080", mux)
}
