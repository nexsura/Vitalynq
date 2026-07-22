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
