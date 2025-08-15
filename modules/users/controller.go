package users

import (
	"go-boilerplate/modules/core"
	"go-boilerplate/modules/users/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func routes(rg *gin.RouterGroup) {
	usersRoute := rg.Group("/users")

	usersRoute.GET("", getAllUser)
	usersRoute.GET("/:id", getUserDetail)

	usersProtectedRoute := rg.Group("/users")
	usersProtectedRoute.Use(core.JwtAuthMiddleware())
	usersProtectedRoute.POST("", createNewUser)
	usersProtectedRoute.PUT("/:id", updateUser)
	usersProtectedRoute.DELETE("/:id", deleteUser)
}

func getAllUser(c *gin.Context) {
	err, count, data := GetAllUser("1", "30")

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

func getUserDetail(c *gin.Context) {
	err, data := GetUserDetail(c.Param("id"))

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

func createNewUser(c *gin.Context) {
	var input dto.UserDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err, data := CreateUser(input)

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

func updateUser(c *gin.Context) {
	var input dto.UserDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err, data := UpdateUser(c.Param("id"), input)

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

func deleteUser(c *gin.Context) {
	err, data := DeleteUser(c.Param("id"))

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
