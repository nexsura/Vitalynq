package main

import (
	"testing"
)

func TestMemoryMeasurementStoreSavesMeasurementWithID(t *testing.T) {
	store := newMemoryMeasurementStore()

	saved, err := store.Save(validMeasurement())
	if err != nil {
		t.Fatalf("Save() error = %v, want nil", err)
	}

	if saved.ID != 1 {
		t.Fatalf("ID = %d, want 1", saved.ID)
	}
}

func TestMemoryMeasurementStoreRejectsInvalidMeasurement(t *testing.T) {
	store := newMemoryMeasurementStore()

	_, err := store.Save(Measurement{})
	if err == nil {
		t.Fatalf("Save() error = nil, want error")
	}
}

func TestMemoryMeasurementStoreListsSaveMeasurements(t *testing.T) {
	store := newMemoryMeasurementStore()

	first, err := store.Save(validMeasurement())
	if err != nil {
		t.Fatalf("Save(first) error = %v, want nil", err)
	}

	secondMeasurement := validMeasurement()
	secondMeasurement.Indicator = "taille"
	secondMeasurement.Value = 175
	secondMeasurement.Unit = "cm"

	second, err := store.Save(secondMeasurement)
	if err != nil {
		t.Fatalf("Save(second) error = %v, want nil", err)
	}

	measurement, err := store.List()
	if err != nil {
		t.Fatalf("List() error = %v, want nil", err)
	}

	if len(measurement) != 2 {
		t.Fatalf("len(List()) = %d, want 2", len(measurement))
	}

	if measurement[0].ID != first.ID {
		t.Fatalf("first ID = %d, want %d", measurement[0].ID, first.ID)
	}

	if measurement[1].ID != second.ID {
		t.Fatalf("second ID = %d, want %d", measurement[1].ID, second.ID)
	}
}
