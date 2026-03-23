package main

import (
	"log"
	"net/http"
	"sum/internal"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/sum", internal.HandlerSum)

	log.Println("🚀 Server starting on :8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("Server failed:", err)
	}
}