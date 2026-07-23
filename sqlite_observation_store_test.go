package main

import (
	"database/sql"
	"testing"
)

func TestNewSQLiteObservationStore(t *testing.T) {
	db, err := openSQLite(":memory:")
	if err != nil {
		t.Fatalf("openSQLite() error = %v, want nil", err)
	}
	defer db.Close()

	store := NewSQLiteObservationStore(db)
	if store == nil {
		t.Fatalf("NewSQLiteObservationStore() = nil, want store")
	}

	if store.db != db {
		t.Fatalf("store db does not match input db")
	}
}

func TestSQLiteObservationStoreImplementsObservationStore(t *testing.T) {
	var _ ObservationStore = (*SQLiteObservationStore)(nil)
}

func newTestSQLiteObservationStore(t *testing.T) (*sql.DB, *SQLiteObservationStore) {
	t.Helper()
	db, err := openSQLite(":memory:")
	if err != nil {
		t.Fatalf("openSQLite() error = %v, want nil", err)
	}

	if err := initializeSQLiteSchema(db); err != nil {
		t.Fatalf("initializeSQLiteSchema() error = %v, want nil", err)
	}

	return db, NewSQLiteObservationStore(db)
}

func TestSQLiteObservationStoreSavesObservation(t *testing.T) {
	db, store := newTestSQLiteObservationStore(t)
	defer db.Close()

	observation := validStoreObservation("Observation fictive de test")

	saved, err := store.Save(observation)
	if err != nil {
		t.Fatalf("Save() error = %v, want nil", err)
	}

	if saved.ID != 1 {
		t.Fatalf("ID = %d, want 1", saved.ID)
	}
}

func TestSQLiteObservationStoreRejectsInvalidObservation(t *testing.T) {
	db, store := newTestSQLiteObservationStore(t)
	defer db.Close()

	observation := Observation{
		OccurredAt: testTime(),
		CreatedAt:  testTime(),
		Text:       "",
		Source:     "saisie manuelle",
	}

	if _, err := store.Save(observation); err == nil {
		t.Fatalf("Save() error = nil, want error")
	}
}
