package location

import (
	"mini-inventory/modules/user"
	"mini-inventory/utils"

	"github.com/gin-gonic/gin"
)

type LocationController struct {
	Service LocationService
}

func (ctrl *LocationController) GetLocations(c *gin.Context) {
	locations, err := ctrl.Service.GetAllLocations()
	if err != nil {
		c.JSON(500, utils.ErrorResponse("Failed to fetch locations", err.Error()))
		return
	}
	c.JSON(200, utils.SuccessResponse("Locations retrieved successfully", locations))
}

func (ctrl *LocationController) CreateLocation(c *gin.Context) {
	var loc Location
	if err := c.ShouldBindJSON(&loc); err != nil {
		c.JSON(400, utils.ErrorResponse("Validation failed", utils.FormatValidationErrors(err)))
		return
	}

	u, exists := c.Get("user")
	if !exists {
		c.JSON(401, utils.ErrorResponse("Unauthorized", "User not found in context"))
		return
	}
	currentUser := u.(user.User)
	loc.CreatedBy = currentUser.ID

	if err := ctrl.Service.CreateLocation(&loc); err != nil {
		c.JSON(500, utils.ErrorResponse("Failed to create location", err.Error()))
		return
	}

	c.JSON(201, utils.SuccessResponse("Location created successfully", loc))
}

func (ctrl *LocationController) GetLocation(c *gin.Context) {
	id := c.Param("id")
	loc, err := ctrl.Service.GetLocationByID(id)
	if err != nil {
		c.JSON(404, utils.ErrorResponse("Location not found", nil))
		return
	}
	c.JSON(200, utils.SuccessResponse("Location retrieved successfully", loc))
}

func (ctrl *LocationController) UpdateLocation(c *gin.Context) {
	id := c.Param("id")
	loc, err := ctrl.Service.GetLocationByID(id)
	if err != nil {
		c.JSON(404, utils.ErrorResponse("Location not found", nil))
		return
	}

	if err := c.ShouldBindJSON(&loc); err != nil {
		c.JSON(400, utils.ErrorResponse("Validation failed", utils.FormatValidationErrors(err)))
		return
	}

	if err := ctrl.Service.UpdateLocation(&loc); err != nil {
		c.JSON(500, utils.ErrorResponse("Failed to update location", err.Error()))
		return
	}

	c.JSON(200, utils.SuccessResponse("Location updated successfully", loc))
}

func (ctrl *LocationController) DeleteLocation(c *gin.Context) {
	id := c.Param("id")
	if err := ctrl.Service.DeleteLocation(id); err != nil {
		c.JSON(500, utils.ErrorResponse("Failed to delete location", err.Error()))
		return
	}
	c.JSON(200, utils.SuccessResponse("Location deleted successfully", nil))
}
