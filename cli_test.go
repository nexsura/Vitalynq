package main

import "testing"

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
  observations list  Liste les observations
  obs list           Alias de observations list
  observations add   Ajoute une observation
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
	got := outputForArgs([]string{"vitalynq"}, NewMemoryObservationStore(), defaultDatabasePath)
	want := "Vitalynq organise des données de santé locales."

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestOutputForArgsHelp(t *testing.T) {
	got := outputForArgs([]string{"vitalynq", "help"}, NewMemoryObservationStore(), defaultDatabasePath)
	want := helpText()

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestOutputForArgsVersion(t *testing.T) {
	got := outputForArgs([]string{"vitalynq", "version"}, NewMemoryObservationStore(), defaultDatabasePath)
	want := "Vitalynq 0.1.0-dev"

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestOutputForArgsAbout(t *testing.T) {
	got := outputForArgs([]string{"vitalynq", "about"}, NewMemoryObservationStore(), defaultDatabasePath)
	want := aboutText()

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestOutputForArgsUnknownCommand(t *testing.T) {
	got := outputForArgs([]string{"vitalynq", "profil"}, NewMemoryObservationStore(), defaultDatabasePath)
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
	got := outputForArgs([]string{"vitalynq", "observations", "list"}, NewMemoryObservationStore(), defaultDatabasePath)
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
- #1 Observation fictive de test`

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

	got := outputForArgs([]string{"vitalynq", "observations", "add", "Observation fictive de test"}, store, defaultDatabasePath)
	want := "Observation #1 ajoutée."

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestOutputForArgsObservationsAddMissingText(t *testing.T) {
	store := NewMemoryObservationStore()

	got := outputForArgs([]string{"vitalynq", "observations", "add"}, store, defaultDatabasePath)
	want := "Texte d'observation manquant."

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestOutputForArgsObsList(t *testing.T) {
	got := outputForArgs([]string{"vitalynq", "obs", "list"}, NewMemoryObservationStore(), defaultDatabasePath)
	want := "Aucune observation enregistrée."

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestOutputForArgsObsAdd(t *testing.T) {
	store := NewMemoryObservationStore()

	got := outputForArgs([]string{"vitalynq", "obs", "add", "Observation fictive de test"}, store, defaultDatabasePath)
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
	got := outputForArgs([]string{"vitalynq", "db", "path"}, NewMemoryObservationStore(), "test.db")
	want := "Base SQLite: test.db"

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}
