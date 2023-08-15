package main

import (
	"backend/app/db"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	Init()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong!",
		})
	})
	r.Run(":8080")
}

func Init() {
	godotenv.Load()
	database := db.Database{}
	database.ConnectDatabase()
	database.Migrate()
}
