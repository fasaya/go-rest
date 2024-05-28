package controllers

import (
	"fmt"
	"go-rest/database"
	"go-rest/helpers"
	"go-rest/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	appJSON = "application/json"
)

func OrderCreate(ctx *gin.Context) {
	contentType := helpers.GetContentType(ctx)

	Order := models.Order{}

	if contentType == appJSON {
		ctx.ShouldBindJSON(&Order)
	} else {
		ctx.ShouldBind(&Order)
	}

	fmt.Printf("%+v\n", Order)

	db := database.GetDB()

	err := db.Transaction(func(tx *gorm.DB) error {

		// Create order
		if err := tx.Debug().Create(&Order).Error; err != nil {
			return err
		}

		// Create order items
		// for _, orderItem := range Order.Items {
		// 	orderItem.ID = 0
		// 	orderItem.OrderID = Order.ID
		// 	fmt.Println("aaa:", orderItem, "bbb:", Order.ID)
		// 	if err := tx.Debug().Create(&orderItem).Error; err != nil {
		// 		return err
		// 	}
		// }

		// return nil will commit the whole transaction
		return nil
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Order created successfully",
		"data":    Order,
	})
}

func Orders(ctx *gin.Context) {
	db := database.GetDB()

	var Orders []models.Order

	err := db.Preload("Items").Find(&Orders).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Failed to fetch orders",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Order fetched successfully",
		"data":    Orders,
	})
}

func OrderDetail(ctx *gin.Context) {
	db := database.GetDB()

	var Order []models.Order

	OrderID := ctx.Param("id")

	err := db.Preload("Items").First(&Order, OrderID).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Order not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Order detail fetched successfully",
		"data":    Order,
	})
}
