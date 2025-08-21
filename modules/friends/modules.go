package friends

import (
	"go-boilerplate/modules/core"
	"go-boilerplate/modules/friends/models"

	"github.com/gin-gonic/gin"
)

func Init(rg *gin.RouterGroup) {
	initModel()
	routes(rg)
}

func initModel() {
	core.AutoMigrate("friends", &models.Friends{})
	core.AutoMigrate("friend_requests", &models.FriendRequests{})
}
