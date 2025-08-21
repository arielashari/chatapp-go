package friends

import (
	"go-boilerplate/modules/core"
	"net/http"

	"github.com/gin-gonic/gin"
)

func routes(rg *gin.RouterGroup) {
	friendsProtectedRoute := rg.Group("/friends")
	friendsProtectedRoute.Use(core.JwtAuthMiddleware())
	friendsProtectedRoute.GET("", getAllFriend)
	friendsProtectedRoute.GET("friend-request", getAllFriendRequests)
	friendsProtectedRoute.POST("friend-request/:userId", sendFriendRequest)
	friendsProtectedRoute.POST("accept-friend-request/:userId", acceptFriendRequest)
	friendsProtectedRoute.POST("reject-friend-request/:userId", rejectFriendRequest)
	friendsProtectedRoute.DELETE("/:userId", deleteFriend)
}

func getAllFriend(c *gin.Context) {
	userIdValue, _ := c.Get("user_id")
	userId, _ := userIdValue.(string)
	err, count, data := GetAllFriends(userId, "1", "30")

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
func getAllFriendRequests(c *gin.Context) {
	userIdValue, _ := c.Get("user_id")
	userId, _ := userIdValue.(string)
	err, count, data := GetFriendRequests(userId, "1", "30")

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

func sendFriendRequest(c *gin.Context) {
	userIdValue, _ := c.Get("user_id")
	userId, _ := userIdValue.(string)

	err, data := SendFriendRequest(userId, c.Param("userId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data":    data,
		"message": "success",
	})
}

func acceptFriendRequest(c *gin.Context) {
	userIdValue, _ := c.Get("user_id")
	userId, _ := userIdValue.(string)

	err, data := AcceptFriendRequest(userId, c.Param("userId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data":    data,
		"message": "success",
	})
}

func rejectFriendRequest(c *gin.Context) {
	userIdValue, _ := c.Get("user_id")
	userId, _ := userIdValue.(string)

	err, data := RejectFriendRequest(userId, c.Param("userId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data":    data,
		"message": "success",
	})
}

func deleteFriend(c *gin.Context) {
	userIdValue, _ := c.Get("user_id")
	userId, _ := userIdValue.(string)

	err, data := DeleteFriend(userId, c.Param("userId"))

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
