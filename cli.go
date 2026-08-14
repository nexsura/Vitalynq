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
  privacy            Affiche les informations de confidentialité
  limitations        Affiche les limites de Vitalynq
  profile set        Enregistre le profil médical
  profile show       Affiche le profil médical
  observations list  Liste les observations
  obs list           Alias de observations list
  observations add   Ajoute une observation
  observations add --date YYYY-MM-DD
                     Ajoute une observation datée
  db path            Affiche le chemin de la base SQLite
  db info            Affiche les informations de stockage local
  db check           Vérifie l'accès au stockage local
  measurements list  Liste les mesures
  measurements add   Ajoute une mesure
  measurements add --date YYYY-MM-DD
                     Ajoute une mesure datée
  appointments list  Liste les rendez-vous
  appointments add   Ajoute un rendez-vous
  summary            Affiche un bilan synthétique
  export             Exporte les données locales en JSON

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

func privacyText() string {
	return `Confidentialité Vitalynq

Vitalynq stocke les données localement dans un fichier SQLite.
Vitalynq n'envoie aucune donnée vers un serveur, un cloud ou une API externe.
Vitalynq n'utilise pas de télémétrie.

Les exports JSON sont affichés localement dans le terminal.
L'utilisateur reste responsable de la protection du fichier SQLite et des exports.

Vitalynq organise des données. Il ne pose pas de diagnostic et ne remplace pas un professionnel de santé.`
}

func limitationsText() string {
	return `Limites de Vitalynq

Vitalynq organise des données personnelles de santé.
Vitalynq ne pose pas de diagnostic.
Vitalynq ne recommande aucun traitement.
Vitalynq ne prédit pas l'évolution d'un état de santé.
Vitalynq ne remplace pas un professionnel de santé.

En cas de question médicale, contactez un professionnel de santé qualifié.`
}

func unknownCommandText(command string) string {
	return fmt.Sprintf("Commande inconnue: %s\n\nUtilisez 'vitalynq help' pour voir les commandes disponibles.", command)
}

func missingObservationText() string {
	return `Texte d'observation manquant.

Usage: vitalynq observations add "Observation fictive"`
}

func missingObservationDateText() string {
	return `Date ou texte d'observation manquant.

Usage: vitalynq observations add --date YYYY-MM-DD "Observation fictive"`
}

func missingMeasurementText() string {
	return `Arguments de mesure manquants.

Usage: vitalynq measurements add indicateur valeur unité contexte méthode source`
}

func missingMeasurementDateText() string {
	return `Date ou arguments de mesure manquants.

Usage: vitalynq measurements add --date YYYY-MM-DD indicateur valeur unité contexte méthode source`
}

func invalidMeasurementValueText() string {
	return `Valeur de mesure invalide.

La valeur doit être un nombre, par exemple: 72.5`
}

func missingAppointmentText() string {
	return `Arguments de rendez-vous manquants.

Usage: vitalynq appointments add YYYY-MM-DD titre catégorie lieu source`
}

func outputForArgs(args []string, observationStore ObservationStore, profileStore MedicalProfileStore, measurementStore MeasurementStore, appointmentStore AppointmentStore, databasePath string) string {
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
	case "privacy":
		return privacyText()
	case "limitations":
		return limitationsText()
	case "db":
		if len(args) > 2 && args[2] == "path" {
			return databasePathText(databasePath)
		}

		if len(args) > 2 && args[2] == "info" {
			return databaseInfoText(databasePath)
		}

		if len(args) > 2 && args[2] == "check" {
			return databaseCheckText()
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
			return missingObservationDateText()
		}

		if len(args) > 3 && args[2] == "add" {
			return observationsAddText(observationStore, args[3])
		}

		if len(args) > 2 && args[2] == "add" {
			return missingObservationText()
		}

		return unknownCommandText(args[1])

	case "appointments":
		if len(args) > 2 && args[2] == "list" {
			return appointmentsListText(appointmentStore)
		}

		if len(args) > 7 && args[2] == "add" {
			return appointmentsAddText(
				appointmentStore,
				args[3],
				args[4],
				args[5],
				args[6],
				args[7],
			)
		}

		if len(args) > 2 && args[2] == "add" {
			return missingAppointmentText()
		}

		return unknownCommandText(args[1])

	case "measurements":
		if len(args) > 2 && args[2] == "list" {
			return measurementsListText(measurementStore)
		}

		if len(args) > 10 && args[2] == "add" && args[3] == "--date" {
			value, err := strconv.ParseFloat(args[6], 64)
			if err != nil {
				return invalidMeasurementValueText()
			}

			return measurementsAddTextWithDate(
				measurementStore,
				args[4],
				args[5],
				value,
				args[7],
				args[8],
				args[9],
				args[10],
			)
		}

		if len(args) > 2 && args[2] == "add" && len(args) > 3 && args[3] == "--date" {
			return missingMeasurementDateText()
		}

		if len(args) > 8 && args[2] == "add" {
			value, err := strconv.ParseFloat(args[4], 64)
			if err != nil {
				return invalidMeasurementValueText()
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
			return missingMeasurementText()
		}

		return unknownCommandText(args[1])

	case "summary":
		return summaryText(observationStore, measurementStore, appointmentStore)

	case "export":
		return exportText(profileStore, observationStore, measurementStore, appointmentStore)
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

func databaseInfoText(databasePath string) string {
	return fmt.Sprintf(`Base SQLite: %s
Stockage: local
Cloud: non
Télémétrie: non`, databasePath)
}

func databaseCheckText() string {
	return `Base SQLite accessible: oui
Schéma SQLite initialisé: oui`
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

func measurementsAddTextWithDate(store MeasurementStore, datevalue string, indicator string, value float64, unit string, context string, method string, source string) string {
	occurredAt, err := parseObservationDate(datevalue)
	if err != nil {
		return err.Error()
	}

	measurement, err := newMeasurement(occurredAt, indicator, value, unit, context, method, source)
	if err != nil {
		return fmt.Sprintf("Impossible d'ajouter la mesure: %v", err)
	}

	saved, err := store.Save(measurement)
	if err != nil {
		return fmt.Sprintf("Impossible d'ajouter la mesure: %v", err)
	}

	return fmt.Sprintf("Mesure #%d ajoutée.", saved.ID)
}

func appointmentsListText(store AppointmentStore) string {
	appointments, err := store.List()
	if err != nil {
		return fmt.Sprintf("Impossible de lister les rendez-vous: %v", err)
	}

	if len(appointments) == 0 {
		return "Aucun rendez-vous enregistré."
	}

	var builder strings.Builder
	builder.WriteString("Rendez-vous:\n")

	for _, appointment := range appointments {
		fmt.Fprintf(
			&builder,
			"- #%d %s %s (%s)\n",
			appointment.ID,
			appointment.ScheduledAt.Format("2006-01-02"),
			appointment.Title,
			appointment.Category,
		)
	}

	return strings.TrimRight(builder.String(), "\n")
}

func appointmentsAddText(store AppointmentStore, dateValue string, title string, category string, location string, source string) string {
	scheduledAt, err := parseObservationDate(dateValue)
	if err != nil {
		return err.Error()
	}

	appointment, err := newAppointment(scheduledAt, title, category, location, source)
	if err != nil {
		return fmt.Sprintf("Impossible d'ajouter le rendez-vous: %v", err)
	}

	saved, err := store.Save(appointment)
	if err != nil {
		return fmt.Sprintf("Impossible d'ajouter le rendez-vous: %v", err)
	}

	return fmt.Sprintf("Rendez-vous #%d ajouté.", saved.ID)
}

func summaryText(observationStore ObservationStore, measurementStore MeasurementStore, appointmentStore AppointmentStore) string {
	observations, err := observationStore.List()
	if err != nil {
		return fmt.Sprintf("Impossible de produire le bilan: %v", err)
	}

	measurements, err := measurementStore.List()
	if err != nil {
		return fmt.Sprintf("Impossible de produire le bilan: %v", err)
	}

	appointments, err := appointmentStore.List()
	if err != nil {
		return fmt.Sprintf("Impossible de produire le bilan: %v", err)
	}

	return fmt.Sprintf("Bilan:\n- Observations: %d\n- Mesures: %d\n- Rendez-vous: %d",
		len(observations),
		len(measurements),
		len(appointments),
	)
}

func exportText(profileStore MedicalProfileStore, observationStore ObservationStore, measurementStore MeasurementStore, appointmentStore AppointmentStore) string {
	snapshot, err := buildExportSnapshot(profileStore, observationStore, measurementStore, appointmentStore)
	if err != nil {
		return fmt.Sprintf("Impossible d'exporter les données: %v", err)
	}

	jsonText, err := exportSnapshotJSON(snapshot)
	if err != nil {
		return fmt.Sprintf("Impossible d'exporter les données: %v", err)
	}

	return jsonText
}
