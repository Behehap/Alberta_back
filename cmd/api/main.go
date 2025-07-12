// cmd/api/main.go
package main

import (
	"context"
	"database/sql"
	"flag"
	"log"
	"os"
	"time"

	// Make sure this module path matches your go.mod file

	"github.com/Behehap/Alberta/internal/store"
	_ "github.com/lib/pq" // PostgreSQL driver
)

const version = "1.0.0"

func main() {
	var cfg config

	// Read configuration settings from command-line flags.
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")

	// Use the ALBERTA_DB_DSN environment variable for the database connection string.
	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("ALBERTA_DB_DSN"), "PostgreSQL DSN")

	// Read the database connection pool settings from flags.
	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.StringVar(&cfg.db.maxIdleTime, "db-max-idle-time", "15m", "PostgreSQL max connection idle time")

	flag.Parse()

	// Initialize a new logger.
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// Connect to the database.
	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()
	logger.Println("database connection pool established")

	// Initialize the data store.
	storage := store.NewStorage(db)

	// Create the application instance.
	app := &application{
		config: cfg,
		logger: logger,
		store:  storage,
	}

	// Run the server.
	err = app.run()
	if err != nil {
		logger.Fatal(err)
	}
}

// openDB creates and returns a new sql.DB object, which represents a connection pool.
func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(cfg.db.maxOpenConns)
	db.SetMaxIdleConns(cfg.db.maxIdleConns)
	duration, err := time.ParseDuration(cfg.db.maxIdleTime)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}
