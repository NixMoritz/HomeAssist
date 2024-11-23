package database

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var migrationsDir = `internal\storage\SQL Migration`

func Migration(db *sql.DB) {
	// Read all files in the migrations directory
	err := filepath.Walk(migrationsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if the file is a regular file and ends with .sql
		if !info.IsDir() && filepath.Ext(path) == ".sql" {
			// Read the SQL file
			sqlFile, err := os.ReadFile(path)
			if err != nil {
				log.Printf("Error reading SQL file %s: %v", path, err)
				return err
			}

			// Split SQL commands if there are multiple statements
			commands := strings.Split(string(sqlFile), ";")
			for _, command := range commands {
				// Trim spaces and check if the command is not empty
				if trimmedCommand := strings.TrimSpace(command); trimmedCommand != "" {
					//Execute the SQL statement
					_, err := db.Exec(trimmedCommand)
					if err != nil {
						log.Printf("Error executing SQL command from file %s: %v", path, err)
					} else {
						// Log success
						log.Printf("Successfully executed command from %s: %s", path, trimmedCommand)
					}
				}
			}
		}
		return nil
	})

	if err != nil {
		log.Fatalf("Error walking through migrations directory: %v", err)
	}
}
