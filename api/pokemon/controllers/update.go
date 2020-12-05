package controllers

import (
	"net/http"
	"strconv"
	"workshop-go-gin-basic/api/pokemon/models"
	"workshop-go-gin-basic/database"

	"github.com/gin-gonic/gin"
)

// Update update pokemon controller
func Update(c *gin.Context) {
	pokemonID, err := strconv.ParseInt(c.Param("pokemonId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pokemon := models.Pokemon{}
	if err := c.Bind(&pokemon); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pokemon.ID = pokemonID
	conn := database.GetDatabase(c)
	if err := pokemon.UpdateMe(conn); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSONP(http.StatusNoContent, pokemon)
}
