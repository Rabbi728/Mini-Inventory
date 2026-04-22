package inventory

import (
	"mini-inventory/modules/user"
	"mini-inventory/utils"

	"github.com/gin-gonic/gin"
)

type InventoryController struct {
	Service InventoryService
}

func (ctrl *InventoryController) Receive(c *gin.Context) {
	var inv Inventory
	if err := c.ShouldBindJSON(&inv); err != nil {
		c.JSON(400, utils.ErrorResponse("Validation failed", utils.FormatValidationErrors(err)))
		return
	}

	u, exists := c.Get("user")
	if !exists {
		c.JSON(401, utils.ErrorResponse("Unauthorized", "User not found in context"))
		return
	}
	currentUser := u.(user.User)
	inv.CreatedBy = currentUser.ID
	inv.RecordType = "IN"

	if err := ctrl.Service.CreateInventory(&inv); err != nil {
		c.JSON(500, utils.ErrorResponse("Failed to receive inventory", err.Error()))
		return
	}

	c.JSON(201, utils.SuccessResponse("Inventory received successfully", inv))
}

func (ctrl *InventoryController) Delivery(c *gin.Context) {
	var inv Inventory
	if err := c.ShouldBindJSON(&inv); err != nil {
		c.JSON(400, utils.ErrorResponse("Validation failed", utils.FormatValidationErrors(err)))
		return
	}

	u, exists := c.Get("user")
	if !exists {
		c.JSON(401, utils.ErrorResponse("Unauthorized", "User not found in context"))
		return
	}
	currentUser := u.(user.User)
	inv.CreatedBy = currentUser.ID
	inv.RecordType = "OUT"

	if err := ctrl.Service.CreateInventory(&inv); err != nil {
		c.JSON(500, utils.ErrorResponse("Failed to deliver inventory", err.Error()))
		return
	}

	c.JSON(201, utils.SuccessResponse("Inventory delivered successfully", inv))
}
