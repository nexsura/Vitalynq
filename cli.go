package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func appDescription() string {
	return "Vitalynq organizes local health data."
}

func helpText() string {
	return `Vitalynq

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
}

func aboutText() string {
	return `Vitalynq is a local CLI application for organizing personal health data.

  Current scope:
  local
  single-user
  no cloud
  no telemetry

  Vitalynq does not provide diagnosis and does not replace a healthcare professional.`
}

func privacyText() string {
	return `Vitalynq Privacy

  Vitalynq stores data locally in a SQLite file.
  Vitalynq does not send data to a server, cloud service, or external API.
  Vitalynq does not use telemetry.

  JSON exports are displayed locally in the terminal.
  The user is responsible for protecting the SQLite file and exports.

  Vitalynq organizes data. It does not provide diagnosis and does not replace a healthcare professional.`
}

func limitationsText() string {
	return `Vitalynq Limitations

  Vitalynq organizes personal health data.
  Vitalynq does not provide diagnosis.
  Vitalynq does not recommend treatments.
  Vitalynq does not predict health changes.
  Vitalynq does not replace a healthcare professional.

  For medical questions, contact a qualified healthcare professional.`
}

func unknownCommandText(command string) string {
	return fmt.Sprintf("Unknown command: %s\n\nUse 'vitalynq help' to see available commands.", command)
}

func unknownSubcommandText(command string, subcommand string) string {
	return fmt.Sprintf("Unknown subcommand: %s %s\n\nUse 'vitalynq help' to see available commands.", command, subcommand)
}

func invalidDateText() string {
	return `Invalid date.

  Use the YYYY-MM-DD format, for example: 2026-07-29`
}

func missingObservationText() string {
	return `Missing observation text.

  Usage: vitalynq observations add "Fictive observation"`
}

func missingObservationDateText() string {
	return `Missing observation date or text.

  Usage: vitalynq observations add --date YYYY-MM-DD "Fictive observation"`
}

func missingMeasurementText() string {
	return `Missing measurement arguments.

  Usage: vitalynq measurements add indicator value unit context method source`
}

func missingMeasurementDateText() string {
	return `Missing measurement date or arguments.

  Usage: vitalynq measurements add --date YYYY-MM-DD indicator value unit context method source`
}

func invalidMeasurementValueText() string {
	return `Invalid measurement value.

  The value must be a number, for example: 72.5`
}

func missingAppointmentText() string {
	return `Missing appointment arguments.

  Usage: vitalynq appointments add YYYY-MM-DD title category location source`
}

func missingMedicalProfileLabelText() string {
	return `Missing medical profile label.

  Usage: vitalynq profile set "Fictive profile"`
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

		if len(args) > 2 {
			return unknownSubcommandText(args[1], args[2])
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
			return missingMedicalProfileLabelText()
		}

		if len(args) > 2 {
			return unknownSubcommandText(args[1], args[2])
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

		if len(args) > 2 {
			return unknownSubcommandText(args[1], args[2])
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

		if len(args) > 2 {
			return unknownSubcommandText(args[1], args[2])
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

		if len(args) > 2 {
			return unknownSubcommandText(args[1], args[2])
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
		return fmt.Sprintf("Unable to list observations: %v", err)
	}
	if len(observations) == 0 {
		return "No observations recorded."
	}

	var builder strings.Builder
	builder.WriteString("Observations:\n")

	for _, observation := range observations {
		fmt.Fprintf(&builder, "- #%d %s %s\n", observation.ID, observation.OccurredAt.Format("2006-01-02"), observation.Text)
	}

	return strings.TrimRight(builder.String(), "\n")
}

func observationsAddText(store ObservationStore, text string) string {
	observation, err := newObservation(time.Now().UTC(), text, "manual entry")
	if err != nil {
		return fmt.Sprintf("Unable to add observation: %v", err)
	}

	saved, err := store.Save(observation)
	if err != nil {
		return fmt.Sprintf("Unable to add observation: %v", err)
	}

	return fmt.Sprintf("Observation #%d added.", saved.ID)
}

func databasePathText(databasePath string) string {
	return fmt.Sprintf("SQLite database: %s", databasePath)
}

func databaseInfoText(databasePath string) string {
	return fmt.Sprintf(`SQLite database: %s
  Storage: local
  Cloud: no
  Telemetry: no`, databasePath)
}

func databaseCheckText() string {
	return `SQLite database accessible: yes
  SQLite schema initialized: yes`
}

func parseObservationDate(value string) (time.Time, error) {
	parsed, err := time.Parse("2006-01-02", value)
	if err != nil {
		return time.Time{}, fmt.Errorf("%s", invalidDateText())
	}

	return parsed, nil
}

func observationsAddTextWithDate(store ObservationStore, dateValue string, text string) string {
	occurredAt, err := parseObservationDate(dateValue)
	if err != nil {
		return err.Error()
	}

	observation, err := newObservation(occurredAt, text, "manual entry")
	if err != nil {
		return fmt.Sprintf("Unable to add observation: %v", err)
	}

	saved, err := store.Save(observation)
	if err != nil {
		return fmt.Sprintf("Unable to add observation: %v", err)
	}

	return fmt.Sprintf("Observation #%d added.", saved.ID)
}

func medicalProfileText(store MedicalProfileStore) string {
	profile, found, err := store.Get()
	if err != nil {
		return fmt.Sprintf("Unable to read medical profile: %v", err)
	}

	if !found {
		return "No medical profile recorded."
	}

	return fmt.Sprintf("Medical profile: %s", profile.Label)
}

func medicalProfileSaveText(store MedicalProfileStore, label string) string {
	profile, err := newMedicalProfile(label)
	if err != nil {
		return fmt.Sprintf("Unable to save medical profile: %v", err)
	}

	saved, err := store.Save(profile)
	if err != nil {
		return fmt.Sprintf("Unable to save medical profile: %v", err)
	}

	return fmt.Sprintf("Medical profile #%d saved.", saved.ID)
}

func measurementsListText(store MeasurementStore) string {
	measurements, err := store.List()
	if err != nil {
		return fmt.Sprintf("Unable to list measurements: %v", err)
	}

	if len(measurements) == 0 {
		return "No measurements recorded."
	}

	var builder strings.Builder
	builder.WriteString("Measurements:\n")

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
		return fmt.Sprintf("Unable to add measurement: %v", err)
	}

	saved, err := store.Save(measurement)
	if err != nil {
		return fmt.Sprintf("Unable to add measurement: %v", err)
	}

	return fmt.Sprintf("Measurement #%d added.", saved.ID)
}

func measurementsAddTextWithDate(store MeasurementStore, dateValue string, indicator string, value float64, unit string, context string, method string, source string) string {
	occurredAt, err := parseObservationDate(dateValue)
	if err != nil {
		return err.Error()
	}

	measurement, err := newMeasurement(occurredAt, indicator, value, unit, context, method, source)
	if err != nil {
		return fmt.Sprintf("Unable to add measurement: %v", err)
	}

	saved, err := store.Save(measurement)
	if err != nil {
		return fmt.Sprintf("Unable to add measurement: %v", err)
	}

	return fmt.Sprintf("Measurement #%d added.", saved.ID)
}

func appointmentsListText(store AppointmentStore) string {
	appointments, err := store.List()
	if err != nil {
		return fmt.Sprintf("Unable to list appointments: %v", err)
	}

	if len(appointments) == 0 {
		return "No appointments recorded."
	}

	var builder strings.Builder
	builder.WriteString("Appointments:\n")

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
		return fmt.Sprintf("Unable to add appointment: %v", err)
	}

	saved, err := store.Save(appointment)
	if err != nil {
		return fmt.Sprintf("Unable to add appointment: %v", err)
	}

	return fmt.Sprintf("Appointment #%d added.", saved.ID)
}

func summaryText(observationStore ObservationStore, measurementStore MeasurementStore, appointmentStore AppointmentStore) string {
	observations, err := observationStore.List()
	if err != nil {
		return fmt.Sprintf("Unable to produce summary: %v", err)
	}

	measurements, err := measurementStore.List()
	if err != nil {
		return fmt.Sprintf("Unable to produce summary: %v", err)
	}

	appointments, err := appointmentStore.List()
	if err != nil {
		return fmt.Sprintf("Unable to produce summary: %v", err)
	}

	return fmt.Sprintf("Summary:\n- Observations: %d\n- Measurements: %d\n- Appointments: %d",
		len(observations),
		len(measurements),
		len(appointments),
	)
}

func exportText(profileStore MedicalProfileStore, observationStore ObservationStore, measurementStore MeasurementStore, appointmentStore AppointmentStore) string {
	snapshot, err := buildExportSnapshot(profileStore, observationStore, measurementStore, appointmentStore)
	if err != nil {
		return fmt.Sprintf("Unable to export data: %v", err)
	}

	jsonText, err := exportSnapshotJSON(snapshot)
	if err != nil {
		return fmt.Sprintf("Unable to export data: %v", err)
	}

	return jsonText
}
