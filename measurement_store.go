package main

import "fmt"

type MeasurementStore interface {
	Save(measurement Measurement) (Measurement, error)
	List() ([]Measurement, error)
}

type MemoryMeasurementStore struct {
	nextID       int64
	measurements []Measurement
}

func newMemoryMeasurementStore() *MemoryMeasurementStore {
	return &MemoryMeasurementStore{
		nextID: 1,
	}
}

func (store *MemoryMeasurementStore) Save(measurement Measurement) (Measurement, error) {
	if err := validateMeasurement(measurement); err != nil {
		return Measurement{}, fmt.Errorf("save measurement: %w", err)
	}

	measurement.ID = store.nextID
	store.nextID++

	store.measurements = append(store.measurements, measurement)

	return measurement, nil
}

func (store *MemoryMeasurementStore) List() ([]Measurement, error) {
	measurements := make([]Measurement, len(store.measurements))
	copy(measurements, store.measurements)

	return measurements, nil
}
