package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"workshop-go-gin-basic/api/pokemon/models"
	"workshop-go-gin-basic/database"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	pokemonID, err := strconv.ParseInt(c.Param("pokemonId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	conn := database.GetDatabase(c)
	pokemon := models.Pokemon{ID: pokemonID}
	if err := pokemon.DeleteByPk(conn); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Not Found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Pokemon deleted",
	})
}
