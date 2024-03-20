package database

import (
	"context"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type DbConn struct {
	Dbpool *pgxpool.Pool
}

// initializes the database connection pool and runs migrations
func InitPool(config *DbConfig) *DbConn {
	// set up the connection string to the db
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.dbHost, config.dbPort, config.dbUser, config.dbPassword, config.dbName)

	// initialize the connection pool
	dbpool, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		log.Panicf("error opening database: %v", err)
	}
	defer dbpool.Close()

	// verify the connection
	if err := dbpool.Ping(context.Background()); err != nil {
		log.Panicf("error pinging database: %v", err)
	}

	// run migrations if needed only if migration path is set
	if config.dbMigrationPath != nil {
		if err := runMigration(connectionString, *config.dbMigrationPath); err != nil {
			log.Panicf("error pinging database: %v", err)
		}
	}

	// return the connection pool pointer singleton connection
	return &DbConn{
		Dbpool: dbpool,
	}
}

// closes the database connection pool
func (db *DbConn) ClosePool() {
	if db.Dbpool != nil {
		db.Dbpool.Close()
	}
}

// run database migration files only if migration path is set
func runMigration(connString string, path string) error {
	m, err := migrate.New(
		fmt.Sprintf("file://%s", path),
		connString,
	)
	if err != nil {
		return err
	}

	// Migrate up to the latest active version
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
