package main

import "fmt"

type ExportSnapshot struct {
	Profile      *MedicalProfile
	Observations []Observation
	Measurements []Measurement
	Appointments []Appointment
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
		Observations: observations,
		Measurements: measurements,
		Appointments: appointments,
	}

	if found {
		snapshot.Profile = &profile
	}

	return snapshot, nil
}
