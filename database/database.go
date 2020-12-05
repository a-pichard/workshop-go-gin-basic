package database

import (
	"os"

	"github.com/gin-gonic/gin"
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

// GetDatabase get database connection from context
func GetDatabase(c *gin.Context) *pg.DB {
	db, _ := c.Get("db")
	return db.(*pg.DB)
}

// Database set database in gin context
func Database(conn *pg.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", conn)
		c.Next()
	}
}
