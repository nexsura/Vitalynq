package main

import (
	"fmt"
	"os"
)

func appDescription() string {
	return "Vitalynq organise des données de santé locales."
}

func outputForArgs(args []string) string {
	if len(args) > 1 && args[1] == "version" {
		return "Vitalynq 0.1.0-dev"
	}

	return appDescription()
}

func main() {
	fmt.Println(outputForArgs(os.Args))
}
