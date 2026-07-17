package main

import (
	"errors"
	"strings"
	"time"
)

type Observation struct {
	ID         int64
	OccurredAt time.Time
	Text       string
	Source     string
}

func validateObservation(observation Observation) error {
	if observation.OccurredAt.IsZero() {
		return errors.New("observation date is required")
	}

	if strings.TrimSpace(observation.Text) == "" {
		return errors.New("observation text is required")
	}

	if strings.TrimSpace(observation.Source) == "" {
		return errors.New("observation source is required")
	}

	return nil
}

func newObservation(occurredAt time.Time, text string, source string) (Observation, error) {
	observation := Observation{
		OccurredAt: occurredAt,
		Text:       strings.TrimSpace(text),
		Source:     strings.TrimSpace(source),
	}

	if err := validateObservation(observation); err != nil {
		return Observation{}, err
	}

	return observation, nil
}
