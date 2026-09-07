package main

import (
	"database/sql"
	"fmt"
	"sort"
	"time"

	_ "modernc.org/sqlite"
)

type SQLiteMigration struct {
	Version int
	Apply   func(*sql.Tx) error
}

func openSQLite(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, fmt.Errorf("open sqlite: %w", err)
	}

	if err := db.Ping(); err != nil {
		if closeErr := db.Close(); closeErr != nil {
			return nil, fmt.Errorf("ping sqlite: %w; close sqlite: %v", err, closeErr)
		}

		return nil, fmt.Errorf("ping sqlite: %w", err)
	}

	return db, nil
}

func initializeSQLiteSchema(db *sql.DB) error {
	statements := []string{
		`CREATE TABLE IF NOT EXISTS schema_migrations (
    version INTEGER PRIMARY KEY,
    applied_at TEXT NOT NULL
    );`,
		`CREATE TABLE IF NOT EXISTS observations (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  occurred_at TEXT NOT NULL,
  created_at TEXT NOT NULL,
  text TEXT NOT NULL,
  source TEXT NOT NULL
  );`,
		`CREATE TABLE IF NOT EXISTS medical_profiles (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  created_at TEXT NOT NULL,
  updated_at TEXT NOT NULL,
  label TEXT NOT NULL
  );`,
		`CREATE TABLE IF NOT EXISTS measurements (
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
		`CREATE TABLE IF NOT EXISTS appointments (
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
			return fmt.Errorf("initialize sqlite schema: %w", err)
		}
	}

	return nil
}

func appliedSQLiteMigrationVersions(db *sql.DB) ([]int, error) {
	rows, err := db.Query("SELECT version FROM schema_migrations ORDER BY version ASC")
	if err != nil {
		return nil, fmt.Errorf("list sqlite migration versions: %w", err)
	}
	defer rows.Close()

	var versions []int

	for rows.Next() {
		var version int
		if err := rows.Scan(&version); err != nil {
			return nil, fmt.Errorf("scan sqlite migration version: %w", err)
		}

		versions = append(versions, version)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate sqlite migration versions: %w", err)
	}

	return versions, nil
}

func recordSQLiteMigrationVersion(tx *sql.Tx, version int, appliedAt time.Time) error {
	if version <= 0 {
		return fmt.Errorf("sqlite migration version must be positive")
	}

	if appliedAt.IsZero() {
		return fmt.Errorf("sqlite migration applied date is required")
	}

	_, err := tx.Exec(
		"INSERT INTO schema_migrations (version, applied_at) VALUES (?, ?)",
		version,
		appliedAt.UTC().Format(time.RFC3339),
	)
	if err != nil {
		return fmt.Errorf("record sqlite migration version: %w", err)
	}

	return nil
}

func hasSQLiteMigrationVersion(db *sql.DB, version int) (bool, error) {
	var exists bool

	err := db.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM schema_migrations WHERE version = ?)",
		version,
	).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("check sqlite migration version: %w", err)
	}

	return exists, nil
}

func applySQLiteMigration(db *sql.DB, migration SQLiteMigration, appliedAt time.Time) error {
	if migration.Version <= 0 {
		return fmt.Errorf("sqlite migration version must be positive")
	}

	if migration.Apply == nil {
		return fmt.Errorf("sqlite migration apply function is required")
	}

	if appliedAt.IsZero() {
		return fmt.Errorf("sqlite migration applied date is required")
	}

	applied, err := hasSQLiteMigrationVersion(db, migration.Version)
	if err != nil {
		return err
	}

	if applied {
		return nil
	}

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("begin sqlite migration transaction: %w", err)
	}
	defer tx.Rollback()

	if err := migration.Apply(tx); err != nil {
		return fmt.Errorf("apply sqlite migration %d: %w", migration.Version, err)
	}

	if err := recordSQLiteMigrationVersion(tx, migration.Version, appliedAt); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit sqlite migration transaction: %w", err)
	}

	return nil
}

func applySQLiteMigrations(db *sql.DB, migrations []SQLiteMigration, appliedAt time.Time) error {
	orderedMigrations := append([]SQLiteMigration(nil), migrations...)

	sort.Slice(orderedMigrations, func(i, j int) bool {
		return orderedMigrations[i].Version < orderedMigrations[j].Version
	})

	for index, migration := range orderedMigrations {
		if index > 0 && migration.Version == orderedMigrations[index-1].Version {
			return fmt.Errorf("duplicate sqlite migration version: %d", migration.Version)
		}

		if err := applySQLiteMigration(db, migration, appliedAt); err != nil {
			return err
		}
	}

	return nil
}

func sqliteMigrations() []SQLiteMigration {
	return []SQLiteMigration{}
}

func hasCurrentSQLiteSchema(db *sql.DB) (bool, error) {
	expectedTables := []string{
		"schema_migrations",
		"observations",
		"medical_profiles",
		"measurements",
		"appointments",
	}

	for _, tableName := range expectedTables {
		var exists bool
		err := db.QueryRow(
			"SELECT EXISTS(SELECT 1 FROM sqlite_master WHERE type = 'table' AND name = ?)",
			tableName,
		).Scan(&exists)
		if err != nil {
			return false, fmt.Errorf("check sqlite table %s: %w", tableName, err)
		}
		if !exists {
			return false, nil
		}
	}

	return true, nil
}

func markCurrentSQLiteSchemaBaseline(db *sql.DB, appliedAt time.Time) error {
	if appliedAt.IsZero() {
		return fmt.Errorf("sqlite schema baseline applied date is required")
	}

	hasSchema, err := hasCurrentSQLiteSchema(db)
	if err != nil {
		return err
	}
	if !hasSchema {
		return fmt.Errorf("sqlite schema does not match current baseline")
	}

	versions, err := appliedSQLiteMigrationVersions(db)
	if err != nil {
		return err
	}
	if len(versions) != 0 {
		return fmt.Errorf("sqlite schema baseline already has applied migrations")
	}

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("begin sqlite schema baseline transaction: %w", err)
	}
	defer tx.Rollback()

	if err := recordSQLiteMigrationVersion(tx, 1, appliedAt); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit sqlite schema baseline transaction: %w", err)
	}

	return nil
}
