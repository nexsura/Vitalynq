package main

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

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
