package main

import (
	"errors"
	"strings"
	"time"
)

type MedicalProfile struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Label     string    `json:"label"`
}

func validateMedicalProfile(profile MedicalProfile) error {
	if profile.CreatedAt.IsZero() {
		return errors.New("medical profile creation date is required")
	}

	if profile.UpdatedAt.IsZero() {
		return errors.New("medical profile update date is required")
	}

	if strings.TrimSpace(profile.Label) == "" {
		return errors.New("medical profile label is required")
	}

	return nil
}

func newMedicalProfile(label string) (MedicalProfile, error) {
	now := time.Now().UTC()

	profile := MedicalProfile{
		CreatedAt: now,
		UpdatedAt: now,
		Label:     strings.TrimSpace(label),
	}

	if err := validateMedicalProfile(profile); err != nil {
		return MedicalProfile{}, err
	}

	return profile, nil
}
