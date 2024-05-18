package routes

import (
	controller "restaurant-management/controller"

	"github.com/gin-gonic/gin"
)

func MenuRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/menus", controller.Getmenus())
	incomingRoutes.GET("/menus/:menu_id", controller.Getmenu())

	incomingRoutes.POST("/menus", controller.Createmenu())
	incomingRoutes.PATCH("/menus/:menu_id", controller.Updatemenu())
}
