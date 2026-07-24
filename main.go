package main

import (
	"fmt"
	"os"
)

const defaultDatabasePath = "vitalynq.db"

func main() {
	databasePath, args := databasePathAndArgs(os.Args)

	db, err := openSQLite(databasePath)
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
	fmt.Println(outputForArgs(args, store))
}

func databasePathAndArgs(args []string) (string, []string) {
	if len(args) > 2 && args[1] == "--db" {
		cleanedArgs := append([]string{args[0]}, args[3:]...)
		return args[2], cleanedArgs
	}

	return defaultDatabasePath, args
}
