package models

import "time"

type TrafficOrigin struct {
	Time         time.Time `gorm:"column:time;type:timestamptz;autoUpdateTime"`
	EntranceFlow int       `gorm:"column:entrance_flow;"`
	TunnelFlow   int       `gorm:"column:tunnel_flow"`
	ExportFlow   int       `gorm:"column:export_flow"`
	Direction    bool      `gorm:"column:direction"` //f表示右,t表示左
}
