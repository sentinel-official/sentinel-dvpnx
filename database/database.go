package database

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/sentinel-official/dvpn-node/database/models"
)

// New initializes a new database connection with the specified file path and configuration.
// It also performs migrations to ensure the database schema is up to date with the models.
func New(filepath string, cfg *gorm.Config) (*gorm.DB, error) {
	// Build the SQLite DSN
	dsn := fmt.Sprintf("%s?_busy_timeout=5000&_journal_mode=WAL", filepath)

	// Open a database connection using the provided filepath and configuration.
	db, err := gorm.Open(sqlite.Open(dsn), cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Set connection pooling for SQLite
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve database instance: %w", err)
	}

	sqlDB.SetMaxIdleConns(1) // Keep one idle connection for reuse
	sqlDB.SetMaxOpenConns(1) // Restrict to one open connection

	// Run migrations to apply the schema of the `Session` model to the database.
	if err := db.AutoMigrate(&models.Session{}); err != nil {
		return nil, fmt.Errorf("failed to auto migrate session: %w", err)
	}

	// Return the database connection if everything is successful.
	return db, nil
}

// NewDefault uses default configuration settings and calls the New function to initialize the database.
func NewDefault(filepath string) (*gorm.DB, error) {
	// Define default GORM configuration settings.
	cfg := gorm.Config{
		Logger:         logger.Discard,
		PrepareStmt:    false,
		TranslateError: true,
	}

	// Call New with the default configuration.
	return New(filepath, &cfg)
}
