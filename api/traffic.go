package api

import (
	"time"
	"traffic_flow/db"
	"traffic_flow/models"
	"traffic_flow/utils"
)

func FindDataInOneMin(ro, lo *models.TrafficOrigin) models.TrafficDispose {
	sql1 := "SELECT \"sum\"(entrance_flow) AS \"entrance_flow\",\"sum\"(tunnel_flow) AS \"tunnel_flow\",\"sum\"(export_flow) AS \"export_flow\" FROM \"traffic\".\"db_traffic_left\" WHERE time > TIMESTAMP ?"
	sql2 := "SELECT \"sum\"(entrance_flow) AS \"entrance_flow\",\"sum\"(tunnel_flow) AS \"tunnel_flow\",\"sum\"(export_flow) AS \"export_flow\" FROM \"traffic\".\"db_traffic_right\" WHERE time > TIMESTAMP ?"
	utcMinusOneMin := time.Now().UTC().Add(-time.Minute)
	db.DB.Raw(sql1, utcMinusOneMin).Scan(ro)
	db.DB.Raw(sql2, utcMinusOneMin).Scan(lo)
	return utils.MergeLeftRight(ro, lo)
}

func InsertOri(dir string, ori *models.TrafficOrigin) {
	var sql string
	if dir == "left" {
		sql = "INSERT INTO \"traffic\".\"db_traffic_left\" ( entrance_flow, tunnel_flow, export_flow) VALUES (?, ?, ?);"
	} else if dir == "right" {
		sql = "INSERT INTO \"traffic\".\"db_traffic_right\" ( entrance_flow, tunnel_flow, export_flow) VALUES (?, ?, ?);"
	}
	db.DB.Raw(sql, &ori.EntranceFlow, &ori.TunnelFlow)
}
