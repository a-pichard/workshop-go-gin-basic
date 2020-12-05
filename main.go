package main

import (
	"fmt"
	"workshop-go-gin-basic/database"
	"workshop-go-gin-basic/middleware"

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
	router.Use(middleware.Database(conn))

	router.Run(":8080")
}
