package middlewares

import (
	"real_estate/src/utils/context"
	"real_estate/src/utils/db"
	"strings"

	"github.com/gin-gonic/gin"
)

func getContext(c *gin.Context) *context.Context {
	return &context.Context{
		Context: c,
		DB:      db.New(),
	}
}

func GetBearerToken(c *gin.Context) string {
	bearerToken := c.Request.Header.Get("Authorization")
	strs := strings.Split(bearerToken, " ")
	if len(strs) > 1 {
		return strs[1]
	}
	return ""
}

// func AccountFromRequest(req *dtos.UserReq) *dtos.Users {
// 	return &dtos.Users{
// 		Id:             req.Id,
// 		Email:          req.Email,
// 		Contact_number: req.Contact_number,
// 		UserName:       req.Username,
// 		Role:           req.Role_id,
// 		Gender:         req.Gender,
// 		RoleName:      req.RoleName,
// 	}
// }
