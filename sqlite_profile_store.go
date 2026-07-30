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
	row := store.db.QueryRow(
		`SELECT id, created_at, updated_at, label
  FROM medical_profiles
  ORDER BY id DESC
  LIMIT 1`,
	)

	var profile MedicalProfile
	var createdAt string
	var updatedAt string

	if err := row.Scan(&profile.ID, &createdAt, &updatedAt, &profile.Label); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return MedicalProfile{}, false, nil
		}

		return MedicalProfile{}, false, fmt.Errorf("get sqlite medical profile: %w", err)
	}
	parsedCreatedAt, err := time.Parse(time.RFC3339, createdAt)
	if err != nil {
		return MedicalProfile{}, false, fmt.Errorf("parse sqlite medical profile created_at: %w", err)
	}

	parsedUpdatedAt, err := time.Parse(time.RFC3339, updatedAt)
	if err != nil {
		return MedicalProfile{}, false, fmt.Errorf("parse sqlite medical profile updated_at: %w", err)
	}

	profile.CreatedAt = parsedCreatedAt
	profile.UpdatedAt = parsedUpdatedAt

	return profile, true, nil
}
