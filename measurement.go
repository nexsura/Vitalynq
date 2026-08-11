package main

import (
	"errors"
	"strings"
	"time"
)

type Measurement struct {
	ID         int64     `json:"id"`
	OccurredAt time.Time `json:"occurred_at"`
	CreatedAt  time.Time `json:"created_at"`
	Indicator  string    `json:"indicator"`
	Value      float64   `json:"value"`
	Unit       string    `json:"unit"`
	Context    string    `json:"context"`
	Method     string    `json:"method"`
	Source     string    `json:"source"`
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

	if strings.TrimSpace(measurement.Context) == "" {
		return errors.New("measurement context is required")
	}

	if strings.TrimSpace(measurement.Method) == "" {
		return errors.New("measurement method is required")
	}

	if strings.TrimSpace(measurement.Source) == "" {
		return errors.New("measurement source is required")
	}

	return nil
}

func newMeasurement(occurredAt time.Time, indicator string, value float64, unit string, context string, method string, source string) (Measurement, error) {
	measurement := Measurement{
		OccurredAt: occurredAt,
		CreatedAt:  time.Now().UTC(),
		Indicator:  strings.TrimSpace(indicator),
		Value:      value,
		Unit:       strings.TrimSpace(unit),
		Context:    strings.TrimSpace(context),
		Method:     strings.TrimSpace(method),
		Source:     strings.TrimSpace(source),
	}

	if err := validateMeasurement(measurement); err != nil {
		return Measurement{}, err
	}

	return measurement, nil
}
