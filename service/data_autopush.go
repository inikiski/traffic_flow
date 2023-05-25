package service

import (
	"log"
	"time"
)

// 每间隔1min,数据推送给前端
func AutoPushEveryMin() {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			GetDataOneMin()
			log.Println("Auto push")
		}
	}
}
