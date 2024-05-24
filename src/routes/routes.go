package routes
 
import (
    "real_estate/src/middlewares"
    "github.com/gin-gonic/gin"
)
 
func GetRouter() *gin.Engine {
    router := gin.Default()
 
    router.Use(middlewares.CORSMiddleware())
 
    pingRoutes(router)
    UserRoutes(router) 
    PropertyRoutes(router)
    BookingRoutes(router)
 
    return router
}