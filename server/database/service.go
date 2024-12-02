package database

import (
	"time"

	"github.com/carabalonepaulo/origin/server/config"
	"github.com/carabalonepaulo/origin/shared/service"
	"github.com/carabalonepaulo/origin/shared/services/scheduler"
	"github.com/carabalonepaulo/origin/shared/weave"
)

type Service struct {
	config *config.Database
}

func New(config *config.Database) func() service.Service {
	return func() service.Service {
		return &Service{
			config: config,
		}
	}
}

func (s *Service) Start(services service.Services, shutdown func()) error {
	scheduler, err := service.Get[*scheduler.Service](services)
	if err != nil {
		return err
	}

	interval, err := time.ParseDuration(s.config.TickInterval)
	scheduler.Every(interval, s.poll)

	var value int
	task := weave.NewChain[int](&value, 10).Add(0, func(value *int) {})
	scheduler.Dispatch(task)

	return nil
}

func (s *Service) Stop() {
	// TODO: wait for all delayed tasks
}

func (s *Service) poll() {}
