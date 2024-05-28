package router

import (
	"go-rest/controllers"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	router := gin.Default()

	router.GET("/order/:id", controllers.OrderDetail)
	router.PUT("/orders/:id", controllers.OrderUpdate)
	router.DELETE("/orders/:id", controllers.OrderDelete)
	router.GET("/orders", controllers.Orders)
	router.POST("/orders", controllers.OrderCreate)

	return router
}
