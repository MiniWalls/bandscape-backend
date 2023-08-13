package main

import (
	mydb "bandscape-backend/pkg/config"
	routes "bandscape-backend/pkg/routes"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("APP_ENV") != "production" {
		godotenv.Load() //Load .env file and error check
	}

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	mydb.DbConnection()
	router := gin.Default()
	router.Use(cors.Default())
	router.Use(gin.Logger())
	routes.SetupRoutes(router)

	log.Println("Server running on port " + port)
	log.Println("###########################################################################################")
	log.Println("###########################################################################################")

	router.Run(":" + port)
}
