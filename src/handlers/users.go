package handlers

import (
	"net/http"

	"real_estate/src/dtos"
	"real_estate/src/services/users"
	"real_estate/src/utils/context"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context){
	req:=&dtos.RegisterReq{}
	if err:=c.Bind(req); err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	} 

	err:=users.New().Register(getContext(c),req)
	if err!=nil{
		c.JSON(http.StatusUnauthorized,gin.H{
			"error":err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Registration successful",
	})

}

func Login(c *gin.Context){
	req:=&dtos.LoginReq{}
	if err:=c.Bind(req);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}
	
	res,err:=users.New().Login(getContext(c),req)
	if err!=nil{
		c.JSON(http.StatusUnauthorized,gin.H{
			"error":err.Error(),
		})
		return
	}  
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"Response":res,
	})
}

func GetAccountByUsingAccessToken(c *context.Context) {
	c.JSON(http.StatusOK, c.Users)
}

func GetAccessFromRefreshToken(c *gin.Context) {
	at, err := users.New().GetAccessFromRefreshToken(getContext(c), c.Query("refresh-token"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": at,
	})
}