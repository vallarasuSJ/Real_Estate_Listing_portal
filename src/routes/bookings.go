package routes

import (
	"real_estate/src/handlers"
	"real_estate/src/middlewares"

	"github.com/gin-gonic/gin"
)

func BookingRoutes(router *gin.Engine){
	router.POST("/booking",middlewares.WithAuth(handlers.CreateBookings))
	router.DELETE("/booking/:id",middlewares.WithAuth(handlers.CancelBooking))
	router.GET("/booking",middlewares.WithAuth(handlers.GetBooking))   
}