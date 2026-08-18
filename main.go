package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const address = ":8080"

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ok")
	})

	http.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}

		events := []map[string]any{
			{
				"id":    0,
				"title": "lowlands",
				"city":  "Biddinghuizen",
			},
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(events)
	})

	// Start the server

	fmt.Printf("server started at http://localhost%s\n", address)

	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println("server stopped", err)
	}
}
