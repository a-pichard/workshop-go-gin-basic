package database

import (
	"os"

	"github.com/go-pg/pg/v10"
)

// ConnectDatabase create database connection
func ConnectDatabase() *pg.DB {
	db := pg.Connect(&pg.Options{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Addr:     os.Getenv("DB_HOST"),
		Database: os.Getenv("DB_NAME"),
	})
	if err := db.Ping(db.Context()); err != nil {
		panic(err)
	}
	return db
}
