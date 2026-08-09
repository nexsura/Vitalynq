package main

import (
	"database/sql"
	"errors"
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
		return Appointment{}, fmt.Errorf("get sqlite apointment id: %w", err)
	}

	appointment.ID = id

	return appointment, nil
}

func (store *SQLiteAppointmentStore) List() ([]Appointment, error) {
	return nil, errors.New("sqlite appointment list is not implemented")
}
