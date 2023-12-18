package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tanapon395/sa-66-example/controller"
	"github.com/tanapon395/sa-66-example/entity"
)

func main() {
	entity.SetupDatabase()
	r := gin.Default()
	r.Use(CORSMiddleware())
	// User Routes
	r.GET("/users", controller.ListUsers)
	r.GET("/user/:id", controller.GetUser)
	r.POST("/users", controller.CreateUser)
	r.PATCH("/users", controller.UpdateUser)
	r.DELETE("/users/:id", controller.DeleteUser)
	// Gender Routes
	r.GET("/genders", controller.ListGenders)
	// Run the server
	r.Run("localhost:8000")
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
