package product

import (
	"mini-inventory/utils"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	Service ProductService
}

func (ctrl *ProductController) GetProducts(c *gin.Context) {
	products, err := ctrl.Service.GetAllProducts()
	if err != nil {
		c.JSON(500, utils.ErrorResponse("Failed to fetch products", err.Error()))
		return
	}
	c.JSON(200, utils.SuccessResponse("Products retrieved successfully", products))
}

func (ctrl *ProductController) CreateProduct(c *gin.Context) {
	var product Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(400, utils.ErrorResponse("Validation failed", utils.FormatValidationErrors(err)))
		return
	}

	if err := ctrl.Service.CreateProduct(&product); err != nil {
		c.JSON(500, utils.ErrorResponse("Failed to create product", err.Error()))
		return
	}

	c.JSON(201, utils.SuccessResponse("Product created successfully", product))
}

func (ctrl *ProductController) GetProduct(c *gin.Context) {
	id := c.Param("id")
	product, err := ctrl.Service.GetProductByID(id)
	if err != nil {
		c.JSON(404, utils.ErrorResponse("Product not found", nil))
		return
	}
	c.JSON(200, utils.SuccessResponse("Product retrieved successfully", product))
}

func (ctrl *ProductController) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	product, err := ctrl.Service.GetProductByID(id)
	if err != nil {
		c.JSON(404, utils.ErrorResponse("Product not found", nil))
		return
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(400, utils.ErrorResponse("Validation failed", utils.FormatValidationErrors(err)))
		return
	}

	if err := ctrl.Service.UpdateProduct(&product); err != nil {
		c.JSON(500, utils.ErrorResponse("Failed to update product", err.Error()))
		return
	}

	c.JSON(200, utils.SuccessResponse("Product updated successfully", product))
}

func (ctrl *ProductController) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if err := ctrl.Service.DeleteProduct(id); err != nil {
		c.JSON(500, utils.ErrorResponse("Failed to delete product", err.Error()))
		return
	}
	c.JSON(200, utils.SuccessResponse("Product deleted successfully", nil))
}
