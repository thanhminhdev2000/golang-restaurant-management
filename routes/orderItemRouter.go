package routes

import (
	controller "golang-restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.GET("/orderItems", controller.GetOrderItems())
	incommingRoutes.GET("/orderItems/:orderItem_id", controller.GetOrderItem())
	incommingRoutes.GET("/orderItems-order/:order_id", controller.GetOrderItemByOrder())
	incommingRoutes.POST("/orderItems", controller.CreateOrderItem())
	incommingRoutes.PATCH("/orderItems:orderItem_id", controller.UpdateOrderItem())
}
