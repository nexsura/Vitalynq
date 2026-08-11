package main

import (
	"encoding/json"
	"fmt"
)

const exportVersion = "1"

type ExportSnapshot struct {
	ExportVersion string          `json:"export_version"`
	Profile       *MedicalProfile `json:"profile"`
	Observations  []Observation   `json:"observations"`
	Measurements  []Measurement   `json:"measurements"`
	Appointments  []Appointment   `json:"appointments"`
}

func buildExportSnapshot(profileStore MedicalProfileStore, observationStore ObservationStore, measurementStore MeasurementStore, appointmentStore AppointmentStore) (ExportSnapshot, error) {
	profile, found, err := profileStore.Get()
	if err != nil {
		return ExportSnapshot{}, fmt.Errorf("export profile: %w", err)
	}

	observations, err := observationStore.List()
	if err != nil {
		return ExportSnapshot{}, fmt.Errorf("export observations: %w", err)
	}

	measurements, err := measurementStore.List()
	if err != nil {
		return ExportSnapshot{}, fmt.Errorf("export measurements: %w", err)
	}

	appointments, err := appointmentStore.List()
	if err != nil {
		return ExportSnapshot{}, fmt.Errorf("export appointments: %w", err)
	}

	snapshot := ExportSnapshot{
		ExportVersion: exportVersion,
		Observations:  observations,
		Measurements:  measurements,
		Appointments:  appointments,
	}

	if found {
		snapshot.Profile = &profile
	}

	return snapshot, nil
}

func exportSnapshotJSON(snapshot ExportSnapshot) (string, error) {
	data, err := json.MarshalIndent(snapshot, "", " ")
	if err != nil {
		return "", fmt.Errorf("export snapshot json: %w", err)
	}

	return string(data), nil
}
