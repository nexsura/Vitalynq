package main

import (
	"strings"
	"testing"
	"time"
)

func TestAppDescription(t *testing.T) {
	got := appDescription()
	want := "Vitalynq organise des données de santé locales."

	if got != want {
		t.Fatalf("appDescription() = %q, want %q", got, want)
	}
}

func TestHelpText(t *testing.T) {
	got := helpText()
	want := `Vitalynq

Commandes:
  help               Affiche cette aide
  version            Affiche la version
  about              Affiche le périmètre actuel
  privacy            Affiche les informations de confidentialité
  profile set        Enregistre le profil médical
  profile show       Affiche le profil médical
  observations list  Liste les observations
  obs list           Alias de observations list
  observations add   Ajoute une observation
  observations add --date YYYY-MM-DD
                     Ajoute une observation datée
  db path            Affiche le chemin de la base SQLite
  measurements list  Liste les mesures
  measurements add   Ajoute une mesure
  measurements add --date YYYY-MM-DD
                     Ajoute une mesure datée
  appointments list  Liste les rendez-vous
  appointments add   Ajoute un rendez-vous
  summary            Affiche un bilan synthétique
  export             Exporte les données locales en JSON

Vitalynq organise des données. Il ne pose pas de diagnostic.`

	if got != want {
		t.Fatalf("helpText() = %q, want %q", got, want)
	}
}

func TestAboutText(t *testing.T) {
	got := aboutText()
	want := `Vitalynq est une application CLI locale pour organiser des données personnelles de santé.

Périmètre actuel:
	local
	mono-utilisateur
	sans cloud
	sans télémétrie

Vitalynq ne pose pas de diagnostic et ne remplace pas un professionnel de santé.`

	if got != want {
		t.Fatalf("aboutText() = %q, want %q", got, want)
	}
}

func TestPrivacyText(t *testing.T) {
	got := privacyText()
	want := `Confidentialité Vitalynq

Vitalynq stocke les données localement dans un fichier SQLite.
Vitalynq n'envoie aucune donnée vers un serveur, un cloud ou une API externe.
Vitalynq n'utilise pas de télémétrie.

Les exports JSON sont affichés localement dans le terminal.
L'utilisateur reste responsable de la protection du fichier SQLite et des exports.

Vitalynq organise des données. Il ne pose pas de diagnostic et ne remplace pas un professionnel de santé.`

	if got != want {
		t.Fatalf("privacyText() = %q, want %q", got, want)
	}
}

func TestUnknownCommandText(t *testing.T) {
	got := unknownCommandText("profil")
	want := "Commande inconnue: profil\n\nUtilisez 'vitalynq help' pour voir les commandes disponibles."

	if got != want {
		t.Fatalf("unknownCommandText() = %q, want %q", got, want)
	}
}

func TestOutputForArgsWithoutCommand(t *testing.T) {
	got := outputForArgs([]string{"vitalynq"}, NewMemoryObservationStore(), NewMemoryMedicalProfileStore(), NewMemoryMeasurementStore(), NewMemoryAppointmentStore(), defaultDatabasePath)
	want := "Vitalynq organise des données de santé locales."

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
	got := outputForArgs([]string{"vitalynq", "profil"}, NewMemoryObservationStore(), NewMemoryMedicalProfileStore(), NewMemoryMeasurementStore(), NewMemoryAppointmentStore(), defaultDatabasePath)
	want := unknownCommandText("profil")

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestObservationsListTextWithoutObservations(t *testing.T) {
	store := NewMemoryObservationStore()

	got := observationsListText(store)
	want := "Aucune observation enregistrée."

	if got != want {
		t.Fatalf("observationsListText() = %q, want %q", got, want)
	}
}

func TestOutputForArgsObservationsList(t *testing.T) {
	got := outputForArgs([]string{"vitalynq", "observations", "list"}, NewMemoryObservationStore(), NewMemoryMedicalProfileStore(), NewMemoryMeasurementStore(), NewMemoryAppointmentStore(), defaultDatabasePath)
	want := "Aucune observation enregistrée."

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestObservationsListTextWithObservations(t *testing.T) {
	store := NewMemoryObservationStore()

	if _, err := store.Save(validStoreObservation("Observation fictive de test")); err != nil {
		t.Fatalf("Save() error = %v, want nil", err)
	}

	got := observationsListText(store)
	want := `Observations:
- #1 2026-07-17 Observation fictive de test`

	if got != want {
		t.Fatalf("observationsListText() = %q, want %q", got, want)
	}
}

func TestObservationsAddText(t *testing.T) {
	store := NewMemoryObservationStore()

	got := observationsAddText(store, "Observation fictive de test")
	want := "Observation #1 ajoutée."

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

	if observations[0].Text != "Observation fictive de test" {
		t.Fatalf("Text = %q, want %q", observations[0].Text, "Observation fictive de test")
	}
}

func TestOutputForArgsObservationsAdd(t *testing.T) {
	store := NewMemoryObservationStore()

	got := outputForArgs([]string{"vitalynq", "observations", "add", "Observation fictive de test"}, store, NewMemoryMedicalProfileStore(), NewMemoryMeasurementStore(), NewMemoryAppointmentStore(), defaultDatabasePath)
	want := "Observation #1 ajoutée."

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestOutputForArgsObservationsAddMissingText(t *testing.T) {
	store := NewMemoryObservationStore()

	got := outputForArgs([]string{"vitalynq", "observations", "add"}, store, NewMemoryMedicalProfileStore(), NewMemoryMeasurementStore(), NewMemoryAppointmentStore(), defaultDatabasePath)
	want := "Texte d'observation manquant."

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestOutputForArgsObsList(t *testing.T) {
	got := outputForArgs([]string{"vitalynq", "obs", "list"}, NewMemoryObservationStore(), NewMemoryMedicalProfileStore(), NewMemoryMeasurementStore(), NewMemoryAppointmentStore(), defaultDatabasePath)
	want := "Aucune observation enregistrée."

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestOutputForArgsObsAdd(t *testing.T) {
	store := NewMemoryObservationStore()

	got := outputForArgs([]string{"vitalynq", "obs", "add", "Observation fictive de test"}, store, NewMemoryMedicalProfileStore(), NewMemoryMeasurementStore(), NewMemoryAppointmentStore(), defaultDatabasePath)
	want := "Observation #1 ajoutée."

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestDatabasePathText(t *testing.T) {
	got := databasePathText("test.db")
	want := "Base SQLite: test.db"

	if got != want {
		t.Fatalf("databasePathText() = %q, want %q", got, want)
	}
}

func TestOutputForArgsDatabasePath(t *testing.T) {
	got := outputForArgs([]string{"vitalynq", "db", "path"}, NewMemoryObservationStore(), NewMemoryMedicalProfileStore(), NewMemoryMeasurementStore(), NewMemoryAppointmentStore(), "test.db")
	want := "Base SQLite: test.db"

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

	got := observationsAddTextWithDate(store, "2026-07-29", "Observation fictive de test")
	want := "Observation #1 ajoutée."

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

	got := outputForArgs([]string{"vitalynq", "observations", "add", "--date", "2026-07-29", "Observation fictive de test"}, store, NewMemoryMedicalProfileStore(), NewMemoryMeasurementStore(), NewMemoryAppointmentStore(), defaultDatabasePath)
	want := "Observation #1 ajoutée."

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
	want := "Date ou texte d'observation manquant."

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestOutputForArgsObservationsAddWithDateMissingText(t *testing.T) {
	store := NewMemoryObservationStore()

	got := outputForArgs([]string{"vitalynq", "observations", "add", "--date", "2026-07-29"}, store, NewMemoryMedicalProfileStore(), NewMemoryMeasurementStore(), NewMemoryAppointmentStore(), defaultDatabasePath)
	want := "Date ou texte d'observation manquant."

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestOutputForArgsObservationsAddWithInvalidDate(t *testing.T) {
	store := NewMemoryObservationStore()

	got := outputForArgs([]string{"vitalynq", "observations", "add", "--date", "29-07-2026", "Observation fictive de test"}, store, NewMemoryMedicalProfileStore(), NewMemoryMeasurementStore(), NewMemoryAppointmentStore(), defaultDatabasePath)
	want := "date invalide, utilisez YYYY-MM-DD"

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestMedicalProfileTextWithoutProfile(t *testing.T) {
	store := NewMemoryMedicalProfileStore()

	got := medicalProfileText(store)
	want := "Aucun profil médical enregistré."

	if got != want {
		t.Fatalf("medicalProfileText() = %q, want %q", got, want)
	}
}

func TestMedicalProfileTextWithProfile(t *testing.T) {
	store := NewMemoryMedicalProfileStore()

	if _, err := store.Save(validMedicalProfile("Profil fictif de test")); err != nil {
		t.Fatalf("Save() error = %v, want nil", err)
	}

	got := medicalProfileText(store)
	want := "Profil médical: Profil fictif de test"

	if got != want {
		t.Fatalf("medicalProfileText() = %q, want %q", got, want)
	}
}

func TestOutputForArgsProfileShow(t *testing.T) {
	profileStore := NewMemoryMedicalProfileStore()

	if _, err := profileStore.Save(validMedicalProfile("Profil fictif de test")); err != nil {
		t.Fatalf("Save() error = %v, want nil", err)
	}

	got := outputForArgs([]string{"vitalynq", "profile", "show"}, NewMemoryObservationStore(), profileStore, NewMemoryMeasurementStore(), NewMemoryAppointmentStore(), defaultDatabasePath)
	want := "Profil médical: Profil fictif de test"

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestMedicalProfileSaveText(t *testing.T) {
	store := NewMemoryMedicalProfileStore()

	got := medicalProfileSaveText(store, "Profil fictif de test")
	want := "Profil médical #1 enregistré."

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

	if profile.Label != "Profil fictif de test" {
		t.Fatalf("Label = %q, want %q", profile.Label, "Profil fictif de test")
	}
}

func TestMedicalProfileSaveTextRejectsBlankLabel(t *testing.T) {
	store := NewMemoryMedicalProfileStore()

	got := medicalProfileSaveText(store, "  ")
	want := "Impossible d'enregistrer le profil médical: medical profile label is required"

	if got != want {
		t.Fatalf("medicalProfileSaveText() = %q, want %q", got, want)
	}
}

func TestOutputForArgsProfilSet(t *testing.T) {
	profileStore := NewMemoryMedicalProfileStore()

	got := outputForArgs(
		[]string{"vitalynq", "profile", "set", "Profil fictif de test"},
		NewMemoryObservationStore(),
		profileStore,
		NewMemoryMeasurementStore(),
		NewMemoryAppointmentStore(),
		defaultDatabasePath,
	)
	want := "Profil médical #1 enregistré."

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestOutputForArgsProfileSetMissingLabel(t *testing.T) {
	got := outputForArgs([]string{"vitalynq", "profile", "set"}, NewMemoryObservationStore(), NewMemoryMedicalProfileStore(), NewMemoryMeasurementStore(), NewMemoryAppointmentStore(), defaultDatabasePath)
	want := "Libellé du profil médical manquant."

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestMeasurementsListTextWithoutMeasurements(t *testing.T) {
	store := NewMemoryMeasurementStore()

	got := measurementsListText(store)
	want := "Aucune mesure enregistrée."

	if got != want {
		t.Fatalf("measurementsListText() %q, want %q", got, want)
	}
}

func TestMeasurementsListTextWithMeasurements(t *testing.T) {
	store := NewMemoryMeasurementStore()

	if _, err := store.Save(validMeasurement()); err != nil {
		t.Fatalf("Save() error = %v, want nil", err)
	}

	got := measurementsListText(store)
	want := `Mesures:
- #1 2026-07-17 poids 72.50 kg`

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
	want := "Aucune mesure enregistrée."

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestMeasurementsAddText(t *testing.T) {
	store := NewMemoryMeasurementStore()

	got := measurementsAddText(
		store,
		"poids",
		72.5,
		"kg",
		"test fictif",
		"saisie manuelle",
		"saisie manuelle",
	)
	want := "Mesure #1 ajoutée."

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

	if measurements[0].Indicator != "poids" {
		t.Fatalf("Indicator = %q, want %q", measurements[0].Indicator, "poids")
	}

	if measurements[0].Unit != "kg" {
		t.Fatalf("Unit = %q, want %q", measurements[0].Unit, "kg")
	}
}

func TestMeasurementsAddTextRejectsMissingUnit(t *testing.T) {
	store := NewMemoryMeasurementStore()

	got := measurementsAddText(
		store,
		"poids",
		72.5,
		"	",
		"test fictif",
		"saisie manuelle",
		"saisie manuelle",
	)
	want := "Impossible d'ajouter la mesure: measurement unit is required"

	if got != want {
		t.Fatalf("measurementsAddText() = %q, want %q", got, want)
	}
}

func TestOutputForArgsMeasurementsAdd(t *testing.T) {
	measurementStore := NewMemoryMeasurementStore()

	got := outputForArgs(
		[]string{"vitalynq", "measurements", "add", "poids", "72.5", "kg", "test fictif", "saisie manuelle", "saisie manuelle"},
		NewMemoryObservationStore(),
		NewMemoryMedicalProfileStore(),
		measurementStore,
		NewMemoryAppointmentStore(),
		defaultDatabasePath,
	)
	want := "Mesure #1 ajoutée."

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestOutputForArgsMeasurementsAddRejectsInvalidValue(t *testing.T) {
	got := outputForArgs(
		[]string{"vitalynq", "measurements", "add", "poids", "abc", "kg", "test fictif", "saisie manuelle", "saisie manuelle"},
		NewMemoryObservationStore(),
		NewMemoryMedicalProfileStore(),
		NewMemoryMeasurementStore(),
		NewMemoryAppointmentStore(),
		defaultDatabasePath,
	)
	want := "Valeur de mesure invalide."

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestMeasurementsAddTextWithDate(t *testing.T) {
	store := NewMemoryMeasurementStore()

	got := measurementsAddTextWithDate(
		store,
		"2026-07-29",
		"poids",
		72.5,
		"kg",
		"test fictif",
		"saisie manuelle",
		"saisie manuelle",
	)
	want := "Mesure #1 ajoutée."

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
		[]string{"vitalynq", "measurements", "add", "--date", "2026-07-29", "poids", "72.5", "kg", "test fictif", "saisie manuelle", "saisie manuelle"},
		NewMemoryObservationStore(),
		NewMemoryMedicalProfileStore(),
		measurementStore,
		NewMemoryAppointmentStore(),
		defaultDatabasePath,
	)
	want := "Mesure #1 ajoutée."

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestAppointmentsListTextWithoutAppointments(t *testing.T) {
	store := NewMemoryAppointmentStore()

	got := appointmentsListText(store)
	want := "Aucun rendez-vous enregistré."

	if got != want {
		t.Fatalf("Save() error = %q, want %q", got, want)
	}
}

func TestAppointmentsListTextWithAppointments(t *testing.T) {
	store := NewMemoryAppointmentStore()

	if _, err := store.Save(validAppointment()); err != nil {
		t.Fatalf("Save() error = %v, want nil", err)
	}

	got := appointmentsListText(store)
	want := `Rendez-vous:
- #1 2026-07-17 consultation fictive (rendez-vous)`

	if got != want {
		t.Fatalf("appointmentsListText() = %q, want %q", got, want)
	}
}

func TestAppointmentsAddText(t *testing.T) {
	store := NewMemoryAppointmentStore()

	got := appointmentsAddText(
		store,
		"2026-07-29",
		"consultation fictive",
		"rendez-vous",
		"cabinet fictif",
		"saisie manuelle",
	)
	want := "Rendez-vous #1 ajouté."

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

	if appointments[0].Title != "consultation fictive" {
		t.Fatalf("Title = %q, want %q", appointments[0].Title, "consultation fictive")
	}
}

func TestAppointmentsAddTextRejectsInvalidDate(t *testing.T) {
	store := NewMemoryAppointmentStore()

	got := appointmentsAddText(
		store,
		"29-07-2026",
		"consultation fictive",
		"rendez-vous",
		"cabinet fictif",
		"saisie manuelle",
	)
	want := "date invalide, utilisez YYYY-MM-DD"

	if got != want {
		t.Fatalf("appointmentsAddText() = %q, want %q", got, want)
	}
}

func TestOutputForArgsAppointmentsAdd(t *testing.T) {
	AppointmentStore := NewMemoryAppointmentStore()

	got := outputForArgs(
		[]string{"vitalynq", "appointments", "add", "2026-07-29", "consultation fictive", "rendez-vous", "cabinet fictif", "saisie manuelle"},
		NewMemoryObservationStore(),
		NewMemoryMedicalProfileStore(),
		NewMemoryMeasurementStore(),
		AppointmentStore,
		defaultDatabasePath,
	)
	want := "Rendez-vous #1 ajouté."

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
	want := "Arguments de rendez-vous manquants."

	if got != want {
		t.Fatalf("outputForArgs() %q, want %q", got, want)
	}
}

func TestSummaryText(t *testing.T) {
	observationStore := NewMemoryObservationStore()
	measurementStore := NewMemoryMeasurementStore()
	appointmentStore := NewMemoryAppointmentStore()

	if _, err := observationStore.Save(validStoreObservation("Observation fictive")); err != nil {
		t.Fatalf("Save(observation) error = %v, want nil", err)
	}

	if _, err := measurementStore.Save(validMeasurement()); err != nil {
		t.Fatalf("Save(measurement) error = %v, want nil", err)
	}

	if _, err := appointmentStore.Save(validAppointment()); err != nil {
		t.Fatalf("Save(appointment) error = %v, want nil", err)
	}

	got := summaryText(observationStore, measurementStore, appointmentStore)
	want := `Bilan:
- Observations: 1
- Mesures: 1
- Rendez-vous: 1`

	if got != want {
		t.Fatalf("summaryText() = %q, want %q", got, want)
	}
}

func TestOutputForArgsSummary(t *testing.T) {
	observationStore := NewMemoryObservationStore()
	measurementStore := NewMemoryMeasurementStore()
	appointmentStore := NewMemoryAppointmentStore()

	if _, err := observationStore.Save(validStoreObservation("Observation fictive")); err != nil {
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
	want := `Bilan:
- Observations: 1
- Mesures: 0
- Rendez-vous: 0`

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestOutputForArgsExport(t *testing.T) {
	profileStore := NewMemoryMedicalProfileStore()
	observationStore := NewMemoryObservationStore()
	measurementStore := NewMemoryMeasurementStore()
	appointmentStore := NewMemoryAppointmentStore()

	if _, err := observationStore.Save(validStoreObservation("Observation fictive")); err != nil {
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
	assertContains(t, got, `"Observation fictive"`)
}

func assertContains(t *testing.T, got string, want string) {
	t.Helper()

	if !strings.Contains(got, want) {
		t.Fatalf("got %q, want it to contain %q", got, want)
	}
}
