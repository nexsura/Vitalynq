package main

import (
	"fmt"
	"os"
)

func main() {
	store := NewMemoryObservationStore()
	fmt.Println(outputForArgs(os.Args, store))
}
