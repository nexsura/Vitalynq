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
