package utils

import (
	"basic-inventory-app/config"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
)

func RunMigrations() {
	migrationDir := "database/migrations"
	files, err := os.ReadDir(migrationDir)
	if err != nil {
		log.Fatal("Could not read migrations directory: ", err)
	}

	var sqlFiles []string
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".sql" {
			sqlFiles = append(sqlFiles, file.Name())
		}
	}

	sort.Strings(sqlFiles)

	for _, fileName := range sqlFiles {
		filePath := filepath.Join(migrationDir, fileName)
		content, err := os.ReadFile(filePath)
		if err != nil {
			log.Printf("Could not read migration file %s: %v", fileName, err)
			continue
		}

		fmt.Printf("Running migration: %s\n", fileName)
		_, err = config.DB.Exec(string(content))
		if err != nil {
			log.Fatalf("Failed to execute migration %s: %v", fileName, err)
		}
	}

	fmt.Println("All migrations completed successfully!")
}
