package middlewares

import (
	"real_estate/src/utils/context"
	"net/http"
	"real_estate/src/constants"
	"real_estate/src/services/users"

	"github.com/gin-gonic/gin"
)

func WithAuth(next func(*context.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := GetBearerToken(c)
		ctx := getContext(c)

		users, err := users.New().GetAccountWithAccessToken(ctx, token)
		if err == constants.ErrAccessTokenExpire {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.Users=users
		next(ctx)
	}
}


