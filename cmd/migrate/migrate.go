package main

import (
	"database/sql"
	"path/filepath"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

const (
	migrationsPath  = "db"
	migrationsTable = "migrations"
	sqliteDatabase  = "matchers.db"
)

func main() {
	logger := logrus.StandardLogger()

	db, err := sql.Open("sqlite3", filepath.Join(migrationsPath, sqliteDatabase))
	if err != nil {
		logger.WithError(err).Fatal("cannot open db")
	}

	defer func() {
		if err := db.Close(); err != nil {
			logger.WithError(err).Fatal("cannot close db")
		}
	}()

	config := &sqlite3.Config{
		MigrationsTable: migrationsTable,
	}

	driver, err := sqlite3.WithInstance(db, config)
	if err != nil {
		logger.WithError(err).Fatal("cannot open sqlite3 driver")
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://./db/migrations",
		"sqlite3", driver)
	if err != nil {
		logger.WithError(err).Fatal("cannot open migrations")
	}

	m.Up()
}
