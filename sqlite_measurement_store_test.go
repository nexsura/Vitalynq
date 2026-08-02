package main

import (
	"database/sql"
	"testing"
)

func TestNewSQLiteMeasurementStore(t *testing.T) {
	db, err := openSQLite(":memory:")
	if err != nil {
		t.Fatalf("openSQLite() error = %v, want nil", err)
	}
	defer db.Close()

	store := NewSQLiteMeasurementStore(db)
	if store == nil {
		t.Fatalf("NewSQLiteMeasurementsStore() = nil, want store")
	}

	if store.db != db {
		t.Fatalf("store db does not match input db")
	}
}

func TestSQLiteMeasurementStoreImplementsMeasurementStore(t *testing.T) {
	var _ MeasurementStore = (*SQLiteMeasurementStore)(nil)
}

func newTestSQLiteMeasurementStore(t *testing.T) (*sql.DB, *SQLiteMeasurementStore) {
	t.Helper()

	db, err := openSQLite("memory")
	if err != nil {
		t.Fatalf("openSQLite() error = %v, want nil", err)
	}

	if err := initializeSQLiteSchema(db); err != nil {
		t.Fatalf("initializeSQLiteSchema() error = %v, want nil", err)
	}

	return db, NewSQLiteMeasurementStore(db)
}
