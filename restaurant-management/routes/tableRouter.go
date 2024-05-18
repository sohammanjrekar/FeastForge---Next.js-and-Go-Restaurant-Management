package routes

import (
	controller "restaurant-management/controller"

	"github.com/gin-gonic/gin"
)

func tableRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/tables", controller.Gettables())
	incomingRoutes.GET("/tables/:table_id", controller.Gettable())

	incomingRoutes.POST("/tables", controller.Createtable())
	incomingRoutes.PATCH("/tables/:table_id", controller.Updatetable())
}
