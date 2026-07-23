package main

import (
	"database/sql"
	"fmt"
	"time"
)

type SQLiteObservationStore struct {
	db *sql.DB
}

func NewSQLiteObservationStore(db *sql.DB) *SQLiteObservationStore {
	return &SQLiteObservationStore{
		db: db,
	}
}

func (store *SQLiteObservationStore) Save(observation Observation) (Observation, error) {
	if err := validateObservation(observation); err != nil {
		return Observation{}, fmt.Errorf("save sqlite observation: %w", err)
	}

	result, err := store.db.Exec(
		`INSERT INTO observations (occurred_at, created_at, text, source)
VALUES (?, ?, ?, ?)`,
		observation.OccurredAt.UTC().Format(time.RFC3339),
		observation.CreatedAt.UTC().Format(time.RFC3339),
		observation.Text,
		observation.Source,
	)
	if err != nil {
		return Observation{}, fmt.Errorf("insert sqlite observation: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return Observation{}, fmt.Errorf("get sqlite observation id: %w", err)
	}

	observation.ID = id

	return observation, nil
}

func (store *SQLiteObservationStore) List() []Observation {
	rows, err := store.db.Query(
		`SELECT id, occurred_at, created_at, text, source
FROM observations
ORDER BY occurred_at ASC, id ASC`,
	)
	if err != nil {
		return nil
	}
	defer rows.Close()

	var observations []Observation

	for rows.Next() {
		var observation Observation
		var occurredAt string
		var createdAt string

		if err := rows.Scan(
			&observation.ID,
			&occurredAt,
			&createdAt,
			&observation.Text,
			&observation.Source,
		); err != nil {
			return nil
		}

		parsedOccurredAt, err := time.Parse(time.RFC3339, occurredAt)
		if err != nil {
			return nil
		}

		parsedCreatedAt, err := time.Parse(time.RFC3339, createdAt)
		if err != nil {
			return nil
		}

		observation.OccurredAt = parsedOccurredAt
		observation.CreatedAt = parsedCreatedAt

		observations = append(observations, observation)
	}

	if err := rows.Err(); err != nil {
		return nil
	}

	return observations
}
