package main

import (
	"fmt"
	"os"
)

func main() {
	db, err := openSQLite("vitalynq.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	if err := initializeSQLiteSchema(db); err != nil {
		fmt.Println(err)
		return
	}

	store := NewSQLiteObservationStore(db)
	fmt.Println(outputForArgs(os.Args, store))
}
