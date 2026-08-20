package main

import (
	"database/sql"
	"testing"
)

func TestNewSQLiteAppointmentStore(t *testing.T) {
	db, err := openSQLite(":memory:")
	if err != nil {
		t.Fatalf("openSQLite() error = %v, want nil", err)
	}
	defer db.Close()

	store := NewSQLiteAppointmentStore(db)
	if store == nil {
		t.Fatalf("NewSQLiteAppointmentStore() = nil, want store")
	}

	if store.db != db {
		t.Fatalf("store db does not match input db")
	}
}

func TestSQLiteAppointmentStoreImplementsAppointmentStore(t *testing.T) {
	var _ AppointmentStore = (*SQLiteAppointmentStore)(nil)
}

func newTestSQLiteAppointmentStore(t *testing.T) (*sql.DB, *SQLiteAppointmentStore) {
	t.Helper()

	db, err := openSQLite(":memory:")
	if err != nil {
		t.Fatalf("openSQLite() error = %v, want nil", err)
	}

	if err := initializeSQLiteSchema(db); err != nil {
		t.Fatalf("initializeSQLiteSchema() error = %v, want nil", err)
	}

	return db, NewSQLiteAppointmentStore(db)
}

func TestSQLiteAppointmentStoreSavesAppointment(t *testing.T) {
	db, store := newTestSQLiteAppointmentStore(t)
	defer db.Close()

	saved, err := store.Save(validAppointment())
	if err != nil {
		t.Fatalf("Save() error = %v, want nil", err)
	}

	if saved.ID != 1 {
		t.Fatalf("ID = %d, want 1", saved.ID)
	}
}

func TestSQLiteAppointmentStoreRejectsInvalidAppointment(t *testing.T) {
	db, store := newTestSQLiteAppointmentStore(t)
	defer db.Close()

	_, err := store.Save(Appointment{})
	if err == nil {
		t.Fatalf("Save() error = nil, want error")
	}
}

func TestSQLiteAppointmentStoreListsSavedAppointments(t *testing.T) {
	db, store := newTestSQLiteAppointmentStore(t)
	defer db.Close()

	first, err := store.Save(validAppointment())
	if err != nil {
		t.Fatalf("Save(first) error = %v, want nil", err)
	}

	secondAppointment := validAppointment()
	secondAppointment.Title = "Fictive exam"

	second, err := store.Save(secondAppointment)
	if err != nil {
		t.Fatalf("Save(second) error = %v, want nil", err)
	}

	appointments, err := store.List()
	if err != nil {
		t.Fatalf("List() error = %v, want nil", err)
	}

	if len(appointments) != 2 {
		t.Fatalf("len(List()) = %d, want 2", len(appointments))
	}

	if appointments[0].ID != first.ID {
		t.Fatalf("first ID = %d, want %d", appointments[0].ID, first.ID)
	}

	if appointments[1].ID != second.ID {
		t.Fatalf("second ID = %d, want %d", appointments[1].ID, second.ID)
	}

	if appointments[0].Title != "fictive consultation" {
		t.Fatalf("first Title = %q, want %q", appointments[0].Title, "fictive consultation")
	}

	if appointments[1].Title != "Fictive exam" {
		t.Fatalf("second Title = %q, want %q", appointments[1].Title, "Fictive exam")
	}
}
