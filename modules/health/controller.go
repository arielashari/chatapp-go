package health

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func routes(rg *gin.RouterGroup) {
	healthRoute := rg.Group("/health")

	healthRoute.GET("", healthCheck)
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
