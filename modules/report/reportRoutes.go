package report

import (
	"github.com/gin-gonic/gin"
)

func RegisterReportRoutes(router *gin.RouterGroup, authMiddleware gin.HandlerFunc) {
	ctrl := &ReportController{
		Service: ReportService{},
	}

	reportGroup := router.Group("/reports")
	reportGroup.Use(authMiddleware)
	{
		reportGroup.GET("/stock-register", ctrl.StockRegister)
		reportGroup.GET("/receive", ctrl.ReceiveReport)
		reportGroup.GET("/delivery", ctrl.DeliveryReport)
	}
}
