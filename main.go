package main

import (
	"fmt"
	"mini-inventory/config"
	"mini-inventory/modules/auth"
	"mini-inventory/modules/inventory"
	"mini-inventory/modules/location"
	"mini-inventory/modules/product"
	"mini-inventory/modules/report"
	"mini-inventory/modules/user"
	"mini-inventory/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	utils.RunMigrations()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Go Inventory API",
		})
	})

	api := r.Group("/api")
	user.RegisterUserRoutes(api, auth.AuthMiddleware())
	product.RegisterProductRoutes(api, auth.AuthMiddleware())
	location.RegisterLocationRoutes(api, auth.AuthMiddleware())
	inventory.RegisterInventoryRoutes(api, auth.AuthMiddleware())
	report.RegisterReportRoutes(api, auth.AuthMiddleware())
	auth.RegisterAuthRoutes(api)

	fmt.Println("Server starting on :8080...")
	if err := r.Run(":8080"); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
