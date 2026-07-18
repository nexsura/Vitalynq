package main

import (
	"fmt"
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

	return "Observations enregistrées."
}
