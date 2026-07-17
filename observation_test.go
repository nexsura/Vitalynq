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

func TestValidateOservationRejectsMissingDate(t *testing.T) {
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
		t.Fatalf("validationObservation() error = nil, want error")
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
