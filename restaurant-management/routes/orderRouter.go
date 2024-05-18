package routes

import (
	controller "restaurant-management/controller"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orders", controller.Getorders())
	incomingRoutes.GET("/orders/:order_id", controller.Getorder())

	incomingRoutes.POST("/orders", controller.Createorder())
	incomingRoutes.PATCH("/orders/:order_id", controller.Updateorder())
}
