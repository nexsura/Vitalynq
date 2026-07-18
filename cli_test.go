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
  help     Affiche cette aide
  version  Affiche la version
  about    Affiche le périmètre actuel

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
	got := outputForArgs([]string{"vitalynq"}, NewMemoryObservationStore())
	want := "Vitalynq organise des données de santé locales."

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestOutputForArgsHelp(t *testing.T) {
	got := outputForArgs([]string{"vitalynq", "help"}, NewMemoryObservationStore())
	want := helpText()

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestOutputForArgsVersion(t *testing.T) {
	got := outputForArgs([]string{"vitalynq", "version"}, NewMemoryObservationStore())
	want := "Vitalynq 0.1.0-dev"

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestOutputForArgsAbout(t *testing.T) {
	got := outputForArgs([]string{"vitalynq", "about"}, NewMemoryObservationStore())
	want := aboutText()

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}

func TestOutputForArgsUnknownCommand(t *testing.T) {
	got := outputForArgs([]string{"vitalynq", "profil"}, NewMemoryObservationStore())
	want := unknownCommandText("profil")

	if got != want {
		t.Fatalf("outputForArgs() = %q, want %q", got, want)
	}
}
