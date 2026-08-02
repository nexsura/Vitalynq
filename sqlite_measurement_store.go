package main

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
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
	if err := validateMeasurement(measurement); err != nil {
		return Measurement{}, fmt.Errorf("save sqlite measurement: %w", err)
	}

	result, err := store.db.Exec(
		`INSERT INTO measurements (
    occurred_at,
    created_at,
    indicator,
    value,
    unit,
    context,
    method,
    source
  )
  VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		measurement.OccurredAt.UTC().Format(time.RFC3339),
		measurement.CreatedAt.UTC().Format(time.RFC3339),
		measurement.Indicator,
		measurement.Value,
		measurement.Unit,
		measurement.Context,
		measurement.Method,
		measurement.Source,
	)
	if err != nil {
		return Measurement{}, fmt.Errorf("insert sqlite measurement: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return Measurement{}, fmt.Errorf("get sqlite measurement id: %w", err)
	}

	measurement.ID = id

	return measurement, nil
}

func (store *SQLiteMeasurementStore) List() ([]Measurement, error) {
	return nil, errors.New("sqlite measurement list is not implemented")
}
