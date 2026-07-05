package main

import (
	"fmt"
	"log"
	"net/http"
)

func ping(w http.ResponseWriter, _ *http.Request) {
	if _, err := fmt.Fprintln(w, "pong"); err != nil {
		log.Printf("ping: %v", err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /ping", ping)

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
