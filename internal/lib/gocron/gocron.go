package gocron

// 1
import (
	"time"

	"github.com/go-co-op/gocron"
)

var s *gocron.Scheduler

func CreateSchedule() *gocron.Scheduler {
	s = gocron.NewScheduler(time.UTC)
	return s
}

func GetScheduler() *gocron.Scheduler {
	return s
}
