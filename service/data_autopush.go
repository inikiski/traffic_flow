package service

import (
	"github.com/robfig/cron/v3"
)

// 每间隔1min,数据推送给前端
func AutoPushEveryMin() {
	crontab := cron.New(cron.WithSeconds())
	spec := "@every 1m"
	task := func() {
		GetDataOneMin()
	}
	crontab.AddFunc(spec, task)
	crontab.Start()
	select {}
}
