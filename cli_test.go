package main

import (
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
  profile set        Enregistre le profil médical
  profile show       Affiche le profil médical
  observations list  Liste les observations
  obs list           Alias de observations list
  observations add   Ajoute une observation
  observations add --date YYYY-MM-DD
                     Ajoute une observation datée
  db path            Affiche le chemin de la base SQLite

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

func TestUnknownCommandText(t *testing.T) {
	got := unknownCommandText("profil")
	want := "Commande inconnue: profil\n\nUtilisez 'vitalynq help' pour voir les commandes disponibles."

	if got != want {
		t.Fatalf("unknownCommandText() = %q, want %q", got, want)
	}
}

func TestOutputForArgsWithoutCommand(t *testing.T) {
	got := outputForArgs([]string{"vitalynq"}, NewMemoryObservationStore(), NewMemoryMedicalProfileStore(), NewMemoryMeasurementStore(), defaultDatabasePath)
	want := "Vitalynq organise des données de santé locales."

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestOutputForArgsHelp(t *testing.T) {
	got := outputForArgs([]string{"vitalynq", "help"}, NewMemoryObservationStore(), NewMemoryMedicalProfileStore(), NewMemoryMeasurementStore(), defaultDatabasePath)
	want := helpText()

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestOutputForArgsVersion(t *testing.T) {
	got := outputForArgs([]string{"vitalynq", "version"}, NewMemoryObservationStore(), NewMemoryMedicalProfileStore(), NewMemoryMeasurementStore(), defaultDatabasePath)
	want := "Vitalynq 0.1.0-dev"

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestOutputForArgsAbout(t *testing.T) {
	got := outputForArgs([]string{"vitalynq", "about"}, NewMemoryObservationStore(), NewMemoryMedicalProfileStore(), NewMemoryMeasurementStore(), defaultDatabasePath)
	want := aboutText()

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestOutputForArgsUnknownCommand(t *testing.T) {
	got := outputForArgs([]string{"vitalynq", "profil"}, NewMemoryObservationStore(), NewMemoryMedicalProfileStore(), NewMemoryMeasurementStore(), defaultDatabasePath)
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
	got := outputForArgs([]string{"vitalynq", "observations", "list"}, NewMemoryObservationStore(), NewMemoryMedicalProfileStore(), NewMemoryMeasurementStore(), defaultDatabasePath)
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

	got := outputForArgs([]string{"vitalynq", "observations", "add", "Observation fictive de test"}, store, NewMemoryMedicalProfileStore(), NewMemoryMeasurementStore(), defaultDatabasePath)
	want := "Observation #1 ajoutée."

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestOutputForArgsObservationsAddMissingText(t *testing.T) {
	store := NewMemoryObservationStore()

	got := outputForArgs([]string{"vitalynq", "observations", "add"}, store, NewMemoryMedicalProfileStore(), NewMemoryMeasurementStore(), defaultDatabasePath)
	want := "Texte d'observation manquant."

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestOutputForArgsObsList(t *testing.T) {
	got := outputForArgs([]string{"vitalynq", "obs", "list"}, NewMemoryObservationStore(), NewMemoryMedicalProfileStore(), NewMemoryMeasurementStore(), defaultDatabasePath)
	want := "Aucune observation enregistrée."

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestOutputForArgsObsAdd(t *testing.T) {
	store := NewMemoryObservationStore()

	got := outputForArgs([]string{"vitalynq", "obs", "add", "Observation fictive de test"}, store, NewMemoryMedicalProfileStore(), NewMemoryMeasurementStore(), defaultDatabasePath)
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
	got := outputForArgs([]string{"vitalynq", "db", "path"}, NewMemoryObservationStore(), NewMemoryMedicalProfileStore(), NewMemoryMeasurementStore(), "test.db")
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

	got := outputForArgs([]string{"vitalynq", "observations", "add", "--date", "2026-07-29", "Observation fictive de test"}, store, NewMemoryMedicalProfileStore(), NewMemoryMeasurementStore(), defaultDatabasePath)
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

	got := outputForArgs([]string{"vitalynq", "observations", "add", "--date"}, store, NewMemoryMedicalProfileStore(), NewMemoryMeasurementStore(), defaultDatabasePath)
	want := "Date ou texte d'observation manquant."

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestOutputForArgsObservationsAddWithDateMissingText(t *testing.T) {
	store := NewMemoryObservationStore()

	got := outputForArgs([]string{"vitalynq", "observations", "add", "--date", "2026-07-29"}, store, NewMemoryMedicalProfileStore(), NewMemoryMeasurementStore(), defaultDatabasePath)
	want := "Date ou texte d'observation manquant."

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestOutputForArgsObservationsAddWithInvalidDate(t *testing.T) {
	store := NewMemoryObservationStore()

	got := outputForArgs([]string{"vitalynq", "observations", "add", "--date", "29-07-2026", "Observation fictive de test"}, store, NewMemoryMedicalProfileStore(), NewMemoryMeasurementStore(), defaultDatabasePath)
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

	got := outputForArgs([]string{"vitalynq", "profile", "show"}, NewMemoryObservationStore(), profileStore, NewMemoryMeasurementStore(), defaultDatabasePath)
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
		defaultDatabasePath,
	)
	want := "Profil médical #1 enregistré."

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestOutputForArgsProfileSetMissingLabel(t *testing.T) {
	got := outputForArgs([]string{"vitalynq", "profile", "set"}, NewMemoryObservationStore(), NewMemoryMedicalProfileStore(), NewMemoryMeasurementStore(), defaultDatabasePath)
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
