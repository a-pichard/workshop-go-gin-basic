package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

func Database(conn *pg.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", conn)
		c.Next()
	}
}
