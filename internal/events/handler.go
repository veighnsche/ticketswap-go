package events

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
)

type Handler struct {
	DB *sql.DB
}

func (h Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /events", h.list)
	mux.HandleFunc("GET /events/{id}", h.get)
}

// GET /events
func (h Handler) list(w http.ResponseWriter, r *http.Request) {
	events := []Event{
		{
			ID:    1,
			Title: "lowlands",
			City:  "Biddinghuizen",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}

// GET /events/{id}
func (h Handler) get(w http.ResponseWriter, r *http.Request) {
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
}
