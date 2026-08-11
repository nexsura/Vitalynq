package main

import (
	"errors"
	"strings"
	"time"
)

type Appointment struct {
	ID          int64     `json:"id"`
	ScheduledAt time.Time `json:"scheduled_at"`
	CreatedAt   time.Time `json:"created_at"`
	Title       string    `json:"title"`
	Category    string    `json:"category"`
	Location    string    `json:"location"`
	Source      string    `json:"source"`
}

func validateAppointment(appointment Appointment) error {
	if appointment.ScheduledAt.IsZero() {
		return errors.New("appointment scheduled date is required")
	}

	if appointment.CreatedAt.IsZero() {
		return errors.New("appointment creation date is required")
	}

	if strings.TrimSpace(appointment.Title) == "" {
		return errors.New("appointment title is required")
	}

	if strings.TrimSpace(appointment.Category) == "" {
		return errors.New("appointment category is required")
	}

	if strings.TrimSpace(appointment.Source) == "" {
		return errors.New("appointment source is required")
	}

	return nil
}

func newAppointment(scheduledAt time.Time, title string, category string, location string, source string) (Appointment, error) {
	appointment := Appointment{
		ScheduledAt: scheduledAt,
		CreatedAt:   time.Now().UTC(),
		Title:       strings.TrimSpace(title),
		Category:    strings.TrimSpace(category),
		Location:    strings.TrimSpace(location),
		Source:      strings.TrimSpace(source),
	}

	if err := validateAppointment(appointment); err != nil {
		return Appointment{}, err
	}

	return appointment, nil
}
