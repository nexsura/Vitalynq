package main

import (
	"testing"
	"time"
)

func TestValidateMeasurementAcceptsValidMeasurement(t *testing.T) {
	measurement := validMeasurement()

	if err := validateMeasurement(measurement); err != nil {
		t.Fatalf("validateMeasurement() error = %v, want nil", err)
	}
}

func TestValidateMeasurementRejectsMissingDate(t *testing.T) {
	measurement := validMeasurement()
	measurement.OccurredAt = time.Time{}

	if err := validateMeasurement(measurement); err == nil {
		t.Fatalf("validateMeasurement() error = nil, want error")
	}
}

func TestValidateMeasurementRejectsMissingUnit(t *testing.T) {
	measurement := validMeasurement()
	measurement.Unit = "	"

	if err := validateMeasurement(measurement); err == nil {
		t.Fatalf("validateMeasurement() error = nil, want error")
	}
}

func TestValidateMeasurementRejectsMissingIndicator(t *testing.T) {
	measurement := validMeasurement()
	measurement.Indicator = "   "

	if err := validateMeasurement(measurement); err == nil {
		t.Fatalf("validateMeasurement() error = nil, want error")
	}
}

func TestValidateMeasurementRejectsMissingContext(t *testing.T) {
	measurement := validMeasurement()
	measurement.Context = "   "

	if err := validateMeasurement(measurement); err == nil {
		t.Fatalf("validateMeasurement() error = nil, want error")
	}
}

func TestValidateMeasurementRejectsMissingMethod(t *testing.T) {
	measurement := validMeasurement()
	measurement.Method = "   "

	if err := validateMeasurement(measurement); err == nil {
		t.Fatalf("validateMeasurement() error = nil, want error")
	}
}

func TestValidateMeasurementRejectsMissingSource(t *testing.T) {
	measurement := validMeasurement()
	measurement.Source = "   "

	if err := validateMeasurement(measurement); err == nil {
		t.Fatalf("validateMeasurement() error = nil, want error")
	}
}

func validMeasurement() Measurement {
	return Measurement{
		OccurredAt: testTime(),
		CreatedAt:  testTime(),
		Indicator:  "poids",
		Value:      72.5,
		Unit:       "kg",
		Context:    "test fictif",
		Method:     "saisie manuelle",
		Source:     "saisie manuelle",
	}
}
