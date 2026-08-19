package events

import (
	"context"
	"database/sql"
)

type Repository struct {
	DB *sql.DB
}

func (r Repository) Create(ctx context.Context, event *Event) error {
	return r.DB.QueryRowContext(ctx, `
			INSERT INTO events (title, description, venue, city, starts_at, image_url)
			VALUES ($1, $2, $3, $4, $5, $6)
			RETURNING id
		`,
		event.Title,
		event.Description,
		event.Venue,
		event.City,
		event.StartsAt,
		event.ImageURL,
	).Scan(&event.ID)
}

func (r Repository) List(ctx context.Context) ([]Event, error) {
	rows, err := r.DB.QueryContext(ctx, `
		SELECT id, title, venue, city, starts_at, image_url, description
		FROM events
		ORDER BY starts_at
		`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event

		if err := rows.Scan(
			&event.ID,
			&event.Title,
			&event.Venue,
			&event.City,
			&event.StartsAt,
			&event.ImageURL,
			&event.Description,
		); err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, rows.Err()
}

func (r Repository) Get(ctx context.Context, id int64) (Event, error) {
	var event Event

	err := r.DB.QueryRowContext(ctx, `
		SELECT id, title, venue, city, starts_at, image_url, description
		FROM events
		WHERE id = $1
		`, id).Scan(
		&event.ID,
		&event.Title,
		&event.Venue,
		&event.City,
		&event.StartsAt,
		&event.ImageURL,
		&event.Description)

	if err != nil {
		return Event{}, err
	}

	return event, nil
}
