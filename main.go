package main

import (
	"log"
	"net/http"
)

const version string = "v3"

func main() {
	http.HandleFunc("/query", handleQuery)
	http.HandleFunc("/health", handleHealth)

	log.Printf("Starting server on: 1234")
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal(err)
	}
}

func handleQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Hello from version: " + version + "\n"))
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status": "healthy"}`))
}
