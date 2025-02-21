package config

import (
	"os"

	"github.com/MH0R/Gopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")

	// Check if the database exists
	dbPath := "./db/main.db"
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("database file not found, creating... ")
		// Crate database file and directory
		err := os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()

	}

	// Create a DB and Connect

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errf("Error initializing db sqlite: %v", err)
		return nil, err
	}
	// Migrate schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errf("sqlite migration error: %v", err)
	}

	// Return DB
	return db, nil
}
