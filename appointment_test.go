package main

import (
	"testing"
)

func TestValidateAppointmentAcceptsValidAppointment(t *testing.T) {
	appointment := validAppointment()

	if err := validateAppointment(appointment); err != nil {
		t.Fatalf("validateAppointment() error = %v, want nil", err)
	}
}

func TestValidateAppointmentRejectsMissingScheduleAt(t *testing.T) {
	appointment := validAppointment()
	appointment.ScheduledAt = testTimeFromDate(1, 1, 1)

	if err := validateAppointment(appointment); err == nil {
		t.Fatalf("validateAppointment() error = nil, want error")
	}
}

func TestValidateAppointmentRejectsMissingTitle(t *testing.T) {
	appointment := validAppointment()
	appointment.Title = "	"

	if err := validateAppointment(appointment); err == nil {
		t.Fatalf("validateAppointment() error = nil, want error")
	}
}

func TestNewAppointmentCreatesValidAppointment(t *testing.T) {
	appointment, err := newAppointment(testTime(), "consultation fictive", "rendez-vous", "cabinet fictif", "saisie manuelle")
	if err != nil {
		t.Fatalf("newAppointment() error = %v, want nil", err)
	}

	if appointment.Title != "consultation fictive" {
		t.Fatalf("Title = %q, want %q", appointment.Title, "consultation fictive")
	}

	if appointment.Category != "rendez-vous" {
		t.Fatalf("Category = %q, want %q", appointment.Category, "rendez-vous")
	}

	if appointment.CreatedAt.IsZero() {
		t.Fatalf("CreatedAt is zero, want creation date")
	}
}

func validAppointment() Appointment {
	return Appointment{
		ScheduledAt: testTime(),
		CreatedAt:   testTime(),
		Title:       "consultation fictive",
		Category:    "rendez-vous",
		Location:    "cabinet fictif",
		Source:      "saisie manuelle",
	}
}
