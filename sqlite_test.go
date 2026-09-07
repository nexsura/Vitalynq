package main

import (
	"database/sql"
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

func TestInitializeSQLiteSchemaCreatesSchemaMigrationsTable(t *testing.T) {
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
		"SELECT name FROM sqlite_master WHERE type = 'table' AND name = 'schema_migrations'",
	).Scan(&tableName)
	if err != nil {
		t.Fatalf("query schema_migrations table error = %v, want nil", err)
	}

	if tableName != "schema_migrations" {
		t.Fatalf("tableName = %q, want %q", tableName, "schema_migrations")
	}
}

func TestAppliedSQLiteMigrationVersionsReturnsEmptyList(t *testing.T) {
	db, err := openSQLite(":memory:")
	if err != nil {
		t.Fatalf("openSQLite() error = %v, want nil", err)
	}
	defer db.Close()

	if err := initializeSQLiteSchema(db); err != nil {
		t.Fatalf("initializeSQLiteSchema() error = %v, want nil", err)
	}

	versions, err := appliedSQLiteMigrationVersions(db)
	if err != nil {
		t.Fatalf("appliedSQLiteMigrationVersions() error = %v, want nil", err)
	}

	if len(versions) != 0 {
		t.Fatalf("len(versions) = %d, want 0", len(versions))
	}
}

func TestRecordSQLiteMigrationVersion(t *testing.T) {
	db, err := openSQLite(":memory:")
	if err != nil {
		t.Fatalf("openSQLite() error = %v, want nil", err)
	}
	defer db.Close()

	if err := initializeSQLiteSchema(db); err != nil {
		t.Fatalf("initializeSQLiteSchema() error = %v, want nil", err)
	}

	if err := recordSQLiteMigrationVersion(db, 1, testTime()); err != nil {
		t.Fatalf("recordSQLiteMigrationVersion() error = %v, want nil", err)
	}

	versions, err := appliedSQLiteMigrationVersions(db)
	if err != nil {
		t.Fatalf("appliedSQLiteMigrationVersions() error = %v, want nil", err)
	}

	if len(versions) != 1 {
		t.Fatalf("len(versions) = %d, want 1", len(versions))
	}

	if versions[0] != 1 {
		t.Fatalf("version[0] = %d, want 1", versions[0])
	}
}

func TestSQLiteMigrationStoresVersionAndApplyFunction(t *testing.T) {
	called := false

	migration := SQLiteMigration{
		Version: 1,
		Apply: func(tx *sql.Tx) error {
			called = true
			return nil
		},
	}

	if migration.Version != 1 {
		t.Fatalf("Version = %d, want 1", migration.Version)
	}

	if migration.Apply == nil {
		t.Fatalf("Apply = nil, want function")
	}

	db, err := openSQLite(":memory:")
	if err != nil {
		t.Fatalf("openSQLite() error = %v, want nil", err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		t.Fatalf("Begin() error = %v, want nil", err)
	}
	defer tx.Rollback()

	if err := migration.Apply(tx); err != nil {
		t.Fatalf("Apply() error = %v, want nil", err)
	}

	if !called {
		t.Fatalf("called = false, want true")
	}
}

func TestHasSQLiteMigrationVersionReturnsFalseWhenMissing(t *testing.T) {
	db, err := openSQLite(":memory:")
	if err != nil {
		t.Fatalf("openSQLite() error = %v, want nil", err)
	}
	defer db.Close()

	if err := initializeSQLiteSchema(db); err != nil {
		t.Fatalf("initializeSQLiteSchema() error = %v, want nil", err)
	}

	found, err := hasSQLiteMigrationVersion(db, 1)
	if err != nil {
		t.Fatalf("hasSQLiteMigrationVersion() error = %v, want nil", err)
	}

	if found {
		t.Fatalf("found = true, want false")
	}
}

func TestHasSQLiteMigrationVersionReturnsTrueWhenPresent(t *testing.T) {
	db, err := openSQLite(":memory:")
	if err != nil {
		t.Fatalf("openSQLite() error = %v, want nil", err)
	}
	defer db.Close()

	if err := initializeSQLiteSchema(db); err != nil {
		t.Fatalf("initializeSQLiteSchema() error = %v, want nil", err)
	}

	if err := recordSQLiteMigrationVersion(db, 1, testTime()); err != nil {
		t.Fatalf("recordSQLiteMigrationVersion() error = %v, want nil", err)
	}

	found, err := hasSQLiteMigrationVersion(db, 1)
	if err != nil {
		t.Fatalf("hasSQLiteMigrationVersion() error = %v, want nil", err)
	}

	if !found {
		t.Fatalf("found = false, want true")
	}
}
