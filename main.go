package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"traffic_flow/db"
)

func main() {
	e := echo.New()
	db.Init()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello,traffic_flow")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
