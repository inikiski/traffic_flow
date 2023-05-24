package service

import (
	"traffic_flow/api"
	"traffic_flow/models"
)

// 查询当前的车流量
func GetDataOneMin() models.TrafficDispose {
	var o1, o2 models.TrafficOrigin
	oneMinData := api.FindDataInOneMin(o1, o2)
	api.SendTrafficData(&oneMinData)
	return oneMinData
}
