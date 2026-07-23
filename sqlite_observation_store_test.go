package main

import (
	"testing"
)

func TestNewSQLiteObservationStore(t *testing.T) {
	db, err := openSQLite(":memory:")
	if err != nil {
		t.Fatalf("openSQLite() error = %v, want nil", err)
	}
	defer db.Close()

	store := NewSQLiteObservationStore(db)
	if store == nil {
		t.Fatalf("NewSQLiteObservationStore() = nil, want store")
	}

	if store.db != db {
		t.Fatalf("store db does not match input db")
	}
}

func TestSQLiteObservationStoreImplementsObservationStore(t *testing.T) {
	var _ ObservationStore = (*SQLiteObservationStore)(nil)
}
