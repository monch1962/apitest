package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose"
)

// DBSetup initiate connection to a mysql database, running migrations
func DBSetup(dsn string, setup func(db *sqlx.DB)) *sqlx.DB {
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		panic(err)
	}

	if err := goose.SetDialect("mysql"); err != nil {
		panic(err)
	}
	errMigration := goose.Up(db.DB, "./migrations")
	if errMigration != nil {
		panic(errMigration)
	}

	setup(db)
	return db
}
