package main

import (
	"errors"
	"strings"
	"time"
)

type Measurement struct {
	ID         int64
	OccurredAt time.Time
	CreatedAt  time.Time
	Indicator  string
	Value      float64
	Unit       string
	Context    string
	Method     string
	Source     string
}

func validateMeasurement(measurement Measurement) error {
	if measurement.OccurredAt.IsZero() {
		return errors.New("measurement date is required")
	}

	if measurement.CreatedAt.IsZero() {
		return errors.New("measurement creation date is required")
	}

	if strings.TrimSpace(measurement.Indicator) == "" {
		return errors.New("measurement indicator is required")
	}

	if strings.TrimSpace(measurement.Unit) == "" {
		return errors.New("measurement unit is required")
	}

	return nil
}
