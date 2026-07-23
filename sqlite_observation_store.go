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
	return nil
}
