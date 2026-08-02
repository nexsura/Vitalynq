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

func TestNewMeasurementCreatesValidMeasurement(t *testing.T) {
	measurement, err := newMeasurement(
		testTime(),
		" poids ",
		72.5,
		" kg ",
		" test fictif ",
		" saisie manuelle ",
		" saisie manuelle ",
	)

	if err != nil {
		t.Fatalf("newMeasurement() error = %v, want nil", err)
	}

	if measurement.Indicator != "poids" {
		t.Fatalf("Indicator = %q, want %q", measurement.Indicator, "poids")
	}

	if measurement.Unit != "kg" {
		t.Fatalf("Unit = %q, want %q", measurement.Unit, "kg")
	}

	if measurement.CreatedAt.IsZero() {
		t.Fatalf("CreatedAt is zero, want creation date")
	}

	if measurement.ID != 0 {
		t.Fatalf("ID = %d, want 0", measurement.ID)
	}
}

func TestNewMeasurementRejectsInvalidMeasurement(t *testing.T) {
	_, err := newMeasurement(testTime(), "", 72.5, "kg", "test fictif", "saisie manuelle", "saisie manuelle")
	if err == nil {
		t.Fatalf("newMeasurement() error = nil, want error")
	}
}
