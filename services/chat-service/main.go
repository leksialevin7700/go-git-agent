package main

import (
	"fmt"
	"log"
	"net/http"
)

func chatHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		query = "Why did I write this code?"
	}

	response := fmt.Sprintf(
		"👻 Past You says: I probably wrote this because it worked at 2AM 😄\nQuestion: %s",
		query,
	)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	if _, err := w.Write([]byte(response)); err != nil {
		log.Printf("failed to write response: %v", err)
	}
}

func main() {
	http.HandleFunc("/chat", chatHandler)
	log.Println("Chat service running on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("server exited with error: %v", err)
	}
}
