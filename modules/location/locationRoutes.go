package location

import "github.com/gin-gonic/gin"

func RegisterLocationRoutes(rg *gin.RouterGroup, authMiddleware gin.HandlerFunc) {
	locationService := LocationService{}
	locationCtrl := LocationController{Service: locationService}

	locationGroup := rg.Group("/locations")
	locationGroup.Use(authMiddleware)
	{
		locationGroup.GET("/", locationCtrl.GetLocations)
		locationGroup.POST("/", locationCtrl.CreateLocation)
		locationGroup.GET("/:id", locationCtrl.GetLocation)
		locationGroup.PUT("/:id", locationCtrl.UpdateLocation)
		locationGroup.DELETE("/:id", locationCtrl.DeleteLocation)
	}
}
