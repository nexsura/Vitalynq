package main

import (
	"fmt"
)

type ObservationStore interface {
	Save(observation Observation) (Observation, error)
	List() []Observation
}

type MemoryObservationStore struct {
	nextID       int64
	observations []Observation
}

func NewMemoryObservationStore() *MemoryObservationStore {
	return &MemoryObservationStore{
		nextID: 1,
	}
}

func (store *MemoryObservationStore) Save(observation Observation) (Observation, error) {
	if err := validateObservation(observation); err != nil {
		return Observation{}, fmt.Errorf("save observation: %w", err)
	}

	observation.ID = store.nextID
	store.nextID++

	store.observations = append(store.observations, observation)
	return observation, nil
}

func (store *MemoryObservationStore) List() []Observation {
	observations := make([]Observation, len(store.observations))
	copy(observations, store.observations)

	return observations
}
