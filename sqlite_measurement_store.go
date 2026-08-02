package main

import (
	"database/sql"
	"errors"
)

type SQLiteMeasurementStore struct {
	db *sql.DB
}

func NewSQLiteMeasurementStore(db *sql.DB) *SQLiteMeasurementStore {
	return &SQLiteMeasurementStore{
		db: db,
	}
}

func (store *SQLiteMeasurementStore) Save(measurement Measurement) (Measurement, error) {
	return Measurement{}, errors.New("sqlite measurement save it not implemented")
}

func (store *SQLiteMeasurementStore) List() ([]Measurement, error) {
	return nil, errors.New("sqlite measurement list is not implemented")
}
