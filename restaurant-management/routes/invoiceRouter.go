package routes

import (
	controller "restaurant-management/controller"

	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/invoices", controller.Getinvoices())
	incomingRoutes.GET("/invoices/:invoice_id", controller.Getinvoice())

	incomingRoutes.POST("/invoices", controller.Createinvoice())
	incomingRoutes.PATCH("/invoices/:invoice_id", controller.Updateinvoice())
}
