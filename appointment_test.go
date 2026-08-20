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

func TestValidateAppointmentRejectsMissingScheduledAt(t *testing.T) {
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

func TestValidateAppointmentRejectsMissingCategory(t *testing.T) {
	appointment := validAppointment()
	appointment.Category = "	"

	if err := validateAppointment(appointment); err == nil {
		t.Fatalf("validateAppointment() error = nil, want error")
	}
}

func TestValidateAppointmentRejectsMissingSource(t *testing.T) {
	appointment := validAppointment()
	appointment.Source = "	"

	if err := validateAppointment(appointment); err == nil {
		t.Fatalf("validateAppointment() error = nil, want error")
	}
}

func TestNewAppointmentCreatesValidAppointment(t *testing.T) {
	appointment, err := newAppointment(testTime(), "fictive consultation", "appointment", "fictive office", "manual entry")
	if err != nil {
		t.Fatalf("newAppointment() error = %v, want nil", err)
	}

	if appointment.Title != "fictive consultation" {
		t.Fatalf("Title = %q, want %q", appointment.Title, "fictive consultation")
	}

	if appointment.Category != "appointment" {
		t.Fatalf("Category = %q, want %q", appointment.Category, "appointment")
	}

	if appointment.CreatedAt.IsZero() {
		t.Fatalf("CreatedAt is zero, want creation date")
	}
}

func validAppointment() Appointment {
	return Appointment{
		ScheduledAt: testTime(),
		CreatedAt:   testTime(),
		Title:       "fictive consultation",
		Category:    "appointment",
		Location:    "fictive office",
		Source:      "manual entry",
	}
}
