package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"traffic_flow/db"
	"traffic_flow/service"
	"traffic_flow/sim"
)

func main() {
	e := echo.New()
	db.Init()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello,traffic_flow")
	})
	go service.AutoPushEveryMin()
	go sim.Sensor()
	go e.Logger.Fatal(e.Start(":8080"))
}
