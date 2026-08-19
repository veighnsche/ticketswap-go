package events

import (
	"errors"
	"time"
)

type Event struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Venue       string    `json:"venue"`
	City        string    `json:"city"`
	StartsAt    time.Time `json:"startsAt"`
	ImageURL    string    `json:"imageUrl"`
	Description string    `json:"description"`
}

func (e Event) createValidator() error {
	if e.Title == "" {
		return errors.New("missing title to create event")
	}

	if e.Venue == "" {
		return errors.New("missing venue to create event")
	}

	if e.StartsAt.IsZero() {
		return errors.New("missing venue to create event")
	}

	return nil
}
