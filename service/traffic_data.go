package service

import (
	"traffic_flow/api"
	"traffic_flow/models"
)

// 查询当前的车流量
func GetDataOneMin() models.TrafficDispose {
	var ro, lo models.TrafficOrigin
	oneMinData := api.FindDataInOneMin(&ro, &lo)
	api.SendTrafficData(&oneMinData)
	return oneMinData
}
