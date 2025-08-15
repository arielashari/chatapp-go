package friends

import (
	"go-boilerplate/modules/core"
	"net/http"

	"github.com/gin-gonic/gin"
)

func routes(rg *gin.RouterGroup) {
	friendsProtectedRoute := rg.Group("/friends")
	friendsProtectedRoute.Use(core.JwtAuthMiddleware())
	friendsProtectedRoute.GET("list", getAllFriend)
}

func getAllFriend(c *gin.Context) {
	err, count, data := GetAllFriends("1", "30")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"count":   count,
		"message": "success",
	})
}
