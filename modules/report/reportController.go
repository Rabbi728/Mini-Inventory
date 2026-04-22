package report

import (
	"mini-inventory/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReportController struct {
	Service ReportService
}

func (ctrl *ReportController) StockRegister(c *gin.Context) {
	productID, startDate, endDate := ctrl.getParams(c)
	records, err := ctrl.Service.GetStockRegister(productID, startDate, endDate)
	if err != nil {
		c.JSON(500, utils.ErrorResponse("Failed to fetch stock register", err.Error()))
		return
	}
	c.JSON(200, utils.SuccessResponse("Stock register fetched successfully", records))
}

func (ctrl *ReportController) ReceiveReport(c *gin.Context) {
	productID, startDate, endDate := ctrl.getParams(c)
	records, err := ctrl.Service.GetReceiveReport(productID, startDate, endDate)
	if err != nil {
		c.JSON(500, utils.ErrorResponse("Failed to fetch receive report", err.Error()))
		return
	}
	c.JSON(200, utils.SuccessResponse("Receive report fetched successfully", records))
}

func (ctrl *ReportController) DeliveryReport(c *gin.Context) {
	productID, startDate, endDate := ctrl.getParams(c)
	records, err := ctrl.Service.GetDeliveryReport(productID, startDate, endDate)
	if err != nil {
		c.JSON(500, utils.ErrorResponse("Failed to fetch delivery report", err.Error()))
		return
	}
	c.JSON(200, utils.SuccessResponse("Delivery report fetched successfully", records))
}

func (ctrl *ReportController) getParams(c *gin.Context) (uint, string, string) {
	productIDStr := c.Query("productID")
	if productIDStr == "" {
		productIDStr = c.Query("product_id")
	}

	var productID uint
	if productIDStr != "" {
		pid, _ := strconv.ParseUint(productIDStr, 10, 32)
		productID = uint(pid)
	}

	startDate := c.Query("startDate")
	if startDate == "" {
		startDate = c.Query("start_date")
	}

	endDate := c.Query("endDate")
	if endDate == "" {
		endDate = c.Query("end_date")
	}

	return productID, startDate, endDate
}
