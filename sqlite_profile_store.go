package main

import (
	"database/sql"
	"errors"
)

type SQLiteMedicalProfileStore struct {
	db *sql.DB
}

func NewSQLiteMedicalProfileStore(db *sql.DB) *SQLiteMedicalProfileStore {
	return &SQLiteMedicalProfileStore{
		db: db,
	}
}

func (store *SQLiteMedicalProfileStore) Save(profile MedicalProfile) (MedicalProfile, error) {
	return MedicalProfile{}, errors.New("sqlite medical profile save is not implemented")
}

func (store *SQLiteMedicalProfileStore) Get() (MedicalProfile, bool, error) {
	return MedicalProfile{}, false, errors.New("sqlite medical profile get is not implemented")
}
