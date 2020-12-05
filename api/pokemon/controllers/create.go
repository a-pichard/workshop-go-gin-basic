package controllers

import (
	"net/http"
	"workshop-go-gin-basic/api/pokemon/models"
	"workshop-go-gin-basic/database"

	"github.com/gin-gonic/gin"
	"gopkg.in/validator.v2"
)

// Create create pokemon controller
func Create(c *gin.Context) {
	pokemonBody := models.Pokemon{}
	if err := c.ShouldBindJSON(&pokemonBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := validator.Validate(pokemonBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	conn := database.GetDatabase(c)
	if err := pokemonBody.InsertMe(conn); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSONP(http.StatusOK, pokemonBody)
}
