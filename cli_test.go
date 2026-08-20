package main

import (
	"strings"
	"testing"
	"time"
)

func TestAppDescription(t *testing.T) {
	got := appDescription()
	want := "Vitalynq organizes local health data."

	if got != want {
		t.Fatalf("appDescription() = %q, want %q", got, want)
	}
}

func TestHelpText(t *testing.T) {
	got := helpText()
	want := `Vitalynq

  Commands:
    help               Show this help
    version            Show the version
    about              Show the current scope
    privacy            Show privacy information
    limitations        Show Vitalynq limitations
    profile set        Save the medical profile
    profile show       Show the medical profile
    observations list  List observations
    obs list           Alias for observations list
    observations add   Add an observation
    observations add --date YYYY-MM-DD
                       Add a dated observation
    db path            Show the SQLite database path
    db info            Show local storage information
    db check           Check local storage access
    measurements list  List measurements
    measurements add   Add a measurement
    measurements add --date YYYY-MM-DD
                       Add a dated measurement
    appointments list  List appointments
    appointments add   Add an appointment
    summary            Show a summary
    export             Export local data as JSON

  Vitalynq organizes data. It does not provide diagnosis.`

	if got != want {
		t.Fatalf("helpText() = %q, want %q", got, want)
	}
}

func TestAboutText(t *testing.T) {
	got := aboutText()
	want := `Vitalynq is a local CLI application for organizing personal health data.

  Current scope:
  local
  single-user
  no cloud
  no telemetry

  Vitalynq does not provide diagnosis and does not replace a healthcare professional.`

	if got != want {
		t.Fatalf("aboutText() = %q, want %q", got, want)
	}
}

func TestPrivacyText(t *testing.T) {
	got := privacyText()
	want := `Vitalynq Privacy

  Vitalynq stores data locally in a SQLite file.
  Vitalynq does not send data to a server, cloud service, or external API.
  Vitalynq does not use telemetry.

  JSON exports are displayed locally in the terminal.
  The user is responsible for protecting the SQLite file and exports.

  Vitalynq organizes data. It does not provide diagnosis and does not replace a healthcare professional.`

	if got != want {
		t.Fatalf("privacyText() = %q, want %q", got, want)
	}
}

func TestLimitationsText(t *testing.T) {
	got := limitationsText()
	want := `Vitalynq Limitations

  Vitalynq organizes personal health data.
  Vitalynq does not provide diagnosis.
  Vitalynq does not recommend treatments.
  Vitalynq does not predict health changes.
  Vitalynq does not replace a healthcare professional.

  For medical questions, contact a qualified healthcare professional.`

	if got != want {
		t.Fatalf("limitationsText() = %q, want %q", got, want)
	}
}

func TestOutputForArgsLimitations(t *testing.T) {
	got := outputForArgs([]string{"vitalynq", "limitations"}, NewMemoryObservationStore(), NewMemoryMedicalProfileStore(), NewMemoryMeasurementStore(), NewMemoryAppointmentStore(), defaultDatabasePath)
	want := limitationsText()

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestUnknownCommandText(t *testing.T) {
	got := unknownCommandText("unknown")
	want := "Unknown command: unknown\n\nUse 'vitalynq help' to see available commands."

	if got != want {
		t.Fatalf("unknownCommandText() = %q, want %q", got, want)
	}
}

func TestUnknownSubcommandText(t *testing.T) {
	got := unknownSubcommandText("db", "nope")
	want := "Unknown subcommand: db nope\n\nUse 'vitalynq help' to see available commands."

	if got != want {
		t.Fatalf("unknownSubcommandText() = %q, want %q", got, want)
	}
}

func TestInvalidDateText(t *testing.T) {
	got := invalidDateText()
	want := `Invalid date.

  Use the YYYY-MM-DD format, for example: 2026-07-29`

	if got != want {
		t.Fatalf("invalidDateText() = %q, want %q", got, want)
	}
}

func TestMissingObservationText(t *testing.T) {
	got := missingObservationText()
	want := `Missing observation text.

  Usage: vitalynq observations add "Fictive observation"`

	if got != want {
		t.Fatalf("missingObservationText() = %q, want %q", got, want)
	}
}

func TestMissingObservationDateText(t *testing.T) {
	got := missingObservationDateText()
	want := `Missing observation date or text.

  Usage: vitalynq observations add --date YYYY-MM-DD "Fictive observation"`

	if got != want {
		t.Fatalf("missingObservationDateText() = %q, want %q", got, want)
	}
}

func TestMissingMeasurementText(t *testing.T) {
	got := missingMeasurementText()
	want := `Missing measurement arguments.

  Usage: vitalynq measurements add indicator value unit context method source`

	if got != want {
		t.Fatalf("missingMeasurementText() = %q, want %q", got, want)
	}
}

func TestMissingMeasurementDateText(t *testing.T) {
	got := missingMeasurementDateText()
	want := `Missing measurement date or arguments.

  Usage: vitalynq measurements add --date YYYY-MM-DD indicator value unit context method source`

	if got != want {
		t.Fatalf("missingMeasurementDateText() = %q, want %q", got, want)
	}
}

func TestInvalidMeasurementValueText(t *testing.T) {
	got := invalidMeasurementValueText()
	want := `Invalid measurement value.

  The value must be a number, for example: 72.5`

	if got != want {
		t.Fatalf("invalidMeasurementValueText() = %q, want %q", got, want)
	}
}

func TestMissingAppointmentText(t *testing.T) {
	got := missingAppointmentText()
	want := `Missing appointment arguments.

  Usage: vitalynq appointments add YYYY-MM-DD title category location source`

	if got != want {
		t.Fatalf("missingAppointmentText() = %q, want %q", got, want)
	}
}

func TestMissingMedicalProfileLabelText(t *testing.T) {
	got := missingMedicalProfileLabelText()
	want := `Missing medical profile label.

  Usage: vitalynq profile set "Fictive profile"`

	if got != want {
		t.Fatalf("missingMedicalProfileLabelText() = %q, want %q", got, want)
	}
}

func TestOutputForArgsWithoutCommand(t *testing.T) {
	got := outputForArgs([]string{"vitalynq"}, NewMemoryObservationStore(), NewMemoryMedicalProfileStore(), NewMemoryMeasurementStore(), NewMemoryAppointmentStore(), defaultDatabasePath)
	want := "Vitalynq organizes local health data."

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestOutputForArgsHelp(t *testing.T) {
	got := outputForArgs([]string{"vitalynq", "help"}, NewMemoryObservationStore(), NewMemoryMedicalProfileStore(), NewMemoryMeasurementStore(), NewMemoryAppointmentStore(), defaultDatabasePath)
	want := helpText()

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestOutputForArgsVersion(t *testing.T) {
	got := outputForArgs([]string{"vitalynq", "version"}, NewMemoryObservationStore(), NewMemoryMedicalProfileStore(), NewMemoryMeasurementStore(), NewMemoryAppointmentStore(), defaultDatabasePath)
	want := "Vitalynq 0.1.0-dev"

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestOutputForArgsAbout(t *testing.T) {
	got := outputForArgs([]string{"vitalynq", "about"}, NewMemoryObservationStore(), NewMemoryMedicalProfileStore(), NewMemoryMeasurementStore(), NewMemoryAppointmentStore(), defaultDatabasePath)
	want := aboutText()

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestOutputForArgsPrivacy(t *testing.T) {
	got := outputForArgs([]string{"vitalynq", "privacy"}, NewMemoryObservationStore(), NewMemoryMedicalProfileStore(), NewMemoryMeasurementStore(), NewMemoryAppointmentStore(), defaultDatabasePath)
	want := privacyText()

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestOutputForArgsUnknownCommand(t *testing.T) {
	got := outputForArgs([]string{"vitalynq", "unknown"}, NewMemoryObservationStore(), NewMemoryMedicalProfileStore(), NewMemoryMeasurementStore(), NewMemoryAppointmentStore(), defaultDatabasePath)
	want := unknownCommandText("unknown")

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestOutputForArgsUnknownDatabaseSubcommand(t *testing.T) {
	got := outputForArgs([]string{"vitalynq", "db", "nope"}, NewMemoryObservationStore(), NewMemoryMedicalProfileStore(), NewMemoryMeasurementStore(), NewMemoryAppointmentStore(), defaultDatabasePath)
	want := unknownSubcommandText("db", "nope")

	if got != want {
		t.Fatalf("unknownSubcommandText() = %q, want %q", got, want)
	}
}

func TestOutputForArgsUnknownObservationsSubcommand(t *testing.T) {
	got := outputForArgs([]string{"vitalynq", "observations", "nope"}, NewMemoryObservationStore(), NewMemoryMedicalProfileStore(), NewMemoryMeasurementStore(), NewMemoryAppointmentStore(), defaultDatabasePath)
	want := unknownSubcommandText("observations", "nope")

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestObservationsListTextWithoutObservations(t *testing.T) {
	store := NewMemoryObservationStore()

	got := observationsListText(store)
	want := "No observations recorded."

	if got != want {
		t.Fatalf("observationsListText() = %q, want %q", got, want)
	}
}

func TestOutputForArgsObservationsList(t *testing.T) {
	got := outputForArgs([]string{"vitalynq", "observations", "list"}, NewMemoryObservationStore(), NewMemoryMedicalProfileStore(), NewMemoryMeasurementStore(), NewMemoryAppointmentStore(), defaultDatabasePath)
	want := "No observations recorded."

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestObservationsListTextWithObservations(t *testing.T) {
	store := NewMemoryObservationStore()

	if _, err := store.Save(validStoreObservation("Fictive test observation")); err != nil {
		t.Fatalf("Save() error = %v, want nil", err)
	}

	got := observationsListText(store)
	want := `Observations:
- #1 2026-07-17 Fictive test observation`

	if got != want {
		t.Fatalf("observationsListText() = %q, want %q", got, want)
	}
}

func TestObservationsAddText(t *testing.T) {
	store := NewMemoryObservationStore()

	got := observationsAddText(store, "Fictive test observation")
	want := "Observation #1 added."

	if got != want {
		t.Fatalf("observationsAddText() = %q, want %q", got, want)
	}

	observations, err := store.List()
	if err != nil {
		t.Fatalf("List() error = %v, want nil", err)
	}
	if len(observations) != 1 {
		t.Fatalf("len(List()) = %d, want 1", len(observations))
	}

	if observations[0].Text != "Fictive test observation" {
		t.Fatalf("Text = %q, want %q", observations[0].Text, "Fictive test observation")
	}
}

func TestOutputForArgsObservationsAdd(t *testing.T) {
	store := NewMemoryObservationStore()

	got := outputForArgs([]string{"vitalynq", "observations", "add", "Fictive test observation"}, store, NewMemoryMedicalProfileStore(), NewMemoryMeasurementStore(), NewMemoryAppointmentStore(), defaultDatabasePath)
	want := "Observation #1 added."

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestOutputForArgsObservationsAddMissingText(t *testing.T) {
	store := NewMemoryObservationStore()

	got := outputForArgs([]string{"vitalynq", "observations", "add"}, store, NewMemoryMedicalProfileStore(), NewMemoryMeasurementStore(), NewMemoryAppointmentStore(), defaultDatabasePath)
	want := missingObservationText()

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestOutputForArgsObsList(t *testing.T) {
	got := outputForArgs([]string{"vitalynq", "obs", "list"}, NewMemoryObservationStore(), NewMemoryMedicalProfileStore(), NewMemoryMeasurementStore(), NewMemoryAppointmentStore(), defaultDatabasePath)
	want := "No observations recorded."

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestOutputForArgsObsAdd(t *testing.T) {
	store := NewMemoryObservationStore()

	got := outputForArgs([]string{"vitalynq", "obs", "add", "Fictive test observation"}, store, NewMemoryMedicalProfileStore(), NewMemoryMeasurementStore(), NewMemoryAppointmentStore(), defaultDatabasePath)
	want := "Observation #1 added."

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestDatabasePathText(t *testing.T) {
	got := databasePathText("test.db")
	want := "SQLite database: test.db"

	if got != want {
		t.Fatalf("databasePathText() = %q, want %q", got, want)
	}
}

func TestOutputForArgsDatabasePath(t *testing.T) {
	got := outputForArgs([]string{"vitalynq", "db", "path"}, NewMemoryObservationStore(), NewMemoryMedicalProfileStore(), NewMemoryMeasurementStore(), NewMemoryAppointmentStore(), "test.db")
	want := "SQLite database: test.db"

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestDatabaseInfoText(t *testing.T) {
	got := databaseInfoText("test.db")
	want := `SQLite database: test.db
  Storage: local
  Cloud: no
  Telemetry: no`

	if got != want {
		t.Fatalf("databaseInfoText() = %q, want %q", got, want)
	}
}

func TestOutputForArgsDatabaseInfo(t *testing.T) {
	got := outputForArgs([]string{"vitalynq", "db", "info"}, NewMemoryObservationStore(), NewMemoryMedicalProfileStore(), NewMemoryMeasurementStore(), NewMemoryAppointmentStore(), "test.db")
	want := `SQLite database: test.db
  Storage: local
  Cloud: no
  Telemetry: no`

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestDatabaseCheckText(t *testing.T) {
	got := databaseCheckText()
	want := `SQLite database accessible: yes
  SQLite schema initialized: yes`

	if got != want {
		t.Fatalf("databaseCheckText() = %q, want %q", got, want)
	}
}

func TestOutputForArgsDatabaseCheckText(t *testing.T) {
	got := outputForArgs([]string{"vitalynq", "db", "check"}, NewMemoryObservationStore(), NewMemoryMedicalProfileStore(), NewMemoryMeasurementStore(), NewMemoryAppointmentStore(), "test.db")
	want := `SQLite database accessible: yes
  SQLite schema initialized: yes`

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestParseObservationDate(t *testing.T) {
	got, err := parseObservationDate("2026-07-29")
	if err != nil {
		t.Fatalf("parseObservationDate() error = %v, want nil", err)
	}

	want := testTimeFromDate(2026, 7, 29)
	if !got.Equal(want) {
		t.Fatalf("date = %v, want %v", got, want)
	}
}

func TestParseObservationDateRejectsInvalidDate(t *testing.T) {
	_, err := parseObservationDate("29-07-2026")
	if err == nil {
		t.Fatalf("parseObservationDate() error = nil, want error")
	}
}

func TestObservationsAddTextWithDate(t *testing.T) {
	store := NewMemoryObservationStore()

	got := observationsAddTextWithDate(store, "2026-07-29", "Fictive test observation")
	want := "Observation #1 added."

	if got != want {
		t.Fatalf("observationAddTextWithDate() = %q, want %q", got, want)
	}

	observations, err := store.List()
	if err != nil {
		t.Fatalf("List() error = %v, want nil", err)
	}

	if len(observations) != 1 {
		t.Fatalf("len(List()) = %d, want 1", len(observations))
	}

	wantDate := testTimeFromDate(2026, 7, 29)
	if !observations[0].OccurredAt.Equal(wantDate) {
		t.Fatalf("OccurredAt = %v, want %v", observations[0].OccurredAt, wantDate)
	}
}

func TestOutputForArgsObservationsAddWithDate(t *testing.T) {
	store := NewMemoryObservationStore()

	got := outputForArgs([]string{"vitalynq", "observations", "add", "--date", "2026-07-29", "Fictive test observation"}, store, NewMemoryMedicalProfileStore(), NewMemoryMeasurementStore(), NewMemoryAppointmentStore(), defaultDatabasePath)
	want := "Observation #1 added."

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func testTimeFromDate(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

func TestOutputForArgsObservationsAddWithDateMissingValues(t *testing.T) {
	store := NewMemoryObservationStore()

	got := outputForArgs([]string{"vitalynq", "observations", "add", "--date"}, store, NewMemoryMedicalProfileStore(), NewMemoryMeasurementStore(), NewMemoryAppointmentStore(), defaultDatabasePath)
	want := missingObservationDateText()

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestOutputForArgsObservationsAddWithDateMissingText(t *testing.T) {
	store := NewMemoryObservationStore()

	got := outputForArgs([]string{"vitalynq", "observations", "add", "--date", "2026-07-29"}, store, NewMemoryMedicalProfileStore(), NewMemoryMeasurementStore(), NewMemoryAppointmentStore(), defaultDatabasePath)
	want := missingObservationDateText()

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestOutputForArgsObservationsAddWithInvalidDate(t *testing.T) {
	store := NewMemoryObservationStore()

	got := outputForArgs([]string{"vitalynq", "observations", "add", "--date", "29-07-2026", "Fictive test observation"}, store, NewMemoryMedicalProfileStore(), NewMemoryMeasurementStore(), NewMemoryAppointmentStore(), defaultDatabasePath)
	want := invalidDateText()

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestMedicalProfileTextWithoutProfile(t *testing.T) {
	store := NewMemoryMedicalProfileStore()

	got := medicalProfileText(store)
	want := "No medical profile recorded."

	if got != want {
		t.Fatalf("medicalProfileText() = %q, want %q", got, want)
	}
}

func TestMedicalProfileTextWithProfile(t *testing.T) {
	store := NewMemoryMedicalProfileStore()

	if _, err := store.Save(validMedicalProfile("Fictive test profile")); err != nil {
		t.Fatalf("Save() error = %v, want nil", err)
	}

	got := medicalProfileText(store)
	want := "Medical profile: Fictive test profile"

	if got != want {
		t.Fatalf("medicalProfileText() = %q, want %q", got, want)
	}
}

func TestOutputForArgsProfileShow(t *testing.T) {
	profileStore := NewMemoryMedicalProfileStore()

	if _, err := profileStore.Save(validMedicalProfile("Fictive test profile")); err != nil {
		t.Fatalf("Save() error = %v, want nil", err)
	}

	got := outputForArgs([]string{"vitalynq", "profile", "show"}, NewMemoryObservationStore(), profileStore, NewMemoryMeasurementStore(), NewMemoryAppointmentStore(), defaultDatabasePath)
	want := "Medical profile: Fictive test profile"

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestMedicalProfileSaveText(t *testing.T) {
	store := NewMemoryMedicalProfileStore()

	got := medicalProfileSaveText(store, "Fictive test profile")
	want := "Medical profile #1 saved."

	if got != want {
		t.Fatalf("medicalProfileSaveText() = %q, want %q", got, want)
	}
	profile, found, err := store.Get()
	if err != nil {
		t.Fatalf("Get() error = %v, want nil", err)
	}

	if !found {
		t.Fatalf("found = false, want true")
	}

	if profile.Label != "Fictive test profile" {
		t.Fatalf("Label = %q, want %q", profile.Label, "Fictive test profile")
	}
}

func TestMedicalProfileSaveTextRejectsBlankLabel(t *testing.T) {
	store := NewMemoryMedicalProfileStore()

	got := medicalProfileSaveText(store, "  ")
	want := "Unable to save medical profile: medical profile label is required"

	if got != want {
		t.Fatalf("medicalProfileSaveText() = %q, want %q", got, want)
	}
}

func TestOutputForArgsProfilSet(t *testing.T) {
	profileStore := NewMemoryMedicalProfileStore()

	got := outputForArgs(
		[]string{"vitalynq", "profile", "set", "Fictive test profile"},
		NewMemoryObservationStore(),
		profileStore,
		NewMemoryMeasurementStore(),
		NewMemoryAppointmentStore(),
		defaultDatabasePath,
	)
	want := "Medical profile #1 saved."

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestOutputForArgsProfileSetMissingLabel(t *testing.T) {
	got := outputForArgs([]string{"vitalynq", "profile", "set"}, NewMemoryObservationStore(), NewMemoryMedicalProfileStore(), NewMemoryMeasurementStore(), NewMemoryAppointmentStore(), defaultDatabasePath)
	want := missingMedicalProfileLabelText()

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestMeasurementsListTextWithoutMeasurements(t *testing.T) {
	store := NewMemoryMeasurementStore()

	got := measurementsListText(store)
	want := "No measurements recorded."

	if got != want {
		t.Fatalf("measurementsListText() %q, want %q", got, want)
	}
}

func TestMeasurementsListTextWithMeasurements(t *testing.T) {
	store := NewMemoryMeasurementStore()

	if _, err := store.Save(validCLIMeasurement()); err != nil {
		t.Fatalf("Save() error = %v, want nil", err)
	}

	got := measurementsListText(store)
	want := `Measurements:
- #1 2026-07-17 weight 72.50 kg`

	if got != want {
		t.Fatalf("measurementsListText() = %q, want %q", got, want)
	}
}

func TestOutputForArgsMeasurementsList(t *testing.T) {
	got := outputForArgs(
		[]string{"vitalynq", "measurements", "list"},
		NewMemoryObservationStore(),
		NewMemoryMedicalProfileStore(),
		NewMemoryMeasurementStore(),
		NewMemoryAppointmentStore(),
		defaultDatabasePath,
	)
	want := "No measurements recorded."

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestMeasurementsAddText(t *testing.T) {
	store := NewMemoryMeasurementStore()

	got := measurementsAddText(
		store,
		"weight",
		72.5,
		"kg",
		"fictive test",
		"manual entry",
		"manual entry",
	)
	want := "Measurement #1 added."

	if got != want {
		t.Fatalf("measurementsAddText() = %q, want %q", got, want)
	}

	measurements, err := store.List()
	if err != nil {
		t.Fatalf("List() error = %v, want nil", err)
	}

	if len(measurements) != 1 {
		t.Fatalf("len(List()) = %d, want 1", len(measurements))
	}

	if measurements[0].Indicator != "weight" {
		t.Fatalf("Indicator = %q, want %q", measurements[0].Indicator, "weight")
	}

	if measurements[0].Unit != "kg" {
		t.Fatalf("Unit = %q, want %q", measurements[0].Unit, "kg")
	}
}

func TestOutputForArgsMeasurementAddMissingArguments(t *testing.T) {
	got := outputForArgs(
		[]string{"vitalynq", "measurements", "add"},
		NewMemoryObservationStore(),
		NewMemoryMedicalProfileStore(),
		NewMemoryMeasurementStore(),
		NewMemoryAppointmentStore(),
		defaultDatabasePath,
	)
	want := missingMeasurementText()

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestOutputForArgsMeasurementAddWithDateMissingArguments(t *testing.T) {
	got := outputForArgs(
		[]string{"vitalynq", "measurements", "add", "--date"},
		NewMemoryObservationStore(),
		NewMemoryMedicalProfileStore(),
		NewMemoryMeasurementStore(),
		NewMemoryAppointmentStore(),
		defaultDatabasePath,
	)
	want := missingMeasurementDateText()

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestMeasurementsAddTextRejectsMissingUnit(t *testing.T) {
	store := NewMemoryMeasurementStore()

	got := measurementsAddText(
		store,
		"weight",
		72.5,
		"	",
		"fictive test",
		"manual entry",
		"manual entry",
	)
	want := "Unable to add measurement: measurement unit is required"

	if got != want {
		t.Fatalf("measurementsAddText() = %q, want %q", got, want)
	}
}

func TestOutputForArgsMeasurementsAdd(t *testing.T) {
	measurementStore := NewMemoryMeasurementStore()

	got := outputForArgs(
		[]string{"vitalynq", "measurements", "add", "weight", "72.5", "kg", "fictive test", "manual entry", "manual entry"},
		NewMemoryObservationStore(),
		NewMemoryMedicalProfileStore(),
		measurementStore,
		NewMemoryAppointmentStore(),
		defaultDatabasePath,
	)
	want := "Measurement #1 added."

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestOutputForArgsMeasurementsAddRejectsInvalidValue(t *testing.T) {
	got := outputForArgs(
		[]string{"vitalynq", "measurements", "add", "weight", "abc", "kg", "fictive test", "manual entry", "manual entry"},
		NewMemoryObservationStore(),
		NewMemoryMedicalProfileStore(),
		NewMemoryMeasurementStore(),
		NewMemoryAppointmentStore(),
		defaultDatabasePath,
	)
	want := invalidMeasurementValueText()

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestMeasurementsAddTextWithDate(t *testing.T) {
	store := NewMemoryMeasurementStore()

	got := measurementsAddTextWithDate(
		store,
		"2026-07-29",
		"weight",
		72.5,
		"kg",
		"fictive test",
		"manual entry",
		"manual entry",
	)
	want := "Measurement #1 added."

	if got != want {
		t.Fatalf("measurementsAddTextWithDate() = %q, want %q", got, want)
	}

	measurements, err := store.List()
	if err != nil {
		t.Fatalf("List() error = %v, want nil", err)
	}

	wantDate := testTimeFromDate(2026, 7, 29)
	if !measurements[0].OccurredAt.Equal(wantDate) {
		t.Fatalf("OccurredAt = %v, want %v", measurements[0].OccurredAt, wantDate)
	}
}

func TestOutputForArgsMeasurementsAddWithDate(t *testing.T) {
	measurementStore := NewMemoryMeasurementStore()

	got := outputForArgs(
		[]string{"vitalynq", "measurements", "add", "--date", "2026-07-29", "weight", "72.5", "kg", "fictive test", "manual entry", "manual entry"},
		NewMemoryObservationStore(),
		NewMemoryMedicalProfileStore(),
		measurementStore,
		NewMemoryAppointmentStore(),
		defaultDatabasePath,
	)
	want := "Measurement #1 added."

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestAppointmentsListTextWithoutAppointments(t *testing.T) {
	store := NewMemoryAppointmentStore()

	got := appointmentsListText(store)
	want := "No appointments recorded."

	if got != want {
		t.Fatalf("Save() error = %q, want %q", got, want)
	}
}

func TestAppointmentsListTextWithAppointments(t *testing.T) {
	store := NewMemoryAppointmentStore()

	if _, err := store.Save(validCLIAppointment()); err != nil {
		t.Fatalf("Save() error = %v, want nil", err)
	}

	got := appointmentsListText(store)
	want := `Appointments:
- #1 2026-07-17 fictive consultation (appointment)`

	if got != want {
		t.Fatalf("appointmentsListText() = %q, want %q", got, want)
	}
}

func TestAppointmentsAddText(t *testing.T) {
	store := NewMemoryAppointmentStore()

	got := appointmentsAddText(
		store,
		"2026-07-29",
		"fictive consultation",
		"appointment",
		"fictive office",
		"manual entry",
	)
	want := "Appointment #1 added."

	if got != want {
		t.Fatalf("appointmentsAddText() = %q, want %q", got, want)
	}

	appointments, err := store.List()
	if err != nil {
		t.Fatalf("List() error = %v, want nil", err)
	}

	if len(appointments) != 1 {
		t.Fatalf("len(List()) = %d, want 1", len(appointments))
	}

	if appointments[0].Title != "fictive consultation" {
		t.Fatalf("Title = %q, want %q", appointments[0].Title, "fictive consultation")
	}
}

func TestAppointmentsAddTextRejectsInvalidDate(t *testing.T) {
	store := NewMemoryAppointmentStore()

	got := appointmentsAddText(
		store,
		"29-07-2026",
		"fictive consultation",
		"appointment",
		"fictive office",
		"manual entry",
	)
	want := invalidDateText()

	if got != want {
		t.Fatalf("appointmentsAddText() = %q, want %q", got, want)
	}
}

func TestOutputForArgsAppointmentsAdd(t *testing.T) {
	appointmentStore := NewMemoryAppointmentStore()

	got := outputForArgs(
		[]string{"vitalynq", "appointments", "add", "2026-07-29", "fictive consultation", "appointment", "fictive office", "manual entry"},
		NewMemoryObservationStore(),
		NewMemoryMedicalProfileStore(),
		NewMemoryMeasurementStore(),
		appointmentStore,
		defaultDatabasePath,
	)
	want := "Appointment #1 added."

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestOutputForArgsAppointmentsAddMissingArguments(t *testing.T) {
	got := outputForArgs(
		[]string{"vitalynq", "appointments", "add"},
		NewMemoryObservationStore(),
		NewMemoryMedicalProfileStore(),
		NewMemoryMeasurementStore(),
		NewMemoryAppointmentStore(),
		defaultDatabasePath,
	)
	want := missingAppointmentText()

	if got != want {
		t.Fatalf("outputForArgs() %q, want %q", got, want)
	}
}

func TestSummaryText(t *testing.T) {
	observationStore := NewMemoryObservationStore()
	measurementStore := NewMemoryMeasurementStore()
	appointmentStore := NewMemoryAppointmentStore()

	if _, err := observationStore.Save(validStoreObservation("Fictive observation")); err != nil {
		t.Fatalf("Save(observation) error = %v, want nil", err)
	}

	if _, err := measurementStore.Save(validCLIMeasurement()); err != nil {
		t.Fatalf("Save(measurement) error = %v, want nil", err)
	}

	if _, err := appointmentStore.Save(validCLIAppointment()); err != nil {
		t.Fatalf("Save(appointment) error = %v, want nil", err)
	}

	got := summaryText(observationStore, measurementStore, appointmentStore)
	want := `Summary:
- Observations: 1
- Measurements: 1
- Appointments: 1`

	if got != want {
		t.Fatalf("summaryText() = %q, want %q", got, want)
	}
}

func TestOutputForArgsSummary(t *testing.T) {
	observationStore := NewMemoryObservationStore()
	measurementStore := NewMemoryMeasurementStore()
	appointmentStore := NewMemoryAppointmentStore()

	if _, err := observationStore.Save(validStoreObservation("Fictive observation")); err != nil {
		t.Fatalf("Save(observation) error = %v, want nil", err)
	}

	got := outputForArgs(
		[]string{"vitalynq", "summary"},
		observationStore,
		NewMemoryMedicalProfileStore(),
		measurementStore,
		appointmentStore,
		defaultDatabasePath,
	)
	want := `Summary:
- Observations: 1
- Measurements: 0
- Appointments: 0`

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestOutputForArgsExport(t *testing.T) {
	profileStore := NewMemoryMedicalProfileStore()
	observationStore := NewMemoryObservationStore()
	measurementStore := NewMemoryMeasurementStore()
	appointmentStore := NewMemoryAppointmentStore()

	if _, err := observationStore.Save(validStoreObservation("Fictive observation")); err != nil {
		t.Fatalf("Save(observation) error = %v, want nil", err)
	}

	got := outputForArgs(
		[]string{"vitalynq", "export"},
		observationStore,
		profileStore,
		measurementStore,
		appointmentStore,
		defaultDatabasePath,
	)

	assertContains(t, got, `"profile"`)
	assertContains(t, got, `"observations"`)
	assertContains(t, got, `"measurements"`)
	assertContains(t, got, `"appointments"`)
	assertContains(t, got, `"Fictive observation"`)
}

func validCLIMeasurement() Measurement {
	return Measurement{
		OccurredAt: testTime(),
		CreatedAt:  testTime(),
		Indicator:  "weight",
		Value:      72.5,
		Unit:       "kg",
		Context:    "fictive test",
		Method:     "manual entry",
		Source:     "manual entry",
	}
}

func validCLIAppointment() Appointment {
	return Appointment{
		ScheduledAt: testTime(),
		CreatedAt:   testTime(),
		Title:       "fictive consultation",
		Category:    "appointment",
		Location:    "fictive office",
		Source:      "manual entry",
	}
}

func assertContains(t *testing.T, got string, want string) {
	t.Helper()

	if !strings.Contains(got, want) {
		t.Fatalf("got %q, want it to contain %q", got, want)
	}
}
