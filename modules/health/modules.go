package health

import (
	"github.com/gin-gonic/gin"
)

func Init(rg *gin.RouterGroup) {
	routes(rg)
}
