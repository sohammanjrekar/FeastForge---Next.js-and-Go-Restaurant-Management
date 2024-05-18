package routes

import (
	controller "restaurant-management/controller"

	"github.com/gin-gonic/gin"
)

func orderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderItems", controller.GetorderItems())
	incomingRoutes.GET("/orderItems/:orderItem_id", controller.GetorderItem())
	incomingRoutes.GET("/orderItems-order/:order_id", controller.GetorderItemsByOrder())

	incomingRoutes.POST("/orderItems", controller.CreateorderItem())
	incomingRoutes.PATCH("/orderItems/:orderItem_id", controller.UpdateorderItem())
}
