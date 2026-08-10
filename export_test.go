package main

import (
	"testing"
)

func TestBuildExportSnapshot(t *testing.T) {
	profileStore := NewMemoryMedicalProfileStore()
	observationStore := NewMemoryObservationStore()
	measurementStore := NewMemoryMeasurementStore()
	appointmentStore := NewMemoryAppointmentStore()

	if _, err := profileStore.Save(validMedicalProfile("Profil fictif de test")); err != nil {
		t.Fatalf("Save(profile) error = %v, want nil", err)
	}

	if _, err := observationStore.Save(validStoreObservation("Observation fictive")); err != nil {
		t.Fatalf("Save(observation) error = %v, want nil", err)
	}

	if _, err := measurementStore.Save(validMeasurement()); err != nil {
		t.Fatalf("Save(measurement) error = %v, want nil", err)
	}

	if _, err := appointmentStore.Save(validAppointment()); err != nil {
		t.Fatalf("Save(appointment) error = %v, want nil", err)
	}

	snapshot, err := buildExportSnapshot(profileStore, observationStore, measurementStore, appointmentStore)
	if err != nil {
		t.Fatalf("buildExportSnapshot() error = %v, want nil", err)
	}

	if snapshot.Profile == nil {
		t.Fatalf("Profile = nil, want profile")
	}

	if len(snapshot.Observations) != 1 {
		t.Fatalf("len(Observations) = %d, want 1", len(snapshot.Observations))
	}

	if len(snapshot.Measurements) != 1 {
		t.Fatalf("len(Measurements) = %d, want 1", len(snapshot.Measurements))
	}

	if len(snapshot.Appointments) != 1 {
		t.Fatalf("len(Appointments) = %d, want 1", len(snapshot.Appointments))
	}
}
