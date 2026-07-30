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
