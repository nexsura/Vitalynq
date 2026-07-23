package main

import (
	"database/sql"
	"errors"
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
	return Observation{}, errors.New("sqlite observation save is not implemented")
}

func (store *SQLiteObservationStore) List() []Observation {
	return nil
}
