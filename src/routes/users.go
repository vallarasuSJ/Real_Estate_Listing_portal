package routes

import (
	"real_estate/src/handlers"
	"real_estate/src/middlewares"

	"github.com/gin-gonic/gin"
)
 
func UserRoutes(router *gin.Engine) {
	router.GET("/account",middlewares.WithAuth(handlers.GetAccountByUsingAccessToken))
    router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login) 
	router.POST("/role",handlers.RegisterRoles)
}
 