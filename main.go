package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Event struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	City  string `json:"city"`
}

const address = ":8080"

func main() {
	// GET Health
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ok")
	})

	// GET Events List
	http.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
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

	// GET Event Detail
	http.HandleFunc("/event/{id}", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		idStr := r.PathValue("id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "id cannot be converted to a number", http.StatusBadRequest)
			return
		}

		event := Event{
			ID:    id,
			Title: "lowlands",
			City:  "Biddinghuizen",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(event)
	})

	/*
	 * ---
	 * START THE SERVER
	 * ---
	 */

	fmt.Printf("server started at http://localhost%s\n", address)

	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println("server stopped", err)
	}
}
