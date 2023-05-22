package models

type TrafficDispose struct {
	EntranceFlowLeft  int `json:"entrance_flow_left"`
	EntranceFlowRight int `json:"entrance_flow_right"`
	TunnelFlowLeft    int `json:"tunnel_flow_left"`
	TunnelFlowRight   int `json:"tunnel_flow_right"`
	ExportFlowLeft    int `json:"export_flow_left"`
	ExportFlowRight   int `json:"export_flow_right"`
}
