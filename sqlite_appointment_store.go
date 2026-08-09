package main

import (
	"database/sql"
	"errors"
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
	return Appointment{}, errors.New("sqlite appointment save is not implemented")
}

func (store *SQLiteAppointmentStore) List() ([]Appointment, error) {
	return nil, errors.New("sqlite appointment list is not implemented")
}
