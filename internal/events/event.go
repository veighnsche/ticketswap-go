package events

import (
	"errors"
	"time"
)

type Event struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Venue       string    `json:"venue"`
	City        string    `json:"city"` // Optional
	StartsAt    time.Time `json:"startsAt"`
	ImageURL    string    `json:"imageUrl"` // Optional
	Description string    `json:"description"`
}

func (e Event) ValidateForCreate() error {
	if e.Title == "" {
		return errors.New("missing title to create event")
	}

	if e.Venue == "" {
		return errors.New("missing venue to create event")
	}

	if e.StartsAt.IsZero() {
		return errors.New("missing startsAt to create event")
	}

	if e.Description == "" {
		return errors.New("missing description to create event")
	}

	return nil
}
