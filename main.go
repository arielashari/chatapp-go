package main

import (
	"go-boilerplate/config"
	"go-boilerplate/modules"
	"go-boilerplate/modules/core"
)

func main() {
	config.LoadConfig()
	core.InitDatabase()
	modules.Run()
}
