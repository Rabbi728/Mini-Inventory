package user

import (
	"mini-inventory/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Service UserService
}

func (ctrl *UserController) GetUsers(c *gin.Context) {
	users, err := ctrl.Service.GetAllUsers()
	if err != nil {
		c.JSON(500, utils.ErrorResponse("Failed to fetch users", err.Error()))
		return
	}
	c.JSON(200, utils.SuccessResponse("Users retrieved successfully", users))
}

func (ctrl *UserController) CreateUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, utils.ErrorResponse("Validation failed", utils.FormatValidationErrors(err)))
		return
	}

	if err := ctrl.Service.CreateUser(&user); err != nil {
		c.JSON(500, utils.ErrorResponse("Failed to create user", err.Error()))
		return
	}

	c.JSON(201, utils.SuccessResponse("User created successfully", user))
}

func (ctrl *UserController) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := ctrl.Service.GetUserByID(id)
	if err != nil {
		c.JSON(404, utils.ErrorResponse("User not found", nil))
		return
	}
	c.JSON(200, utils.SuccessResponse("User retrieved successfully", user))
}

func (ctrl *UserController) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	user, err := ctrl.Service.GetUserByID(id)
	if err != nil {
		c.JSON(404, utils.ErrorResponse("User not found", nil))
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, utils.ErrorResponse("Validation failed", utils.FormatValidationErrors(err)))
		return
	}

	if err := ctrl.Service.UpdateUser(&user); err != nil {
		c.JSON(500, utils.ErrorResponse("Failed to update user", err.Error()))
		return
	}

	c.JSON(200, utils.SuccessResponse("User updated successfully", user))
}

func (ctrl *UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := ctrl.Service.DeleteUser(id); err != nil {
		c.JSON(500, utils.ErrorResponse("Failed to delete user", err.Error()))
		return
	}
	c.JSON(200, utils.SuccessResponse("User deleted successfully", nil))
}
