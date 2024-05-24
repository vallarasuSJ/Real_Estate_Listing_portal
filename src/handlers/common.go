package handlers

import (
	"real_estate/src/utils/context"
	"real_estate/src/utils/db"
	"github.com/gin-gonic/gin"
)

func getContext(c *gin.Context) *context.Context{
	return &context.Context{
		Context:c, 
		DB: db.New(),
	}
}