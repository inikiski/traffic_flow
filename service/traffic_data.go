package service

import (
	"github.com/nats-io/nats.go"
	"traffic_flow/models"
)

func SendTrafficData(dispose *models.TrafficDispose) {
	nc, _ := nats.Connect("nats://39.108.214.230:4222")
	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer c.Close()

	err := c.Publish("TRAFFIC-FLOW.data", dispose)
	if err != nil {
		panic("发送错误")
	}
}

func GetDataByDay() {

}
