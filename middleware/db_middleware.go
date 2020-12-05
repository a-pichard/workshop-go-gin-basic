package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

func GetDatabase(c *gin.Context) *pg.DB {
	db, _ := c.Get("db")
	return db.(*pg.DB)
}

func Database(conn *pg.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", conn)
		c.Next()
	}
}
