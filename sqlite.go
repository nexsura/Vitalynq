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
