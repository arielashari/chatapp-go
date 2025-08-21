package message

import (
	"go-boilerplate/modules/core"
	"go-boilerplate/modules/message/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func routes(rg *gin.RouterGroup) {
	messageProtectedRoute := rg.Group("/message")
	messageProtectedRoute.Use(core.JwtAuthMiddleware())
	messageProtectedRoute.GET("", getProfileDetail)
	messageProtectedRoute.PUT("/:id", updateProfile)
	messageProtectedRoute.DELETE("/:id", deleteProfile)
}

func getProfileDetail(c *gin.Context) {
	userIdValue, _ := c.Get("user_id")
	userId, ok := userIdValue.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user ID"})
		return
	}

	err, data := GetProfileDetail(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"message": "success",
	})
}

func updateProfile(c *gin.Context) {
	var input dto.UpdateProfileDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err, data := UpdateProfile(c.Param("id"), input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"message": "success",
	})
}

func deleteProfile(c *gin.Context) {
	err, data := DeleteProfile(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"message": "success",
	})
}
