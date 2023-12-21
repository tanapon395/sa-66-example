package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tanapon395/sa-66-example/controller"
	"github.com/tanapon395/sa-66-example/entity"
)

func main() {
	entity.ConnectDB()
	r := gin.Default()
	r.Use(CORSMiddleware())
	// Auth Routes
	r.POST("/login", controller.Login)
	router := r.Group("")
	{
		// router.Use(middlewares.Authorizes())
		// {
		// User Routes
		router.GET("/users", controller.ListUsers)
		router.GET("/user/:id", controller.GetUser)
		router.POST("/users", controller.CreateUser)
		router.PATCH("/users", controller.UpdateUser)
		router.DELETE("/users/:id", controller.DeleteUser)
		// Gender Routes
		router.GET("/genders", controller.ListGenders)
		// }
	}

	// Run the server
	r.Run()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
