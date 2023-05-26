package service

import (
	"github.com/labstack/echo/v4"
	"traffic_flow/api"
	"traffic_flow/models"
)

// 查询当前的车流量
func GetDataOneMin(e echo.Context) error {
	var ro, lo models.TrafficOrigin
	oneMinData := api.FindDataInOneMin(&ro, &lo)
	api.SendTrafficData(&oneMinData)
	return e.JSON(200, oneMinData)
}
