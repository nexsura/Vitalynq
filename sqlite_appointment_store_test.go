package main

import (
	"database/sql"
	"testing"
)

func TestNewSQLiteAppointmentStore(t *testing.T) {
	db, err := openSQLite(":memory:")
	if err != nil {
		t.Fatalf("openSQLite() error = %v, want nil", err)
	}
	defer db.Close()

	store := NewSQLiteAppointmentStore(db)
	if store == nil {
		t.Fatalf("NewSQLiteAppointmentStore() = nil, want store")
	}

	if store.db != db {
		t.Fatalf("store db does not match input db")
	}
}

func TestSQLiteAppointmentStoreImplementsAppointmentStore(t *testing.T) {
	var _ AppointmentStore = (*SQLiteAppointmentStore)(nil)
}

func newTestSQLiteAppointmentStore(t *testing.T) (*sql.DB, *SQLiteAppointmentStore) {
	t.Helper()

	db, err := openSQLite(":memory:")
	if err != nil {
		t.Fatalf("openSQLite() error = %v, want nil", err)
	}

	if err := initializeSQLiteSchema(db); err != nil {
		t.Fatalf("initializeSQLiteSchema() error = %v, want nil", err)
	}

	return db, NewSQLiteAppointmentStore(db)
}
