package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func appDescription() string {
	return "Vitalynq organise des données de santé locales."
}

func helpText() string {
	return `Vitalynq

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
  measurements list  Liste les mesures
  measurements add   Ajoute une mesure

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

func outputForArgs(args []string, observationStore ObservationStore, profileStore MedicalProfileStore, measurementStore MeasurementStore, databasePath string) string {
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
	case "db":
		if len(args) > 2 && args[2] == "path" {
			return databasePathText(databasePath)
		}

		return unknownCommandText(args[1])
	case "profile":
		if len(args) > 2 && args[2] == "show" {
			return medicalProfileText(profileStore)
		}

		if len(args) > 3 && args[2] == "set" {
			return medicalProfileSaveText(profileStore, args[3])
		}

		if len(args) > 2 && args[2] == "set" {
			return "Libellé du profil médical manquant."
		}

		return unknownCommandText(args[1])
	case "observations", "obs":
		if len(args) > 2 && args[2] == "list" {
			return observationsListText(observationStore)
		}

		if len(args) > 5 && args[2] == "add" && args[3] == "--date" {
			return observationsAddTextWithDate(observationStore, args[4], args[5])
		}

		if len(args) > 3 && args[2] == "add" && args[3] == "--date" {
			return "Date ou texte d'observation manquant."
		}

		if len(args) > 3 && args[2] == "add" {
			return observationsAddText(observationStore, args[3])
		}

		if len(args) > 2 && args[2] == "add" {
			return "Texte d'observation manquant."
		}

		return unknownCommandText(args[1])

	case "measurements":
		if len(args) > 2 && args[2] == "list" {
			return measurementsListText(measurementStore)
		}

		if len(args) > 8 && args[2] == "add" {
			value, err := strconv.ParseFloat(args[4], 64)
			if err != nil {
				return "Valeur de mesure invalide."
			}

			return measurementsAddText(
				measurementStore,
				args[3],
				value,
				args[5],
				args[6],
				args[7],
				args[8],
			)
		}

		if len(args) > 2 && args[2] == "add" {
			return "Arguments de mesure manquants."
		}

		return unknownCommandText(args[1])
	default:
		return unknownCommandText(args[1])
	}
}

func observationsListText(store ObservationStore) string {
	observations, err := store.List()
	if err != nil {
		return fmt.Sprintf("Impossible de lister les observations: %v", err)
	}
	if len(observations) == 0 {
		return "Aucune observation enregistrée."
	}

	var builder strings.Builder
	builder.WriteString("Observations:\n")

	for _, observation := range observations {
		fmt.Fprintf(&builder, "- #%d %s %s\n", observation.ID, observation.OccurredAt.Format("2006-01-02"), observation.Text)
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

func databasePathText(databasePath string) string {
	return fmt.Sprintf("Base SQLite: %s", databasePath)
}

func parseObservationDate(value string) (time.Time, error) {
	parsed, err := time.Parse("2006-01-02", value)
	if err != nil {
		return time.Time{}, fmt.Errorf("date invalide, utilisez YYYY-MM-DD")
	}

	return parsed, nil
}

func observationsAddTextWithDate(store ObservationStore, dateValue string, text string) string {
	occurredAt, err := parseObservationDate(dateValue)
	if err != nil {
		return err.Error()
	}

	observation, err := newObservation(occurredAt, text, "saisie manuelle")
	if err != nil {
		return fmt.Sprintf("Impossible d'ajouter l'observation: %v", err)
	}

	saved, err := store.Save(observation)
	if err != nil {
		return fmt.Sprintf("Impossible d'ajouter l'observation: %v", err)
	}

	return fmt.Sprintf("Observation #%d ajoutée.", saved.ID)
}

func medicalProfileText(store MedicalProfileStore) string {
	profile, found, err := store.Get()
	if err != nil {
		return fmt.Sprintf("Impossible de lire le profil médical: %v", err)
	}

	if !found {
		return "Aucun profil médical enregistré."
	}

	return fmt.Sprintf("Profil médical: %s", profile.Label)
}

func medicalProfileSaveText(store MedicalProfileStore, label string) string {
	profile, err := newMedicalProfile(label)
	if err != nil {
		return fmt.Sprintf("Impossible d'enregistrer le profil médical: %v", err)
	}

	saved, err := store.Save(profile)
	if err != nil {
		return fmt.Sprintf("Impossible d'enregistrer le profil médical: %v", err)
	}

	return fmt.Sprintf("Profil médical #%d enregistré.", saved.ID)
}

func measurementsListText(store MeasurementStore) string {
	measurements, err := store.List()
	if err != nil {
		return fmt.Sprintf("Impossible de lister les mesures: %v", err)
	}

	if len(measurements) == 0 {
		return "Aucune mesure enregistrée."
	}

	var builder strings.Builder
	builder.WriteString("Mesures:\n")

	for _, measurement := range measurements {
		fmt.Fprintf(
			&builder,
			"- #%d %s %s %.2f %s\n",
			measurement.ID,
			measurement.OccurredAt.Format("2006-01-02"),
			measurement.Indicator,
			measurement.Value,
			measurement.Unit,
		)
	}

	return strings.TrimRight(builder.String(), "\n")
}

func measurementsAddText(store MeasurementStore, indicator string, value float64, unit string, context string, method string, source string) string {
	measurement, err := newMeasurement(time.Now().UTC(), indicator, value, unit, context, method, source)
	if err != nil {
		return fmt.Sprintf("Impossible d'ajouter la mesure: %v", err)
	}

	saved, err := store.Save(measurement)
	if err != nil {
		return fmt.Sprintf("Impossible d'ajouter la mesure: %v", err)
	}

	return fmt.Sprintf("Mesure #%d ajoutée.", saved.ID)
}
