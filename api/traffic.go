package api

import (
	"time"
	"traffic_flow/db"
	"traffic_flow/models"
	"traffic_flow/utils"
)

func FindDataInOneMin(o1, o2 models.TrafficOrigin) models.TrafficDispose {
	sql1 := "SELECT \"sum\"(entrance_flow) AS \"entrance_flow\",\"sum\"(tunnel_flow) AS \"tunnel_flow\",\"sum\"(export_flow) AS \"export_flow\" FROM \"traffic\".\"db_traffic\" WHERE time > TIMESTAMP ? - INTERVAL '1 minutes' AND direction = TRUE"
	sql2 := "SELECT \"sum\"(entrance_flow) AS \"entrance_flow\",\"sum\"(tunnel_flow) AS \"tunnel_flow\",\"sum\"(export_flow) AS \"export_flow\" FROM \"traffic\".\"db_traffic\" WHERE time > TIMESTAMP ? - INTERVAL '1 minutes' AND direction = FALSE"
	db.DB.Raw(sql1, time.UTC).Scan(o1)
	o1.Direction = true
	db.DB.Raw(sql2, time.UTC).Scan(o2)
	o2.Direction = false
	return utils.MergeLeftRight(o1, o2)
}
