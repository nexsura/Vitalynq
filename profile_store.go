package main

import "fmt"

type MedicalProfileStore interface {
	Save(profile MedicalProfile) (MedicalProfile, error)
	Get() (MedicalProfile, bool, error)
}

type MemoryMedicalProfileStore struct {
	nextID     int64
	profile    MedicalProfile
	hasProfile bool
}

func NewMemoryMedicalProfileStore() *MemoryMedicalProfileStore {
	return &MemoryMedicalProfileStore{
		nextID: 1,
	}
}

func (store *MemoryMedicalProfileStore) Save(profile MedicalProfile) (MedicalProfile, error) {
	if err := validateMedicalProfile(profile); err != nil {
		return MedicalProfile{}, fmt.Errorf("save medical profile: %w", err)
	}

	if profile.ID == 0 {
		profile.ID = store.nextID
		store.nextID++
	}

	store.profile = profile
	store.hasProfile = true

	return profile, nil
}

func (store *MemoryMedicalProfileStore) Get() (MedicalProfile, bool, error) {
	if !store.hasProfile {
		return MedicalProfile{}, false, nil
	}

	return store.profile, true, nil
}
