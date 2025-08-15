package profile

import (
	"go-boilerplate/modules/core"
	"go-boilerplate/modules/profile/models"

	"github.com/gin-gonic/gin"
)

func Init(rg *gin.RouterGroup) {
	initModel()
	routes(rg)
}

func initModel() {
	core.AutoMigrate("profile", &models.Profile{})
}
