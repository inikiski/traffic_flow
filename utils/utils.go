package utils

import "traffic_flow/models"

// 将两条双通道原始数据转为单条数据
func MergeLeftRight(ro, lo *models.TrafficOrigin) models.TrafficDispose {
	var d models.TrafficDispose //false == right   true == left
	d.EntranceFlowLeft = lo.EntranceFlow
	d.EntranceFlowRight = ro.EntranceFlow
	d.TunnelFlowLeft = lo.TunnelFlow
	d.TunnelFlowRight = ro.TunnelFlow
	d.ExportFlowLeft = lo.ExportFlow
	d.ExportFlowRight = ro.ExportFlow
	return d
}
