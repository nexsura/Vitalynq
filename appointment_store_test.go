package main

import (
	"testing"
)

func TestMemoryAppointmentStoreSavesAppointmentWithID(t *testing.T) {
	store := NewMemoryAppointmentStore()

	saved, err := store.Save(validAppointment())
	if err != nil {
		t.Fatalf("Save() error = %v, want nil", err)
	}

	if saved.ID != 1 {
		t.Fatalf("ID = %d, want 1", saved.ID)
	}
}

func TestMemoryAppointmentStoreRejectsInvalidAppointment(t *testing.T) {
	store := NewMemoryAppointmentStore()

	_, err := store.Save(Appointment{})
	if err == nil {
		t.Fatalf("Save() error = nil, want error")
	}
}

func TestMemoryAppointmentStoreListsSavedAppointments(t *testing.T) {
	store := NewMemoryAppointmentStore()

	first, err := store.Save(validAppointment())
	if err != nil {
		t.Fatalf("Save(first) error = %v, want nil", err)
	}

	secondAppointment := validAppointment()
	secondAppointment.Title = "Examen fictif"

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
}
