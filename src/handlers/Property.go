package handlers

import (
	"net/http"
	"real_estate/src/dtos"
	properties "real_estate/src/services/Properties"
	"real_estate/src/services/categories"
	"real_estate/src/utils/context"

	"github.com/gin-gonic/gin"
)

func CreateCategory(c *context.Context) {
	if c.Users.RoleName != "agent" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Agent only can access..",
		})
		return
	}

	req := &dtos.CategoryReq{}
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := categories.New().CreateCategory(c, req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Category created successfully",
	})
}

func CreateProperty(c *context.Context) {
	if c.Users.RoleName != "agent" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Agent only can access..",
		})
		return
	}
	req := &dtos.PropertyReq{}
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	err := properties.New().CreateProperty(c, req)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Property created successfully",
	})
}

func GetAllProperties(c *context.Context) { 
	properties, err := properties.New().GetAllProperties(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Properties not found",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"Properties": properties,
	})
}

func GetProperty(c *gin.Context) {
	id := c.Param("id")
	property, err := properties.New().GetProperty(getContext(c), id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Property not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Property": property,
	})
}

func DeleteProperty(c *context.Context) {
	if c.Users.RoleName != "agent" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Agent only can access..",
		})
		return
	}
	id := c.Param("id")

	err := properties.New().DeleteProperty(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete property",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Property": "Property deleted successfully",
	})
}

func UpdateProperty(c *context.Context) {
	if c.Users.RoleName != "agent" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Agent only can access..",
		})
		return
	}
	id := c.Param("id")
	req := &dtos.PropertyReq{}
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	

	err := properties.New().UpdateProperty(c, req, id)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Property updated successfully",
	})
}

func ApproveProperty(c *context.Context) {
	if c.Users.RoleName != "admin" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Admin only can access..",
		})
		return
	}
	id := c.Param("id")
	err := properties.New().ApproveProperty(c, id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Property not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Property": "Property has been approved",
	})

}
