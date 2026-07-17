package main

import (
	"errors"
	"strings"
	"time"
)

type Observation struct {
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
