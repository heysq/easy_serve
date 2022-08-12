package easy_serve

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

type easyCron struct {
	scheduler *gocron.Scheduler
}

var cron easyCron

func newEasyCron() *easyCron {
	return &easyCron{
		scheduler: gocron.NewScheduler(time.UTC),
	}
}

func (cron *easyCron) Start() {
	cron.scheduler.StartAsync()
}

func (cron *easyCron) Stop() {
	cron.scheduler.Stop()
	fmt.Println("cron exited")
}
