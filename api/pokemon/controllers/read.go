package controllers

import (
	"net/http"
	"strconv"
	"workshop-go-gin-basic/api/pokemon/models"
	"workshop-go-gin-basic/database"

	"github.com/gin-gonic/gin"
)

// Get get pokemon controller
func Get(c *gin.Context) {
	pokemonID, err := strconv.ParseInt(c.Param("pokemonId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	conn := database.GetDatabase(c)
	pokemon := models.Pokemon{ID: pokemonID}
	if err := pokemon.GetByPk(conn); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSONP(http.StatusOK, pokemon)
}
