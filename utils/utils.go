package utils

import "traffic_flow/models"

// 将两条双通道原始数据转为单条数据
func MergeLeftRight(o1, o2 models.TrafficOrigin) models.TrafficDispose {
	var d models.TrafficDispose //false == right   true == left
	if o1.Direction == false && o2.Direction == true {
		d.EntranceFlowLeft = o2.EntranceFlow
		d.EntranceFlowRight = o1.EntranceFlow
		d.TunnelFlowLeft = o2.TunnelFlow
		d.TunnelFlowRight = o1.TunnelFlow
		d.ExportFlowLeft = o2.ExportFlow
		d.ExportFlowRight = o1.ExportFlow
	} else if o1.Direction == true && o2.Direction == false {
		d.EntranceFlowLeft = o1.EntranceFlow
		d.EntranceFlowRight = o2.EntranceFlow
		d.TunnelFlowLeft = o1.TunnelFlow
		d.TunnelFlowRight = o2.TunnelFlow
		d.ExportFlowLeft = o1.ExportFlow
		d.ExportFlowRight = o2.ExportFlow
	}
	return d
}
