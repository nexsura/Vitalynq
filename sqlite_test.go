package main

import (
	"testing"
)

func TestOpenSQLite(t *testing.T) {
	db, err := openSQLite(":memory:")
	if err != nil {
		t.Fatalf("openSQLite() error = %v, want nil", err)
	}

	if err := db.Close(); err != nil {
		t.Fatalf("Close() error = %v, want nil", err)
	}
}

func TestInitializeSQLiteSchemaCreatesObservationsTable(t *testing.T) {
	db, err := openSQLite(":memory:")
	if err != nil {
		t.Fatalf("openSQLite() error = %v, want nil", err)
	}
	defer db.Close()

	if err := initializeSQLiteSchema(db); err != nil {
		t.Fatalf("initializeSQLiteSchema() error = %v, want nil", err)
	}

	var tableName string
	err = db.QueryRow(
		"SELECT name FROM sqlite_master WHERE type = 'table' AND name = 'observations'",
	).Scan(&tableName)
	if err != nil {
		t.Fatalf("query observations table error = %v, want nil", err)
	}

	if tableName != "observations" {
		t.Fatalf("tableName = %q, want %q", tableName, "observations")
	}
}

func TestInitializeSQLiteSchemaCreatesMedicalProfilesTable(t *testing.T) {
	db, err := openSQLite(":memory:")
	if err != nil {
		t.Fatalf("openSQLite() error %v, want nil", err)
	}
	defer db.Close()

	if err := initializeSQLiteSchema(db); err != nil {
		t.Fatalf("initializeSQLiteSchema() error = %v, want nil", err)
	}

	var tableName string
	err = db.QueryRow(
		"SELECT name FROM sqlite_master WHERE type = 'table' AND name = 'medical_profiles'",
	).Scan(&tableName)
	if err != nil {
		t.Fatalf("query medical_profiles table error = %v, want nil", err)
	}

	if tableName != "medical_profiles" {
		t.Fatalf("tablename = %q, want %q", tableName, "medical_profiles")
	}
}

func TestInitializeSQLiteSchemaCreatesMeasurementTable(t *testing.T) {
	db, err := openSQLite(":memory:")
	if err != nil {
		t.Fatalf("openSQLite() error = %v, want nil", err)
	}
	defer db.Close()

	if err := initializeSQLiteSchema(db); err != nil {
		t.Fatalf("initializeSQLiteSchema() error = %v, want nil", err)
	}

	var tableName string
	err = db.QueryRow(
		"SELECT name FROM sqlite_master WHERE type = 'table' AND name = 'measurements'",
	).Scan(&tableName)
	if err != nil {
		t.Fatalf("query measurements table error = %v, want nil", err)
	}

	if tableName != "measurements" {
		t.Fatalf("tablename = %q, want %q", tableName, "measurements")
	}
}

func TestInitializeSQLiteSchemaCreatesAppointmentsTable(t *testing.T) {
	db, err := openSQLite(":memory:")
	if err != nil {
		t.Fatalf("openSQLite() error = %v, want nil", err)
	}
	defer db.Close()

	if err := initializeSQLiteSchema(db); err != nil {
		t.Fatalf("initializeSQLiteSchema() error = %v, want nil", err)
	}

	var tablename string
	err = db.QueryRow(
		"SELECT name FROM sqlite_master WHERE type = 'table' AND name = 'appointments'",
	).Scan(&tablename)
	if err != nil {
		t.Fatalf("query appointments table error = %v, want nil", err)
	}

	if tablename != "appointments" {
		t.Fatalf("tablename = %q, want %q", tablename, "appointments")
	}
}
