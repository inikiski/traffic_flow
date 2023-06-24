package api

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"traffic_flow/models"
)

func ReceiveOriData(subject string) *models.TrafficOrigin {
	nc, _ := nats.Connect("nats://39.108.214.230:4222")
	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer c.Close()
	var OriData models.TrafficOrigin
	c.Subscribe(subject, func(m *nats.Msg) {
		err := json.Unmarshal(m.Data, &OriData)
		if err != nil {
			return
		}
	})
	return &OriData
}
