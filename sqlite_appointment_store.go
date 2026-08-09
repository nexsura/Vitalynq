package main

import (
	"database/sql"
	"fmt"
	"time"
)

type SQLiteAppointmentStore struct {
	db *sql.DB
}

func NewSQLiteAppointmentStore(db *sql.DB) *SQLiteAppointmentStore {
	return &SQLiteAppointmentStore{
		db: db,
	}
}

func (store *SQLiteAppointmentStore) Save(appointment Appointment) (Appointment, error) {
	if err := validateAppointment(appointment); err != nil {
		return Appointment{}, fmt.Errorf("save sqlite appointment: %w", err)
	}

	result, err := store.db.Exec(
		`INSERT INTO appointments (scheduled_at, created_at, title, category, location, source) VALUES (?, ?, ?, ?, ?, ?)`,
		appointment.ScheduledAt.UTC().Format(time.RFC3339),
		appointment.CreatedAt.UTC().Format(time.RFC3339),
		appointment.Title,
		appointment.Category,
		appointment.Location,
		appointment.Source,
	)
	if err != nil {
		return Appointment{}, fmt.Errorf("insert sqlite appointment: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return Appointment{}, fmt.Errorf("get sqlite appointment id: %w", err)
	}

	appointment.ID = id

	return appointment, nil
}

func (store *SQLiteAppointmentStore) List() ([]Appointment, error) {
	rows, err := store.db.Query(
		`SELECT id, scheduled_at, created_at, title, category, location, source FROM appointments ORDER BY scheduled_at ASC, id ASC`,
	)
	if err != nil {
		return nil, fmt.Errorf("list sqlite appointments: %w", err)
	}
	defer rows.Close()

	var appointments []Appointment

	for rows.Next() {
		var appointment Appointment
		var scheduledAt string
		var createdAt string

		if err := rows.Scan(
			&appointment.ID,
			&scheduledAt,
			&createdAt,
			&appointment.Title,
			&appointment.Category,
			&appointment.Location,
			&appointment.Source,
		); err != nil {
			return nil, fmt.Errorf("scan sqlite appointment: %w", err)
		}

		parsedScheduledAt, err := time.Parse(time.RFC3339, scheduledAt)
		if err != nil {
			return nil, fmt.Errorf("parse sqlite appointment scheduled_at %w", err)
		}

		parsedCreatedAt, err := time.Parse(time.RFC3339, createdAt)
		if err != nil {
			return nil, fmt.Errorf("parse sqlite appointment created_at %w", err)
		}

		appointment.ScheduledAt = parsedScheduledAt
		appointment.CreatedAt = parsedCreatedAt

		appointments = append(appointments, appointment)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate sqlite appointments: %w", err)
	}

	return appointments, nil
}
