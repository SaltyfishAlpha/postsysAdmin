// @title        Golang Service Template
// @version      0.1
// @description  Golang back-end service template, get started with back-end projects quickly
// @BasePath     /api

package main

import (
	"Express/app"
	"Express/model"
)

func main() {
	model.Init()
	app.InitWebFramework()
	app.StartServer()
}
