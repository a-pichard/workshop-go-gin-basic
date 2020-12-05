package main

import (
	"fmt"
	"workshop-go-gin-basic/api/pokemon"
	"workshop-go-gin-basic/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func loadDotEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func main() {
	loadDotEnv()

	conn := database.ConnectDatabase()
	defer conn.Close()

	router := gin.Default()
	router.Use(database.Database(conn))

	pokemon.Init(router, conn)

	router.Run(":8080")
}
