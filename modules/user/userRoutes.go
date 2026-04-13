package user

import "github.com/gin-gonic/gin"

func RegisterUserRoutes(rg *gin.RouterGroup, authMiddleware gin.HandlerFunc) {
	userService := UserService{}
	userCtrl := UserController{Service: userService}

	userGroup := rg.Group("/users")
	userGroup.Use(authMiddleware)
	{
		userGroup.GET("/", userCtrl.GetUsers)
		userGroup.POST("/", userCtrl.CreateUser)
		userGroup.GET("/:id", userCtrl.GetUser)
		userGroup.PUT("/:id", userCtrl.UpdateUser)
		userGroup.DELETE("/:id", userCtrl.DeleteUser)
	}
}