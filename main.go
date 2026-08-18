package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Event struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	City  string `json:"city"`
}

const address = ":8080"

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ok")
	})

	http.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}

		events := []Event{
			{
				ID:    1,
				Title: "lowlands",
				City:  "Biddinghuizen",
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
