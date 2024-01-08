package app

import (
	"Express/app/response"
	"github.com/labstack/echo/v4"
	//echoSwagger "github.com/swaggo/echo-swagger"
	// "github.com/labstack/echo/v4/middleware"
	"Express/app/controller"
)

func ping(c echo.Context) error {
	return response.SendResponse(c, 200, "pong!")
	//return c.String(200, "pong!")
}

func addRoutes() {
	//visit := e.Group("visit", middleware.RedictMiddleware)
	//visit.GET("/:hash", controller.Visit)
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"http://localhost:5060", "http://localhost:1926"},
	// 	AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	// }))
	api := e.Group("api")

	api.GET("/ping", ping)
	//api.POST("/url/Create", controller.CreateUrl)

	api.POST("/upload", controller.UploadExpress)
	api.GET("/allocate", controller.Allocate)
	api.POST("/allocated", controller.Allocated)

	api.POST("/query", controller.QueryMine)
	api.GET("/remove", controller.PickUp)

	api.POST("/info", controller.BeginSend)
	api.GET("/getinfo", controller.GetInfo)
	api.POST("/updateinfo", controller.UpdateInfos)

}
