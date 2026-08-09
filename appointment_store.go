package main

import "fmt"

type AppointmentStore interface {
	Save(appointment Appointment) (Appointment, error)
	List() ([]Appointment, error)
}

type MemoryAppointmentStore struct {
	nextID       int64
	appointments []Appointment
}

func NewMemoryAppointmentStore() *MemoryAppointmentStore {
	return &MemoryAppointmentStore{
		nextID: 1,
	}
}

func (store *MemoryAppointmentStore) Save(appointment Appointment) (Appointment, error) {
	if err := validateAppointment(appointment); err != nil {
		return Appointment{}, fmt.Errorf("save appointment: %w", err)
	}

	appointment.ID = store.nextID
	store.nextID++

	store.appointments = append(store.appointments, appointment)

	return appointment, nil
}

func (store *MemoryAppointmentStore) List() ([]Appointment, error) {
	appointments := make([]Appointment, len(store.appointments))
	copy(appointments, store.appointments)

	return appointments, nil
}
