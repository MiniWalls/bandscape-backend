package routes

import (
	"bandscape-backend/pkg/controllers" // Import your controller package

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	postRoutes := router.Group("/posts")
	{
		postRoutes.GET("", controllers.GetAllPosts)
		postRoutes.GET("/:id", controllers.GetPost)
		postRoutes.POST("", controllers.CreatePost)
		postRoutes.PUT("", controllers.UpdatePost)
		postRoutes.DELETE("/:id", controllers.DeletePost)
	}

	authRoutes := router.Group("/auth")
	{
		authRoutes.GET("", controllers.GetAuth)
		authRoutes.GET("/login", controllers.GetToken)
	}
}
