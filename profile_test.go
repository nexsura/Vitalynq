package main

import (
	"testing"
)

func TestValidateMedicalProfileAcceptsValidProfile(t *testing.T) {
	profile := MedicalProfile{
		CreatedAt: testTime(),
		UpdatedAt: testTime(),
		Label:     "Fictive test profile",
	}

	if err := validateMedicalProfile(profile); err != nil {
		t.Fatalf("validateMedicalProfile() error = %v, want nil", err)
	}
}

func TestValidateMedicalProfileRejectsMissingCreatedAt(t *testing.T) {
	profile := MedicalProfile{
		UpdatedAt: testTime(),
		Label:     "Fictive test profile",
	}

	if err := validateMedicalProfile(profile); err == nil {
		t.Fatalf("validateMedicalProfile() error = nil, want error")
	}
}

func TestValidateMedicalProfileRejectsMissingUpdatedAt(t *testing.T) {
	profile := MedicalProfile{
		CreatedAt: testTime(),
		Label:     "Fictive test profile",
	}

	if err := validateMedicalProfile(profile); err == nil {
		t.Fatalf("validateMedicalProfile() error = nil, want error")
	}
}

func TestValidateMedicalProfileRejectsBlankLabel(t *testing.T) {
	profile := MedicalProfile{
		CreatedAt: testTime(),
		UpdatedAt: testTime(),
		Label:     "   ",
	}

	if err := validateMedicalProfile(profile); err == nil {
		t.Fatalf("validateMedicalProfile() error = nil, want error")
	}
}

func TestNewMedicalProfileCreatesValidProfile(t *testing.T) {
	profile, err := newMedicalProfile("Fictive test profile")
	if err != nil {
		t.Fatalf("newMedicalProfile() error = %v, want nil", err)
	}

	if profile.Label != "Fictive test profile" {
		t.Fatalf("Label = %q, want %q", profile.Label, "Fictive test profile")
	}

	if profile.CreatedAt.IsZero() {
		t.Fatalf("CreatedAt is zero, want creation date")
	}

	if profile.UpdatedAt.IsZero() {
		t.Fatalf("UpdatedAt is zero want updated date")
	}

	if profile.ID != 0 {
		t.Fatalf("ID = %d, want 0", profile.ID)
	}
}
