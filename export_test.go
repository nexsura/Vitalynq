package main

import (
	"encoding/json"
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

	if snapshot.ExportVersion != exportVersion {
		t.Fatalf("ExportVersion = %q, want %q", snapshot.ExportVersion, exportVersion)
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

func TestExportSnapshotJSON(t *testing.T) {
	snapshot := ExportSnapshot{
		ExportVersion: exportVersion,
		Observations:  []Observation{validStoreObservation("Observation fictive")},
		Measurements:  []Measurement{validMeasurement()},
		Appointments:  []Appointment{validAppointment()},
	}

	jsonText, err := exportSnapshotJSON(snapshot)
	if err != nil {
		t.Fatalf("exportSnapshotJSON() error = %v, want nil", err)
	}

	var exported map[string]any
	if err := json.Unmarshal([]byte(jsonText), &exported); err != nil {
		t.Fatalf("json.Unmarshal() error = %v, want nil", err)
	}

	assertJSONKey(t, exported, "profile")
	assertJSONKey(t, exported, "observations")
	assertJSONKey(t, exported, "measurements")
	assertJSONKey(t, exported, "appointments")
	assertJSONKey(t, exported, "export_version")

	assertContains(t, jsonText, `"id"`)
	assertContains(t, jsonText, `"created_at"`)
	assertContains(t, jsonText, `"occurred_at"`)
	assertContains(t, jsonText, `"scheduled_at"`)
	assertContains(t, jsonText, `"indicator"`)
	assertContains(t, jsonText, `"value"`)
	assertContains(t, jsonText, `"unit"`)
	assertContains(t, jsonText, `"export_version"`)
	assertContains(t, jsonText, `"1"`)
}

func assertJSONKey(t *testing.T, exported map[string]any, key string) {
	t.Helper()

	if _, exists := exported[key]; !exists {
		t.Fatalf("JSON key %q is missing", key)
	}
}
