package main

import (
	"fmt"
	"os"
)

func appDescription() string {
	return "Vitalynq organise des données de santé locales."
}

func helpText() string {
	return `Vitalynq

	Commandes:
	help	Affiche cette aide
	version	Affiche la version

	Vitalynq organise des données. Il ne pose pas de diagnostic.`
}

func outputForArgs(args []string) string {
	if len(args) <= 1 {
		return appDescription()
	}

	switch args[1] {
	case "help":
		return helpText()
	case "version":
		return "Vitalynq 0.1.0-dev"
	default:
		return appDescription()
	}
}

func main() {
	fmt.Println(outputForArgs(os.Args))
}
