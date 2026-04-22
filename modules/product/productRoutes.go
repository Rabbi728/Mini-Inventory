package product

import "github.com/gin-gonic/gin"

func RegisterProductRoutes(rg *gin.RouterGroup, authMiddleware gin.HandlerFunc) {
	productService := ProductService{}
	productCtrl := ProductController{Service: productService}

	productGroup := rg.Group("/products")
	productGroup.Use(authMiddleware)
	{
		productGroup.GET("/", productCtrl.GetProducts)
		productGroup.POST("/", productCtrl.CreateProduct)
		productGroup.GET("/:id", productCtrl.GetProduct)
		productGroup.PUT("/:id", productCtrl.UpdateProduct)
		productGroup.DELETE("/:id", productCtrl.DeleteProduct)
	}
}
