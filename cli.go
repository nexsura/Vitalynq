package main

import (
	"fmt"
	"strings"
	"time"
)

func appDescription() string {
	return "Vitalynq organise des données de santé locales."
}

func helpText() string {
	return `Vitalynq

Commandes:
  help     Affiche cette aide
  version  Affiche la version
  about    Affiche le périmètre actuel
  observations list  Liste les observations
  observations add   Ajoute une observation

Vitalynq organise des données. Il ne pose pas de diagnostic.`
}

func aboutText() string {
	return `Vitalynq est une application CLI locale pour organiser des données personnelles de santé.

Périmètre actuel:
	local
	mono-utilisateur
	sans cloud
	sans télémétrie

Vitalynq ne pose pas de diagnostic et ne remplace pas un professionnel de santé.`
}

func unknownCommandText(command string) string {
	return fmt.Sprintf("Commande inconnue: %s\n\nUtilisez 'vitalynq help' pour voir les commandes disponibles.", command)
}

func outputForArgs(args []string, store ObservationStore) string {
	if len(args) <= 1 {
		return appDescription()
	}

	switch args[1] {
	case "help":
		return helpText()
	case "version":
		return "Vitalynq 0.1.0-dev"
	case "about":
		return aboutText()
	case "observations":
		if len(args) > 2 && args[2] == "list" {
			return observationsListText(store)
		}

		if len(args) > 3 && args[2] == "add" {
			return observationsAddText(store, args[3])
		}

		return unknownCommandText(args[1])
	default:
		return unknownCommandText(args[1])
	}
}

func observationsListText(store ObservationStore) string {
	observations := store.List()
	if len(observations) == 0 {
		return "Aucune observation enregistrée."
	}

	var builder strings.Builder
	builder.WriteString("Observations:\n")

	for _, observation := range observations {
		fmt.Fprintf(&builder, "- #%d %s\n", observation.ID, observation.Text)
	}

	return strings.TrimRight(builder.String(), "\n")
}

func observationsAddText(store ObservationStore, text string) string {
	observation, err := newObservation(time.Now().UTC(), text, "saisie manuelle")
	if err != nil {
		return fmt.Sprintf("Impossible d'ajouter l'observation: %v", err)
	}

	saved, err := store.Save(observation)
	if err != nil {
		return fmt.Sprintf("Impossible d'ajouter l'observation: %v", err)
	}

	return fmt.Sprintf("Observation #%d ajoutée.", saved.ID)
}
