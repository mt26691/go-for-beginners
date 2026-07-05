package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mt26691/go-for-beginners/internal/task"
)

func ping(w http.ResponseWriter, _ *http.Request) {
	if _, err := fmt.Fprintln(w, "pong"); err != nil {
		log.Printf("ping: %v", err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /ping", ping)

	store := task.NewStore()
	handler := task.NewHandler(store)
	handler.Routes(mux)

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
