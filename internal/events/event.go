package events

import "time"

type Event struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Venue       string    `json:"venue"`
	City        string    `json:"city"`
	StartsAt    time.Time `json:"startsAt"`
	ImageURL    string    `json:"imageUrl"`
	Description string    `json:"description"`
}
