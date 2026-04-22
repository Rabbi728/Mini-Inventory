package auth

import (
	"mini-inventory/modules/user"
	"mini-inventory/utils"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	Service AuthService
}

func (ctrl *AuthController) Register(c *gin.Context) {
	var input user.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, utils.ErrorResponse("Validation failed", utils.FormatValidationErrors(err)))
		return
	}

	u, err := ctrl.Service.Register(input)
	if err != nil {
		c.JSON(500, utils.ErrorResponse("Failed to create user", err.Error()))
		return
	}

	c.JSON(201, utils.SuccessResponse("User created successfully", u))
}

func (ctrl *AuthController) Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, utils.ErrorResponse("Validation failed", utils.FormatValidationErrors(err)))
		return
	}

	token, err := ctrl.Service.Login(input.Email, input.Password)
	if err != nil {
		c.JSON(401, utils.ErrorResponse("Unauthorized", err.Error()))
		return
	}

	c.JSON(200, utils.SuccessResponse("Login successful", token))
}

func (ctrl *AuthController) Logout(c *gin.Context) {
	token := c.GetString("token")

	err := ctrl.Service.Logout(token)
	if err != nil {
		c.JSON(500, utils.ErrorResponse("Failed to logout", err.Error()))
		return
	}

	c.JSON(200, utils.SuccessResponse("Logged out successfully", nil))
}

func (ctrl *AuthController) Me(c *gin.Context) {
	u, exists := c.Get("user")
	if !exists {
		c.JSON(401, utils.ErrorResponse("Unauthorized", "User not found in context"))
		return
	}

	c.JSON(200, utils.SuccessResponse("User retrieved successfully", u))
}
