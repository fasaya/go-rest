package router

import (
	"go-rest/controllers"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	router := gin.Default()

	router.GET("/order/:id", controllers.OrderDetail)
	router.GET("/orders", controllers.Orders)
	router.POST("/orders", controllers.OrderCreate)

	return router
}
