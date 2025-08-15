package modules

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-boilerplate/config"
	"go-boilerplate/modules/auth"
	"go-boilerplate/modules/health"
	"go-boilerplate/modules/users"
	"log"
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
	health.Init(rootRoute)
	users.Init(rootRoute)
	auth.Init(rootRoute)
}
