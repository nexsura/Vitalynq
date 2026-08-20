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
		t.Fatalf("NewSQLiteMeasurementStore() = nil, want store")
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

	db, err := openSQLite(":memory:")
	if err != nil {
		t.Fatalf("openSQLite() error = %v, want nil", err)
	}

	if err := initializeSQLiteSchema(db); err != nil {
		t.Fatalf("initializeSQLiteSchema() error = %v, want nil", err)
	}

	return db, NewSQLiteMeasurementStore(db)
}

func TestSQLiteMeasurementStoreSavesMeasurement(t *testing.T) {
	db, store := newTestSQLiteMeasurementStore(t)
	defer db.Close()

	saved, err := store.Save(validMeasurement())
	if err != nil {
		t.Fatalf("Save() error = %v, want nil", err)
	}

	if saved.ID != 1 {
		t.Fatalf("ID = %d, want 1", saved.ID)
	}
}

func TestSQLiteMeasurementStoreRejectsInvalidMeasurement(t *testing.T) {
	db, store := newTestSQLiteMeasurementStore(t)
	defer db.Close()

	_, err := store.Save(Measurement{})
	if err == nil {
		t.Fatalf("Save() error = nil, want error")
	}
}

func TestSQLiteMeasurementStoreListsSavedMeasurements(t *testing.T) {
	db, store := newTestSQLiteMeasurementStore(t)
	defer db.Close()

	first, err := store.Save(validMeasurement())
	if err != nil {
		t.Fatalf("Save(first) error = %v, want nil", err)
	}

	secondMeasurement := validMeasurement()
	secondMeasurement.Indicator = "taille"
	secondMeasurement.Value = 175
	secondMeasurement.Unit = "cm"

	second, err := store.Save(secondMeasurement)
	if err != nil {
		t.Fatalf("Save(second) error = %v, want nil", err)
	}

	measurements, err := store.List()
	if err != nil {
		t.Fatalf("List() error = %v, want nil", err)
	}

	if len(measurements) != 2 {
		t.Fatalf("len(List()) = %d, want 2", len(measurements))
	}

	if measurements[0].ID != first.ID {
		t.Fatalf("first ID = %d, want %d", measurements[0].ID, first.ID)
	}

	if measurements[1].ID != second.ID {
		t.Fatalf("second ID = %d, want %d", measurements[1].ID, second.ID)
	}

	if measurements[0].Indicator != "weight" {
		t.Fatalf("first indicator = %q, want %q", measurements[0].Indicator, "weight")
	}

	if measurements[1].Indicator != "taille" {
		t.Fatalf("second indicator = %q, want %q", measurements[1].Indicator, "taille")
	}
}
