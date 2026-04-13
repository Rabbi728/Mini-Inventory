package auth

import "github.com/gin-gonic/gin"

func RegisterAuthRoutes(r *gin.RouterGroup) {
	authService := AuthService{}
	authCtrl := AuthController{Service: authService}

	authGroup := r.Group("/")
	{
		authGroup.POST("register", authCtrl.Register)
		authGroup.POST("login", authCtrl.Login)
		
		protected := authGroup.Group("/")
		protected.Use(AuthMiddleware())
		{
			protected.POST("logout", authCtrl.Logout)
			protected.GET("me", authCtrl.Me)
		}
	}
}
