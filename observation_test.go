package main

import (
	"testing"
	"time"
)

func TestValidateObservationAcceptsValidObservation(t *testing.T) {
	observation := Observation{
		OccurredAt: time.Date(2026, 7, 17, 9, 30, 0, 0, time.UTC),
		Text:       "Observation fictive de test",
		Source:     "saisie manuelle",
	}

	if err := validateObservation(observation); err != nil {
		t.Fatalf("validateObservation() error = %v, want nil", err)
	}
}

func TestValidateObservationRejectsMissingDate(t *testing.T) {
	observation := Observation{
		Text:   "Observation fictive de test",
		Source: "saisie manuelle",
	}

	if err := validateObservation(observation); err == nil {
		t.Fatalf("validateObservation() error = nil, want error")
	}
}

func TestValidateObservationRejectsEmptyText(t *testing.T) {
	observation := Observation{
		OccurredAt: time.Date(2026, 7, 17, 9, 30, 0, 0, time.UTC),
		Text:       "",
		Source:     "saisie manuelle",
	}

	if err := validateObservation(observation); err == nil {
		t.Fatalf("validateObservation() error = nil, want error")
	}
}

func TestValidateObservationRejectsBlankText(t *testing.T) {
	observation := Observation{
		OccurredAt: time.Date(2026, 7, 17, 9, 30, 0, 0, time.UTC),
		Text:       " ",
		Source:     "saisie manuelle",
	}

	if err := validateObservation(observation); err == nil {
		t.Fatalf("validateObservation() error = nil, want error")
	}
}

func TestValidateObservationRejectsEmptySource(t *testing.T) {
	observation := Observation{
		OccurredAt: time.Date(2026, 7, 17, 9, 30, 0, 0, time.UTC),
		Text:       "Observation fictive de test",
		Source:     "",
	}

	if err := validateObservation(observation); err == nil {
		t.Fatalf("validateObservation() error = nil, want error")
	}
}

func TestValidateObservationRejectsBlankSource(t *testing.T) {
	observation := Observation{
		OccurredAt: time.Date(2026, 7, 17, 9, 30, 0, 0, time.UTC),
		Text:       "Observation fictive de test",
		Source:     "   ",
	}

	if err := validateObservation(observation); err == nil {
		t.Fatalf("validateObservation() error = nil, want error")
	}
}

func TestNewObservationCreateValidObservation(t *testing.T) {
	occurredAt := time.Date(2026, 7, 17, 9, 30, 0, 0, time.UTC)

	observation, err := newObservation(occurredAt, " Observation fictive de test ", " saisie manuelle ")
	if err != nil {
		t.Fatalf("newObservation() error = %v, want nil", err)
	}

	if observation.OccurredAt != occurredAt {
		t.Fatalf("OccurredAt = %v , want %v", observation.OccurredAt, occurredAt)
	}

	if observation.Text != "Observation fictive de test" {
		t.Fatalf("Text = %q, want %q", observation.Text, "Observation fictive de test")
	}

	if observation.Source != "saisie manuelle" {
		t.Fatalf("Source = %q, want %q", observation.Source, "saisie manuelle")
	}
}

func TestObservationRejectsInvalidObservation(t *testing.T) {
	_, err := newObservation(time.Time{}, "Observation fictive de test", "saisie manuelle")
	if err == nil {
		t.Fatalf("newObservation() error = nil, want error")
	}
}
