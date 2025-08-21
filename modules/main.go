package modules

import (
	"go-boilerplate/config"
	"go-boilerplate/modules/auth"
	"go-boilerplate/modules/core"
	"go-boilerplate/modules/friends"
	"go-boilerplate/modules/health"
	"go-boilerplate/modules/profile"
	"go-boilerplate/modules/users"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// Run will start the server
func Run() {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	router.Use(cors.New(corsConfig))
	getRoutes()
	err := router.Run(":" + config.App.ApplicationPort)

	if err != nil {
		log.Fatal(err)
	}
}

func getRoutes() {
	router.RedirectTrailingSlash = false
	rootRoute := router.Group("/")
	rootRoute.GET("/ws", core.WebsocketHandler)
	health.Init(rootRoute)
	users.Init(rootRoute)
	auth.Init(rootRoute)
	profile.Init(rootRoute)
	friends.Init(rootRoute)
}
