package routes

import (
	"real_estate/src/handlers"
	"real_estate/src/middlewares"

	"github.com/gin-gonic/gin"
)

func PropertyRoutes(router *gin.Engine){ 
	//agent
	router.POST("/category",middlewares.WithAuth(handlers.CreateCategory))
	router.POST("/property",middlewares.WithAuth(handlers.CreateProperty))
	router.GET("/property",middlewares.WithAuth(handlers.GetAllProperties))
	router.GET("/property/:id",handlers.GetProperty)
	router.PUT("/property/:id",middlewares.WithAuth(handlers.UpdateProperty) )
	router.DELETE("/property/:id",middlewares.WithAuth(handlers.DeleteProperty)) 

	//admin
	router.PUT("/approve_property/:id",middlewares.WithAuth(handlers.ApproveProperty))
}