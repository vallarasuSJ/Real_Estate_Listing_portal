package routes

import (
	"real_estate/src/handlers"
	"real_estate/src/middlewares"
	"github.com/gin-gonic/gin"
)
 
func pingRoutes(router *gin.Engine) {
    router.Use(middlewares.CORSMiddleware())
    router.GET("/ping", handlers.Ping)
}
 