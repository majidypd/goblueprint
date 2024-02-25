package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/majidypd/goblueprint/config"
)

// DB represents a database connection.
type DB struct {
	*gorm.DB
}

// NewDB creates a new database connection.
func NewDB(cfg *config.Config) (*DB, error) {
	db, err := gorm.Open(postgres.Open(generateDSN(cfg)), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

// generateDSN generates a PostgreSQL DSN string from the struct fields.
func generateDSN(cfg *config.Config) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		cfg.PostgresHost, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDbName, cfg.PostgresPort)
}