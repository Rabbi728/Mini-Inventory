package main

import (
	"basic-inventory-app/config"
	"basic-inventory-app/modules/auth"
	"basic-inventory-app/modules/user"
	"basic-inventory-app/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	utils.RunMigrations()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Go Inventory API (powered by sqlx)",
		})
	})
	
	api := r.Group("/api")
	user.RegisterUserRoutes(api, auth.AuthMiddleware())
	auth.RegisterAuthRoutes(api)
	
	fmt.Println("Server starting on :8080...")
	if err := r.Run(":8080"); err != nil {
		fmt.Println("Error starting server:", err)
	}
}