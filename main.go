package main

import (
	"os"
	"restaurant-management/database"
	"restaurant-management/middleware"
	"restaurant-management/routes"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollections *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authencation)

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	router.TableRoutes(router)
	router.OrderRoutes(router)
	router.OrderItemRoutes(router)
	router.InvoiceRoutes(router)

	router.Run(":" + port)

}
