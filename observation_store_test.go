package main

import (
	"testing"
	"time"
)

func TestMemoryObservationStoreSavesObservationWithID(t *testing.T) {
	store := NewMemoryObservationStore()

	observation := Observation{
		OccurredAt: testTime(),
		CreatedAt:  testTime(),
		Text:       "Observation fictive de test",
		Source:     "saisie manuelle",
	}

	saved, err := store.Save(observation)
	if err != nil {
		t.Fatalf("Save() error = %v, want nil", err)
	}

	if saved.ID != 1 {
		t.Fatalf("ID = %d, want 1", saved.ID)
	}
}

func TestMemoryObservationStoreIncrementsIDs(t *testing.T) {
	store := NewMemoryObservationStore()

	first := validStoreObservation("Première observation fictive")
	second := validStoreObservation("Deuxième observation fictive")

	savedFirst, err := store.Save(first)
	if err != nil {
		t.Fatalf("Save(first) error = %v, want nil", err)
	}

	savedSecond, err := store.Save(second)
	if err != nil {
		t.Fatalf("Save(second) error = %v, want nil", err)
	}

	if savedFirst.ID != 1 {
		t.Fatalf("first ID = %d, want 1", savedFirst.ID)
	}

	if savedSecond.ID != 2 {
		t.Fatalf("second ID = %d, want 2", savedSecond.ID)
	}
}

func TestMemoryObservationStoreRejectsInvalidObservation(t *testing.T) {
	store := NewMemoryObservationStore()

	observation := Observation{
		OccurredAt: time.Time{},
		CreatedAt:  testTime(),
		Text:       "Observation fictive de test",
		Source:     "saisie manuelle",
	}

	if _, err := store.Save(observation); err == nil {
		t.Fatalf("Save() error = nil, want error")
	}
}

func TestMemoryObservationStoreListsSavedObservations(t *testing.T) {
	store := NewMemoryObservationStore()

	first := validStoreObservation("Première observation fictive")
	second := validStoreObservation("Deuxième observation fictive")

	if _, err := store.Save(first); err != nil {
		t.Fatalf("Save(first) error = %v, want nil", err)
	}

	if _, err := store.Save(second); err != nil {
		t.Fatalf("Save(second) error = %v, want nil", err)
	}

	observations := store.List()

	if len(observations) != 2 {
		t.Fatalf("len(List()) = %d, want 2", len(observations))
	}

	if observations[0].ID != 1 {
		t.Fatalf("first ID = %d, want 1", observations[0].ID)
	}

	if observations[1].ID != 2 {
		t.Fatalf("second ID = %d, want 2", observations[1].ID)
	}
}

func TestMemoryObservationStoreListReturnsCopy(t *testing.T) {
	store := NewMemoryObservationStore()

	if _, err := store.Save(validStoreObservation("Observation fictive de test")); err != nil {
		t.Fatalf("Save() error = %v, want nil", err)
	}

	observations := store.List()
	observations[0].Text = "Texte modifié hors du store"

	observationsAgain := store.List()

	if observationsAgain[0].Text != "Observation fictive de test" {
		t.Fatalf("stored Text = %q, want %q", observationsAgain[0].Text, "Observation fictive de test")
	}
}

func validStoreObservation(text string) Observation {
	return Observation{
		OccurredAt: testTime(),
		CreatedAt:  testTime(),
		Text:       text,
		Source:     "saisie manuelle",
	}
}
