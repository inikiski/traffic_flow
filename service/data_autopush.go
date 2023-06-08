package service

import (
	"github.com/labstack/echo/v4"
	"log"
	"time"
)

// 每间隔1min,数据推送给前端
func AutoPushEveryMin(e echo.Context) {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			GetDataOneMin(e)
			log.Println("Auto push")
		}
	}
}
