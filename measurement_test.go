package main

import (
	"testing"
)

func TestValidateMeasurementAcceptsValidMeasurement(t *testing.T) {
	measurement := Measurement{
		OccurredAt: testTime(),
		CreatedAt:  testTime(),
		Indicator:  "poids",
		Value:      72.5,
		Unit:       "kg",
		Context:    "test fictif",
		Method:     "saisie manuelle",
		Source:     "saisie manuelle",
	}

	if err := validateMeasurement(measurement); err != nil {
		t.Fatalf("validateMeasurement() error = %v, want nil", err)
	}
}

func TestValidateMeasurementRejectsMissingDate(t *testing.T) {
	measurement := Measurement{
		CreatedAt: testTime(),
		Indicator: "poids",
		Value:     72.5,
		Unit:      "kg",
		Context:   "test fictif",
		Method:    "saisie manuelle",
		Source:    "saisie manuelle",
	}

	if err := validateMeasurement(measurement); err == nil {
		t.Fatalf("validateMeasurement() error = nil, want error")
	}
}

func TestValidateMeasurementRejectsMissingUnit(t *testing.T) {
	measurement := Measurement{
		OccurredAt: testTime(),
		CreatedAt:  testTime(),
		Indicator:  "poids",
		Value:      72.5,
		Unit:       "	",
		Context:    "test fictif",
		Method:     "saisie manuelle",
		Source:     "saisie manuelle",
	}

	if err := validateMeasurement(measurement); err == nil {
		t.Fatalf("validateMeasurement() error = nil, want error")
	}
}
