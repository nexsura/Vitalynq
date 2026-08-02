package main

import (
	"database/sql"
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
	rows, err := store.db.Query(
		`SELECT id, occurred_at, created_at, indicator, value, unit, context, method, source
  FROM measurements
  ORDER BY occurred_at ASC, id ASC`,
	)
	if err != nil {
		return nil, fmt.Errorf("list sqlite measurements: %w", err)
	}
	defer rows.Close()

	var measurements []Measurement

	for rows.Next() {
		var measurement Measurement
		var occuredAt string
		var createAt string

		if err := rows.Scan(
			&measurement.ID,
			&occuredAt,
			&createAt,
			&measurement.Indicator,
			&measurement.Value,
			&measurement.Unit,
			&measurement.Context,
			&measurement.Method,
			&measurement.Source,
		); err != nil {
			return nil, fmt.Errorf("scan sqlite measurement: %w", err)
		}
		parseOccurredAt, err := time.Parse(time.RFC3339, occuredAt)
		if err != nil {
			return nil, fmt.Errorf("parse sqlite measurement occurred_at: %w", err)
		}

		parseCreateAt, err := time.Parse(time.RFC3339, createAt)
		if err != nil {
			return nil, fmt.Errorf("parse sqlite measurement create_at: %w", err)
		}

		measurement.OccurredAt = parseOccurredAt
		measurement.CreatedAt = parseCreateAt

		measurements = append(measurements, measurement)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate sqlite measurements: %w", err)
	}

	return measurements, nil
}
