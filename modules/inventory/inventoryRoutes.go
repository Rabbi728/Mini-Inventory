package inventory

import "github.com/gin-gonic/gin"

func RegisterInventoryRoutes(rg *gin.RouterGroup, authMiddleware gin.HandlerFunc) {
	inventoryService := InventoryService{}
	inventoryCtrl := InventoryController{Service: inventoryService}

	inventoryGroup := rg.Group("/inventories")
	inventoryGroup.Use(authMiddleware)
	{
		inventoryGroup.POST("/receive", inventoryCtrl.Receive)
		inventoryGroup.POST("/delivery", inventoryCtrl.Delivery)
	}
}
