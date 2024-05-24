package context

import (
	"real_estate/src/dtos"
	"real_estate/src/utils/db"

	"github.com/gin-gonic/gin"
)

type Context struct {
    *gin.Context
	DB *db.DB  
	*dtos.Users
	
}
