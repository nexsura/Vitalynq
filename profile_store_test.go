package main

import (
	"testing"
)

func TestMemoryProfileStoreSavesProfileWithID(t *testing.T) {
	store := NewMemoryMedicalProfileStore()

	saved, err := store.Save(validMedicalProfile("Fictive test profile"))
	if err != nil {
		t.Fatalf("Save() error = %v, want nil", err)
	}

	if saved.ID != 1 {
		t.Fatalf("ID = %d, want 1", saved.ID)
	}
}

func TestMemoryMedicalProfileStoreRejectsInvalidProfile(t *testing.T) {
	store := NewMemoryMedicalProfileStore()

	_, err := store.Save(MedicalProfile{})
	if err == nil {
		t.Fatalf("Save() error = nil, want error")
	}
}

func TestMemoryMedicalProfileStoreGetWithoutProfile(t *testing.T) {
	store := NewMemoryMedicalProfileStore()

	_, found, err := store.Get()
	if err != nil {
		t.Fatalf("Get() error = %v, want nil", err)
	}

	if found {
		t.Fatalf("found = true, want false")
	}
}

func TestMemoryMedicalProfileStoreGetSavedProfile(t *testing.T) {
	store := NewMemoryMedicalProfileStore()

	saved, err := store.Save(validMedicalProfile("Fictive test profile"))
	if err != nil {
		t.Fatalf("Save() error = %v, want nil", err)
	}

	got, found, err := store.Get()
	if err != nil {
		t.Fatalf("Get() error = %v, want nil", err)
	}

	if !found {
		t.Fatalf("found = false, want true")
	}

	if got.ID != saved.ID {
		t.Fatalf("ID = %d, want %d", got.ID, saved.ID)
	}

	if got.Label != "Fictive test profile" {
		t.Fatalf("Label = %q, want %q", got.Label, "Fictive test profile")
	}
}

func validMedicalProfile(label string) MedicalProfile {
	return MedicalProfile{
		CreatedAt: testTime(),
		UpdatedAt: testTime(),
		Label:     label,
	}
}
