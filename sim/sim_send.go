package sim

import (
	"github.com/duke-git/lancet/v2/random"
	"strings"
	"time"
	"traffic_flow/api"
	"traffic_flow/models"
)

// 生成随机检测器数据
func CreateRandOriData(location string, direction string) *models.TrafficOrigin {
	var ori models.TrafficOrigin
	if strings.Compare(location, "entrance") == 0 {
		ori.EntranceFlow = random.RandInt(1, 20)
	} else if strings.Compare(location, "export") == 0 {
		ori.ExportFlow = random.RandInt(1, 20)
	}
	if strings.Compare(direction, "right") == 0 {
		ori.Direction = false
	} else if strings.Compare(direction, "left") == 0 {
		ori.Direction = true
	}

	return &ori
}

// 发送到nats
func Sensor() {
	c := make(chan int, 1)
	go func() {
		time.Sleep(500 * time.Millisecond)
		c <- random.RandInt(0, 3)
	}()

	for {
		select {
		case num := <-c:
			if num == 0 {
				data := CreateRandOriData("entrance", "right")
				api.SendOriginData(data)
			} else if num == 1 {
				data := CreateRandOriData("entrance", "left")
				api.SendOriginData(data)
			} else if num == 2 {
				data := CreateRandOriData("export", "left")
				api.SendOriginData(data)
			} else if num == 3 {
				data := CreateRandOriData("export", "right")
				api.SendOriginData(data)
			}
		}
	}

}
