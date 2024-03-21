package database

import (
	"context"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type DbConn struct {
	Dbpool *pgxpool.Pool
}

// initializes the database connection pool and runs migrations
func InitPool(config *DbConfig) *DbConn {
	// initialize the connection pool
	dbpool, err := pgxpool.New(context.Background(), config.dbUrl)
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
		if err := runMigration(config.dbUrl, *config.dbMigrationPath); err != nil {
			log.Panicf("error running migration: %v", err)
		}
	}

	log.Print("success database connection")

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

	// migrate up to the latest active version
	err = m.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			log.Println("no new migrations to apply")
			return nil
		}
		return err
	}

	log.Println("migrations were applied successfully")

	return nil
}
