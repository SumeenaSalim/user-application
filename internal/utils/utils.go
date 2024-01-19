package utils

import (
	"database/sql"
	"fmt"
)

// ConnectDB establishes a connection to a PostgreSQL database
func ConnectDB(username, password, host, dbName string, port int) (*sql.DB, error) {
	// Construct the PostgreSQL connection string with port
	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable", username, password, host, port, dbName)

	// Open a database connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %v", err)

	}

	// db.SetMaxIdleConns(5)                  // Set the maximum number of idle connections
	// db.SetConnMaxLifetime(5 * time.Minute) // Set the maximum amount of time a connection may be reused

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	return db, nil
}
