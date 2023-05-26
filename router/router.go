package router

import (
	"github.com/labstack/echo/v4"
	"traffic_flow/service"
)

func Router(e *echo.Echo) {
	//实时查询接口
	e.GET("/get/realTimeData", service.GetDataOneMin)
}
