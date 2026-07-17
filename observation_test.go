package main

import (
	"testing"
	"time"
)

func TestValidateObservationAcceptsValidObservation(t *testing.T) {
	observation := Observation{
		OccurredAt: testTime(),
		CreatedAt:  testTime(),
		Text:       "Observation fictive de test",
		Source:     "saisie manuelle",
	}

	if err := validateObservation(observation); err != nil {
		t.Fatalf("validateObservation() error = %v, want nil", err)
	}
}

func TestValidateObservationRejectsMissingDate(t *testing.T) {
	observation := Observation{
		CreatedAt: testTime(),
		Text:      "Observation fictive de test",
		Source:    "saisie manuelle",
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
		Source:     "saisie manuelle",
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
		Source:     "saisie manuelle",
	}

	if err := validateObservation(observation); err == nil {
		t.Fatalf("validateObservation() error = nil, want error")
	}
}

func TestValidateObservationRejectsEmptySource(t *testing.T) {
	observation := Observation{
		OccurredAt: testTime(),
		CreatedAt:  testTime(),
		Text:       "Observation fictive de test",
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
		Text:       "Observation fictive de test",
		Source:     "   ",
	}

	if err := validateObservation(observation); err == nil {
		t.Fatalf("validateObservation() error = nil, want error")
	}
}

func TestNewObservationCreatesValidObservation(t *testing.T) {
	occurredAt := testTime()

	observation, err := newObservation(occurredAt, " Observation fictive de test ", " saisie manuelle ")
	if err != nil {
		t.Fatalf("newObservation() error = %v, want nil", err)
	}

	if observation.OccurredAt != occurredAt {
		t.Fatalf("OccurredAt = %v, want %v", observation.OccurredAt, occurredAt)
	}

	if observation.Text != "Observation fictive de test" {
		t.Fatalf("Text = %q, want %q", observation.Text, "Observation fictive de test")
	}

	if observation.Source != "saisie manuelle" {
		t.Fatalf("Source = %q, want %q", observation.Source, "saisie manuelle")
	}

	if observation.ID != 0 {
		t.Fatalf("ID = %d, want 0", observation.ID)
	}

	if observation.CreatedAt.IsZero() {
		t.Fatalf("CreatedAt is zero, want creation date")
	}
}

func TestNewObservationRejectsInvalidObservation(t *testing.T) {
	_, err := newObservation(time.Time{}, "Observation fictive de test", "saisie manuelle")
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
		Text:       "Observation fictive de test",
		Source:     "saisie manuelle",
	}

	if err := validateObservation(observation); err == nil {
		t.Fatalf("validateObservation() error = nil, want error")
	}
}
