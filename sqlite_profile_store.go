package main

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
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
	if err := validateMedicalProfile(profile); err != nil {
		return MedicalProfile{}, fmt.Errorf("save sqlite medical profile: %w", err)
	}

	result, err := store.db.Exec(
		`INSERT INTO medical_profiles (created_at, updated_at, label)
  VALUES (?, ?, ?)`,
		profile.CreatedAt.UTC().Format(time.RFC3339),
		profile.UpdatedAt.UTC().Format(time.RFC3339),
		profile.Label,
	)
	if err != nil {
		return MedicalProfile{}, fmt.Errorf("insert sqlite medical profile: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return MedicalProfile{}, fmt.Errorf("get sqlite medical profile id: %w", err)
	}

	profile.ID = id

	return profile, nil
}

func (store *SQLiteMedicalProfileStore) Get() (MedicalProfile, bool, error) {
	return MedicalProfile{}, false, errors.New("sqlite medical profile get is not implemented")
}
