package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/sentinel-official/dvpn-node/database/models"
)

// New initializes a new database connection with the specified file path and configuration.
// It also performs migrations to ensure the database schema is up to date with the models.
func New(filepath string, cfg *gorm.Config) (*gorm.DB, error) {
	// Open a database connection using the provided filepath and configuration.
	db, err := gorm.Open(sqlite.Open(filepath), cfg)
	if err != nil {
		return nil, err
	}

	// Run migrations to apply the schema of the `Session` model to the database.
	if err := db.AutoMigrate(&models.Session{}); err != nil {
		return nil, err
	}

	// Return the database connection if everything is successful.
	return db, nil
}

// NewDefault uses default configuration settings and calls the New function to initialize the database.
func NewDefault(filepath string) (*gorm.DB, error) {
	// Define default GORM configuration settings.
	cfg := gorm.Config{
		Logger:      logger.Discard,
		PrepareStmt: false,
	}

	// Call New with the default configuration.
	return New(filepath, &cfg)
}
