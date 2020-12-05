package controllers

import (
	"net/http"
	"workshop-go-gin-basic/api/pokemon/models"
	"workshop-go-gin-basic/database"

	"github.com/gin-gonic/gin"
)

// List list pokemon controller
func List(c *gin.Context) {
	var pokemons models.Pokemons
	conn := database.GetDatabase(c)
	if err := pokemons.List(conn); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if pokemons == nil {
		c.JSON(http.StatusOK, []gin.H{})
	} else {
		c.JSONP(http.StatusOK, pokemons)
	}
}
