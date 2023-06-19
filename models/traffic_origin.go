package models

type TrafficOrigin struct {
	EntranceFlow int  `gorm:"column:entrance_flow;"`
	TunnelFlow   int  `gorm:"column:tunnel_flow"`
	ExportFlow   int  `gorm:"column:export_flow"`
	Direction    bool `gorm:"column:direction"` //f表示右right,t表示左left
}

func (to *TrafficOrigin) TableName() string {
	return "db_traffic"
}
