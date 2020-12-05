package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "salut",
	})
}
