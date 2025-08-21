package message

import (
	"go-boilerplate/modules/core"
	"go-boilerplate/modules/message/models"

	"github.com/gin-gonic/gin"
)

func Init(rg *gin.RouterGroup) {
	initModel()
	routes(rg)
}

func initModel() {
	core.AutoMigrate("messages", &models.Message{})
	core.AutoMigrate("rooms", &models.Room{})
	core.AutoMigrate("room_members", &models.RoomMember{})
}
