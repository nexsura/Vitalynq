package main

import (
	"database/sql"
	"fmt"
	"testing"
	"time"
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

	tx, err := db.Begin()
	if err != nil {
		t.Fatalf("Begin() error = %v, want nil", err)
	}

	if err := recordSQLiteMigrationVersion(tx, 1, testTime()); err != nil {
		t.Fatalf("recordSQLiteMigrationVersion() error = %v, want nil", err)
	}

	if err := tx.Commit(); err != nil {
		t.Fatalf("Commit() error = %v, want nil", err)
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

	tx, err := db.Begin()
	if err != nil {
		t.Fatalf("Begin() error = %v, want nil", err)
	}

	if err := recordSQLiteMigrationVersion(tx, 1, testTime()); err != nil {
		t.Fatalf("recordSQLiteMigrationVersion() error = %v, want nil", err)
	}

	if err := tx.Commit(); err != nil {
		t.Fatalf("Commit() error = %v, want nil", err)
	}

	found, err := hasSQLiteMigrationVersion(db, 1)
	if err != nil {
		t.Fatalf("hasSQLiteMigrationVersion() error = %v, want nil", err)
	}

	if !found {
		t.Fatalf("found = false, want true")
	}
}

func TestApplySQLiteMigrationAppliesMigration(t *testing.T) {
	db, err := openSQLite(":memory:")
	if err != nil {
		t.Fatalf("openSQLite() error = %v, want nil", err)
	}
	defer db.Close()

	if err := initializeSQLiteSchema(db); err != nil {
		t.Fatalf("initializeSQLiteSchema() error = %v, want nil", err)
	}

	called := false
	migration := SQLiteMigration{
		Version: 1,
		Apply: func(tx *sql.Tx) error {
			called = true
			_, err := tx.Exec("CREATE TABLE migration_test (id INTEGER PRIMARY KEY)")
			return err
		},
	}

	if err := applySQLiteMigration(db, migration, testTime()); err != nil {
		t.Fatalf("applySQLiteMigration() error = %v, want nil", err)
	}

	if !called {
		t.Fatalf("called = false , want true")
	}

	found, err := hasSQLiteMigrationVersion(db, 1)
	if err != nil {
		t.Fatalf("hasSQLiteMigrationVersion() error = %v, want nil", err)
	}

	if !found {
		t.Fatalf("found = false, want true")
	}
}

func TestApplySQLiteMigrationSkipsAppliedMigration(t *testing.T) {
	db, err := openSQLite(":memory:")
	if err != nil {
		t.Fatalf("openSQLite() error = %v, want nil", err)
	}
	defer db.Close()

	if err := initializeSQLiteSchema(db); err != nil {
		t.Fatalf("initializeSQLiteSchema() error = %v, want nil", err)
	}

	tx, err := db.Begin()
	if err != nil {
		t.Fatalf("Begin() error = %v, want nil", err)
	}

	if err := recordSQLiteMigrationVersion(tx, 1, testTime()); err != nil {
		t.Fatalf("recordSQLiteMigrationVersion() error = %v, want nil", err)
	}

	if err := tx.Commit(); err != nil {
		t.Fatalf("Commit() error = %v, want nil", err)
	}

	called := false
	migration := SQLiteMigration{
		Version: 1,
		Apply: func(tx *sql.Tx) error {
			called = true
			return nil
		},
	}

	if err := applySQLiteMigration(db, migration, testTime()); err != nil {
		t.Fatalf("applySQLiteMigration() error = %v, want nil", err)
	}

	if called {
		t.Fatalf("called = true, want false")
	}
}

func TestApplySQLiteMigrationRejectsInvalidVersion(t *testing.T) {
	db, err := openSQLite(":memory:")
	if err != nil {
		t.Fatalf("openSQLite() error = %v, want nil", err)
	}
	defer db.Close()

	migration := SQLiteMigration{
		Version: 0,
		Apply: func(tx *sql.Tx) error {
			return nil
		},
	}

	if err := applySQLiteMigration(db, migration, testTime()); err == nil {
		t.Fatalf("applySQLiteMigration() error = nil, want error")
	}
}

func TestApplySQLiteMigrationRejectsMissingApplyFunction(t *testing.T) {
	db, err := openSQLite(":memory:")
	if err != nil {
		t.Fatalf("openSQLite() error = %v, want nil", err)
	}
	defer db.Close()

	migration := SQLiteMigration{
		Version: 1,
	}

	if err := applySQLiteMigration(db, migration, testTime()); err == nil {
		t.Fatalf("applySQLiteMigration() error = nil, want error")
	}
}

func TestApplySQLiteMigrationRejectsMissingAppliedDate(t *testing.T) {
	db, err := openSQLite(":memory:")
	if err != nil {
		t.Fatalf("openSQLite() error = %v, want nil", err)
	}
	defer db.Close()

	migration := SQLiteMigration{
		Version: 1,
		Apply: func(tx *sql.Tx) error {
			return nil
		},
	}

	err = applySQLiteMigration(db, migration, time.Time{})
	if err == nil {
		t.Fatalf("applySQLiteMigration() error = nil, want error")
	}

	want := "sqlite migration applied date is required"
	if err.Error() != want {
		t.Fatalf("applySQLiteMigration() error = %q, want %q", err.Error(), want)
	}
}

func TestApplySQLiteMigrationsAppliesMigrationsInVersionOrder(t *testing.T) {
	db, err := openSQLite(":memory:")
	if err != nil {
		t.Fatalf("openSQLite() error = %v, want nil", err)
	}
	defer db.Close()

	if err := initializeSQLiteSchema(db); err != nil {
		t.Fatalf("initializeSQLiteSchema() error = %v, want nil", err)
	}

	var appliedVersions []int

	migrations := []SQLiteMigration{
		{
			Version: 2,
			Apply: func(tx *sql.Tx) error {
				appliedVersions = append(appliedVersions, 2)
				return nil
			},
		},
		{
			Version: 1,
			Apply: func(tx *sql.Tx) error {
				appliedVersions = append(appliedVersions, 1)
				return nil
			},
		},
	}

	if err := applySQLiteMigrations(db, migrations, testTime()); err != nil {
		t.Fatalf("applySQLiteMigrations() error = %v, want nil", err)
	}

	wantVersions := []int{1, 2}
	if len(appliedVersions) != len(wantVersions) {
		t.Fatalf("len(appliedVersions) = %d, want %d", len(appliedVersions), len(wantVersions))
	}

	for index := range wantVersions {
		if appliedVersions[index] != wantVersions[index] {
			t.Fatalf("appliedVersions[%d] = %d, want %d", index, appliedVersions[index], wantVersions[index])
		}
	}
}

func TestApplySQLiteMigrationsRejectsDuplicateVersion(t *testing.T) {
	db, err := openSQLite(":memory:")
	if err != nil {
		t.Fatalf("openSQLite() error = %v, want nil", err)
	}
	defer db.Close()

	if err := initializeSQLiteSchema(db); err != nil {
		t.Fatalf("initializeSQLiteSchema() error = %v, want nil", err)
	}

	migrations := []SQLiteMigration{
		{
			Version: 1,
			Apply: func(tx *sql.Tx) error {
				return nil
			},
		},
		{
			Version: 1,
			Apply: func(tx *sql.Tx) error {
				return nil
			},
		},
	}

	err = applySQLiteMigrations(db, migrations, testTime())
	if err == nil {
		t.Fatalf("applySQLiteMigrations() error = nil, want error")
	}

	want := "duplicate sqlite migration version: 1"
	if err.Error() != want {
		t.Fatalf("applySQLiteMigrations() error = %q, want %q", err.Error(), want)
	}
}

func TestApplySQLiteMigrationsStopsAfterFailedMigration(t *testing.T) {
	db, err := openSQLite(":memory:")
	if err != nil {
		t.Fatalf("openSQLite() error = %v, want nil", err)
	}
	defer db.Close()

	if err := initializeSQLiteSchema(db); err != nil {
		t.Fatalf("initializeSQLiteSchema() error = %v, want nil", err)
	}

	var appliedVersions []int

	migrations := []SQLiteMigration{
		{
			Version: 1,
			Apply: func(tx *sql.Tx) error {
				appliedVersions = append(appliedVersions, 1)
				return nil
			},
		},
		{
			Version: 2,
			Apply: func(tx *sql.Tx) error {
				appliedVersions = append(appliedVersions, 2)
				return fmt.Errorf("fictive migration failure")
			},
		},
		{
			Version: 3,
			Apply: func(tx *sql.Tx) error {
				appliedVersions = append(appliedVersions, 3)
				return nil
			},
		},
	}

	err = applySQLiteMigrations(db, migrations, testTime())
	if err == nil {
		t.Fatalf("applySQLiteMigrations() error = nil , want error")
	}

	wantVersions := []int{1, 2}
	if len(appliedVersions) != len(wantVersions) {
		t.Fatalf("len(appliedVersions) = %d, want %d", len(appliedVersions), len(wantVersions))
	}

	for index := range wantVersions {
		if appliedVersions[index] != wantVersions[index] {
			t.Fatalf("appliedVersions[%d] = %d, want %d", index, appliedVersions[index], wantVersions[index])
		}
	}
}

func TestApplySQLiteMigrationsDoesNotRecordFailedMigration(t *testing.T) {
	db, err := openSQLite(":memory:")
	if err != nil {
		t.Fatalf("openSQLite() error = %v, want nil", err)
	}
	defer db.Close()

	if err := initializeSQLiteSchema(db); err != nil {
		t.Fatalf("initializeSQLiteSchema() error = %v, want nil", err)
	}

	migration := SQLiteMigration{
		Version: 1,
		Apply: func(tx *sql.Tx) error {
			return fmt.Errorf("fictive migration failure")
		},
	}

	err = applySQLiteMigrations(db, []SQLiteMigration{migration}, testTime())
	if err == nil {
		t.Fatalf("applySQLiteMigrations() error = nil, want error")
	}

	applied, err := hasSQLiteMigrationVersion(db, 1)
	if err != nil {
		t.Fatalf("hasSQLiteMigrationVersion() error = %v, want nil", err)
	}

	if applied {
		t.Fatalf("applied = true, want false")
	}
}

func TestHasCurrentSQLiteSchemaReturnsTrueForInitializedSchema(t *testing.T) {
	db, err := openSQLite(":memory:")
	if err != nil {
		t.Fatalf("openSQLite() error = %v, want nil", err)
	}
	defer db.Close()

	if err := initializeSQLiteSchema(db); err != nil {
		t.Fatalf("initializeSQLiteSchema() error = %v, want nil", err)
	}

	hasSchema, err := hasCurrentSQLiteSchema(db)
	if err != nil {
		t.Fatalf("hasCurrentSQLiteSchema() error = %v, want nil", err)
	}

	if !hasSchema {
		t.Fatalf("hasSchema = false, want true")
	}
}

func TestHasCurrentSQLiteSchemaReturnsFalseForEmptyDatabase(t *testing.T) {
	db, err := openSQLite(":memory:")
	if err != nil {
		t.Fatalf("openSQLite() error = %v, want nil", err)
	}
	defer db.Close()

	hasSchema, err := hasCurrentSQLiteSchema(db)
	if err != nil {
		t.Fatalf("hasCurrentSQLiteSchema() error = %v, want nil", err)
	}

	if hasSchema {
		t.Fatalf("hasSchemma = true, want false")
	}
}

func TestMarkCurrentSQLiteSchemaBaselineRecordsVersionOne(t *testing.T) {
	db, err := openSQLite(":memory:")
	if err != nil {
		t.Fatalf("openSQLite() error = %v, want nil", err)
	}
	defer db.Close()

	if err := initializeSQLiteSchema(db); err != nil {
		t.Fatalf("initializeSQLiteSchema() error = %v, want nil", err)
	}

	if err := markCurrentSQLiteSchemaBaseline(db, testTime()); err != nil {
		t.Fatalf("markCurrentSQLiteSchemaBaseline() error = %v, want nil", err)
	}

	applied, err := hasSQLiteMigrationVersion(db, 1)
	if err != nil {
		t.Fatalf("hasSQLiteMigrationVersion() error = %v, want nil", err)
	}

	if !applied {
		t.Fatalf("applied = false , want true")
	}
}

func TestMarkCurrentSQLiteSchemaBaselineRejectsEmptyDatabase(t *testing.T) {
	db, err := openSQLite(":memory:")
	if err != nil {
		t.Fatalf("openSQLite() error = %v, want nil", err)
	}
	defer db.Close()

	err = markCurrentSQLiteSchemaBaseline(db, testTime())
	if err == nil {
		t.Fatalf("markCurrentSQLiteSchemaBaseline() error = nil, want error")
	}

	want := "sqlite schema does not match current baseline"
	if err.Error() != want {
		t.Fatalf("markCurrentSQLiteSchemaBaseline() error = %q, want %q", err.Error(), want)
	}
}

func TestMarkCurrentSQLiteSchemaBaselineRejectsExistingMigrations(t *testing.T) {
	db, err := openSQLite(":memory:")
	if err != nil {
		t.Fatalf("openSQLite() error = %v, want nil", err)
	}
	defer db.Close()

	if err := initializeSQLiteSchema(db); err != nil {
		t.Fatalf("initializeSQLiteSchema() error = %v, want nil", err)
	}

	tx, err := db.Begin()
	if err != nil {
		t.Fatalf("Begin() error = %v, want nil", err)
	}

	if err := recordSQLiteMigrationVersion(tx, 1, testTime()); err != nil {
		t.Fatalf("recordSQLiteMigrationVersion() error = %v, want nil", err)
	}

	if err := tx.Commit(); err != nil {
		t.Fatalf("Commit() error = %v, want nil", err)
	}

	err = markCurrentSQLiteSchemaBaseline(db, testTime())
	if err == nil {
		t.Fatalf("markCurrentSQLiteSchemaBaseline() error = nil, want error")
	}

	want := "sqlite schema baseline already has applied migrations"
	if err.Error() != want {
		t.Fatalf("markCurrentSQLiteSchemaBaseline() error = %q, want %q", err.Error(), want)
	}
}

func TestMarkCurrentSQLiteSchemaBaselineRejectsMissingAppliedDate(t *testing.T) {
	db, err := openSQLite(":memory:")
	if err != nil {
		t.Fatalf("openSQLite() error = %v, want nil", err)
	}
	defer db.Close()

	if err := initializeSQLiteSchema(db); err != nil {
		t.Fatalf("initializeSQLiteSchema() error = %v, want nil", err)
	}

	err = markCurrentSQLiteSchemaBaseline(db, time.Time{})
	if err == nil {
		t.Fatalf("markCurrentSQLiteSchemaBaseline() error = nil, want error")
	}

	want := "sqlite schema baseline applied date is required"
	if err.Error() != want {
		t.Fatalf("markCurrentSQLiteSchemaBaseline() error = %q, want %q", err.Error(), want)
	}
}

func TestHasCurrentSQLiteSchemaReturnsFalseForMissingColumn(t *testing.T) {
	db, err := openSQLite(":memory:")
	if err != nil {
		t.Fatalf("openSQLite() error = %v, want nil", err)
	}
	defer db.Close()

	statements := []string{
		`CREATE TABLE schema_migrations (
			version INTEGER PRIMARY KEY,
			applied_at TEXT NOT NULL
		);`,
		`CREATE TABLE observations (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			occurred_at TEXT NOT NULL,
			created_at TEXT NOT NULL,
			text TEXT NOT NULL
		);`,
		`CREATE TABLE medical_profiles (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			created_at TEXT NOT NULL,
			updated_at TEXT NOT NULL,
			label TEXT NOT NULL
		);`,
		`CREATE TABLE measurements (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			occurred_at TEXT NOT NULL,
			created_at TEXT NOT NULL,
			indicator TEXT NOT NULL,
			value REAL NOT NULL,
			unit TEXT NOT NULL,
			context TEXT NOT NULL,
			method TEXT NOT NULL,
			source TEXT NOT NULL
		);`,
		`CREATE TABLE appointments (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			scheduled_at TEXT NOT NULL,
			created_at TEXT NOT NULL,
			title TEXT NOT NULL,
			category TEXT NOT NULL,
			location TEXT NOT NULL,
			source TEXT NOT NULL
		);`,
	}

	for _, statement := range statements {
		if _, err := db.Exec(statement); err != nil {
			t.Fatalf("Exec() error = %v, want nil", err)
		}
	}

	hasSchema, err := hasCurrentSQLiteSchema(db)
	if err != nil {
		t.Fatalf("hasCurrentSQLiteSchema() error = %v, want nil", err)
	}

	if hasSchema {
		t.Fatalf("hasSchema = true, want false")
	}
}

func TestHasCurrentSQLiteSchemaReturnsFalseForUnexpectedColumn(t *testing.T) {
	db, err := openSQLite(":memory:")
	if err != nil {
		t.Fatalf("openSQLite() error = %v, want nil", err)
	}
	defer db.Close()

	if err := initializeSQLiteSchema(db); err != nil {
		t.Fatalf("initializeSQLiteSchema() error = %v, want nil", err)
	}

	if _, err := db.Exec("ALTER TABLE observations ADD COLUMN extra TEXT"); err != nil {
		t.Fatalf("Exec() error = %v, want nil", err)
	}

	hasSchema, err := hasCurrentSQLiteSchema(db)
	if err != nil {
		t.Fatalf("hasCurrentSQLiteSchema() error = %v, want nil", err)
	}

	if hasSchema {
		t.Fatalf("hasSchema = true, want false")
	}
}

func TestMarkCurrentSQLiteSchemaBaselineRejectsUnexpectedColumn(t *testing.T) {
	db, err := openSQLite(":memory:")
	if err != nil {
		t.Fatalf("openSQLite() error = %v, want nil", err)
	}
	defer db.Close()

	if err := initializeSQLiteSchema(db); err != nil {
		t.Fatalf("initializeSQLiteSchema() error = %v, want nil", err)
	}

	if _, err := db.Exec("ALTER TABLE observations ADD COLUMN extra TEXT"); err != nil {
		t.Fatalf("Exec() error = %v, want nil", err)
	}

	err = markCurrentSQLiteSchemaBaseline(db, testTime())
	if err == nil {
		t.Fatalf("markCurrentSQLiteSchemaBaseline() error = nil, want error")
	}

	want := "sqlite schema does not match current baseline"
	if err.Error() != want {
		t.Fatalf("markCurrentSQLiteSchemaBaseline() error = %q, want %q", err.Error(), want)
	}

	applied, err := hasSQLiteMigrationVersion(db, 1)
	if err != nil {
		t.Fatalf("hasSQLiteMigrationVersion() error = %v, want nil", err)
	}

	if applied {
		t.Fatalf("applied = true, want false")
	}
}

func TestMarkCurrentSQLiteSchemaBaselineRejectsMissingColumn(t *testing.T) {
	db, err := openSQLite(":memory:")
	if err != nil {
		t.Fatalf("openSQLite() error = %v, want nil", err)
	}
	defer db.Close()

	statements := []string{
		`CREATE TABLE schema_migrations (
			version INTEGER PRIMARY KEY,
			applied_at TEXT NOT NULL
		);`,
		`CREATE TABLE observations (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			occurred_at TEXT NOT NULL,
			created_at TEXT NOT NULL,
			text TEXT NOT NULL
		);`,
		`CREATE TABLE medical_profiles (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			created_at TEXT NOT NULL,
			updated_at TEXT NOT NULL,
			label TEXT NOT NULL
		);`,
		`CREATE TABLE measurements (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			occurred_at TEXT NOT NULL,
			created_at TEXT NOT NULL,
			indicator TEXT NOT NULL,
			value REAL NOT NULL,
			unit TEXT NOT NULL,
			context TEXT NOT NULL,
			method TEXT NOT NULL,
			source TEXT NOT NULL
		);`,
		`CREATE TABLE appointments (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			scheduled_at TEXT NOT NULL,
			created_at TEXT NOT NULL,
			title TEXT NOT NULL,
			category TEXT NOT NULL,
			location TEXT NOT NULL,
			source TEXT NOT NULL
		);`,
	}

	for _, statement := range statements {
		if _, err := db.Exec(statement); err != nil {
			t.Fatalf("Exec() error = %v, want nil", err)
		}
	}

	err = markCurrentSQLiteSchemaBaseline(db, testTime())
	if err == nil {
		t.Fatalf("markCurrentSQLiteSchemaBaseline() error = nil, want error")
	}

	want := "sqlite schema does not match current baseline"
	if err.Error() != want {
		t.Fatalf("markCurrentSQLiteSchemaBaseline() error = %q, want %q", err.Error(), want)
	}

	applied, err := hasSQLiteMigrationVersion(db, 1)
	if err != nil {
		t.Fatalf("hasSQLiteMigrationVersion() error = %v, want nil", err)
	}

	if applied {
		t.Fatalf("applied = true, want false")
	}
}
