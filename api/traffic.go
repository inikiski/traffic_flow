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
