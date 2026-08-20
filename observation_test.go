package main

import (
	"testing"
	"time"
)

func TestValidateObservationAcceptsValidObservation(t *testing.T) {
	observation := Observation{
		OccurredAt: testTime(),
		CreatedAt:  testTime(),
		Text:       "Fictive test observation",
		Source:     "manual entry",
	}

	if err := validateObservation(observation); err != nil {
		t.Fatalf("validateObservation() error = %v, want nil", err)
	}
}

func TestValidateObservationRejectsMissingDate(t *testing.T) {
	observation := Observation{
		CreatedAt: testTime(),
		Text:      "Fictive test observation",
		Source:    "manual entry",
	}

	if err := validateObservation(observation); err == nil {
		t.Fatalf("validateObservation() error = nil, want error")
	}
}

func TestValidateObservationRejectsEmptyText(t *testing.T) {
	observation := Observation{
		OccurredAt: testTime(),
		CreatedAt:  testTime(),
		Text:       "",
		Source:     "manual entry",
	}

	if err := validateObservation(observation); err == nil {
		t.Fatalf("validateObservation() error = nil, want error")
	}
}

func TestValidateObservationRejectsBlankText(t *testing.T) {
	observation := Observation{
		OccurredAt: testTime(),
		CreatedAt:  testTime(),
		Text:       " ",
		Source:     "manual entry",
	}

	if err := validateObservation(observation); err == nil {
		t.Fatalf("validateObservation() error = nil, want error")
	}
}

func TestValidateObservationRejectsEmptySource(t *testing.T) {
	observation := Observation{
		OccurredAt: testTime(),
		CreatedAt:  testTime(),
		Text:       "Fictive test observation",
		Source:     "",
	}

	if err := validateObservation(observation); err == nil {
		t.Fatalf("validateObservation() error = nil, want error")
	}
}

func TestValidateObservationRejectsBlankSource(t *testing.T) {
	observation := Observation{
		OccurredAt: testTime(),
		CreatedAt:  testTime(),
		Text:       "Fictive test observation",
		Source:     "   ",
	}

	if err := validateObservation(observation); err == nil {
		t.Fatalf("validateObservation() error = nil, want error")
	}
}

func TestNewObservationCreatesValidObservation(t *testing.T) {
	occurredAt := testTime()

	observation, err := newObservation(occurredAt, " Fictive test observation ", " manual entry ")
	if err != nil {
		t.Fatalf("newObservation() error = %v, want nil", err)
	}

	if observation.OccurredAt != occurredAt {
		t.Fatalf("OccurredAt = %v, want %v", observation.OccurredAt, occurredAt)
	}

	if observation.Text != "Fictive test observation" {
		t.Fatalf("Text = %q, want %q", observation.Text, "Fictive test observation")
	}

	if observation.Source != "manual entry" {
		t.Fatalf("Source = %q, want %q", observation.Source, "manual entry")
	}

	if observation.ID != 0 {
		t.Fatalf("ID = %d, want 0", observation.ID)
	}

	if observation.CreatedAt.IsZero() {
		t.Fatalf("CreatedAt is zero, want creation date")
	}
}

func TestNewObservationRejectsInvalidObservation(t *testing.T) {
	_, err := newObservation(time.Time{}, "Fictive test observation", "manual entry")
	if err == nil {
		t.Fatalf("newObservation() error = nil, want error")
	}
}

func testTime() time.Time {
	return time.Date(2026, 7, 17, 9, 30, 0, 0, time.UTC)
}

func TestValidateObservationRejectsMissingCreationDate(t *testing.T) {
	observation := Observation{
		OccurredAt: testTime(),
		Text:       "Fictive test observation",
		Source:     "manual entry",
	}

	if err := validateObservation(observation); err == nil {
		t.Fatalf("validateObservation() error = nil, want error")
	}
}
