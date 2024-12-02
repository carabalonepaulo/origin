package sys

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/carabalonepaulo/origin/shared/service"
)

type Service struct {
	ch       chan os.Signal
	shutdown func()
}

func New() service.Service {
	return &Service{}
}

func (s *Service) Start(services service.Services, shutdown func()) error {
	s.shutdown = shutdown
	s.ch = make(chan os.Signal, 1)
	signal.Notify(s.ch, syscall.SIGINT, syscall.SIGTERM)

	return nil
}

func (s *Service) Stop() {}

func (s *Service) Update(dt float64) {
	select {
	case <-s.ch:
		s.shutdown()
	default:
	}
}
