package main

import (
	mydb "bandscape-backend/pkg/config"
	routes "bandscape-backend/pkg/routes"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	mydb.DbConnection()
	router := gin.Default()
	router.Use(cors.Default())

	routes.SetupRoutes(router)

	router.Run("localhost:3001")
	fmt.Println("Server running on port 3001")
}
