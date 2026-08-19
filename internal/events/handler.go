package events

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Handler struct {
	DB *sql.DB
}

// -- Register
func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /events", h.create)
	mux.HandleFunc("GET /events", h.list)
	mux.HandleFunc("GET /events/{id}", h.get)
}

// POST /events
func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	var event Event

	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		message := fmt.Errorf("request body must be valid JSON: %w", err)
		http.Error(w, message.Error(), http.StatusBadRequest)
		return
	}

	if err := event.CreateValidator(); err != nil {
		message := fmt.Errorf("incomplete event body for creation: %w", err)
		http.Error(w, message.Error(), http.StatusBadRequest)
		return
	}

	repo := Repository{DB: h.DB}
	if err := repo.Create(r.Context(), &event); err != nil {
		fmt.Println("could not create event:", err)
		http.Error(w, "could not create event", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(event)
}

// GET /events
// missing: filters, search, paginator
// future: makes sense if the event list item has less data than the detail item
func (h *Handler) list(w http.ResponseWriter, r *http.Request) {
	repo := Repository{DB: h.DB}
	events, err := repo.List(r.Context())
	if err != nil {
		fmt.Println("could not list events:", err)
		http.Error(w, "could not list events", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}

// GET /events/{id}
// future: makes sense if the event detail item has more details than the list itm
func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		message := fmt.Errorf("id cannot be converted to a number: %w", err)
		http.Error(w, message.Error(), http.StatusBadRequest)
		return
	}

	repo := Repository{DB: h.DB}
	event, err := repo.Get(r.Context(), id)
	if err != nil {
		fmt.Println("could not get event:", err)
		http.Error(w, "could not get event", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(event)
}
