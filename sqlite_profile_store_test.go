package main

import (
	"database/sql"
	"testing"
)

func TestNewSQLiteMedicalProfileStore(t *testing.T) {
	db, err := openSQLite(":memory:")
	if err != nil {
		t.Fatalf("openSQLite() error = %v, want nil", err)
	}
	defer db.Close()

	store := NewSQLiteMedicalProfileStore(db)
	if store == nil {
		t.Fatalf("NewSQLiteMedicalProfileStore() = nil, want store")
	}

	if store.db != db {
		t.Fatalf("store db does not match input db")
	}
}

func TestSQLiteMedicalProfileStoreImplementsMedicalProfileStore(t *testing.T) {
	var _ MedicalProfileStore = (*SQLiteMedicalProfileStore)(nil)
}

func newTestSQLiteMedicalProfileStore(t *testing.T) (*sql.DB, *SQLiteMedicalProfileStore) {
	t.Helper()

	db, err := openSQLite(":memory:")
	if err != nil {
		t.Fatalf("openSQLite() error = %v, want nil", err)
	}

	if err := initializeSQLiteSchema(db); err != nil {
		t.Fatalf("initializeSQLiteSchema() error = %v, want nil", err)
	}

	return db, NewSQLiteMedicalProfileStore(db)
}

func TestSQLiteMedicalProfileStoreSavesProfile(t *testing.T) {
	db, store := newTestSQLiteMedicalProfileStore(t)
	defer db.Close()

	profile := validMedicalProfile("Profile fictif de test")

	saved, err := store.Save(profile)
	if err != nil {
		t.Fatalf("Save() error = %v, want nil", err)
	}

	if saved.ID != 1 {
		t.Fatalf("ID = %d, want 1", saved.ID)
	}
}

func TestSQLiteMedicalProfileStoreRejectsInvalidProfile(t *testing.T) {
	db, store := newTestSQLiteMedicalProfileStore(t)
	defer db.Close()

	_, err := store.Save(MedicalProfile{})
	if err == nil {
		t.Fatalf("Save() error = nil, want error")
	}
}

func TestSQLiteMedicalProfileStoreGetWithoutProfile(t *testing.T) {
	db, store := newTestSQLiteMedicalProfileStore(t)
	defer db.Close()

	_, found, err := store.Get()
	if err != nil {
		t.Fatalf("Get() error = %v, want nil", err)
	}

	if found {
		t.Fatalf("found = true, want false")
	}
}

func TestSQLiteMedicalProfileStoreGetSavedProfile(t *testing.T) {
	db, store := newTestSQLiteMedicalProfileStore(t)
	defer db.Close()

	saved, err := store.Save(validMedicalProfile("Profil fictif de test"))
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

	if got.Label != "Profil fictif de test" {
		t.Fatalf("Label = %q, want %q", got.Label, "Profil fictif de test")
	}
}
