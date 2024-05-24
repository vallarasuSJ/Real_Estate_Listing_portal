package handlers

import (
	"net/http"
	"real_estate/src/dtos"
	bookings "real_estate/src/services/Bookings"
	"real_estate/src/utils/context"

	"github.com/gin-gonic/gin"
)

func CreateBookings(c *context.Context){
	req:=&dtos.BookingReq{}
	if err:=c.Bind(req);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
            "error":err.Error(),
        })
        return
	}
	if c.Users.RoleName != "customer" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Customer only can access..",
		})
		return
	}
	err:=bookings.New().CreateBookings(c,req)

	if err!=nil{
		c.JSON(http.StatusUnauthorized,gin.H{
			"error":err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated,gin.H{
		"message":"Property booked successfully",
	})
}

func GetBooking(c *context.Context) {
	booking,err:= bookings.New().GetBooking(c)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Booking not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Booking": booking,
	})
} 

func CancelBooking(c *context.Context){
	id := c.Param("id")
	if c.Users.RoleName != "customer" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Customer only can access..",
		})
		return
	}
	err:= bookings.New().CancelBooking(c, id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Booking not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Booking": "Booking cancelled successfully",
	})
}