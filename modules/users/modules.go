package users

import (
	"github.com/gin-gonic/gin"
	"go-boilerplate/modules/core"
	"go-boilerplate/modules/users/models"
)

func Init(rg *gin.RouterGroup) {
	initModel()
	routes(rg)
}

func initModel() {
	core.AutoMigrate("users", &models.User{})
}
