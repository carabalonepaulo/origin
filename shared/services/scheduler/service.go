package scheduler

import (
	"time"

	"github.com/carabalonepaulo/origin/shared/service"
	"github.com/carabalonepaulo/origin/shared/slab"
	"github.com/carabalonepaulo/origin/shared/stack"
	"github.com/carabalonepaulo/origin/shared/weave"
)

type (
	Config struct {
		MaxTimers    int    `json:"max_timers"`
		MaxWorkers   int    `json:"max_workers"`
		TickInterval string `json:"tick_interval"`
	}

	timer struct {
		time     time.Time
		interval time.Duration
		cb       func()
	}

	Service struct {
		timers slab.Slab[timer]
		remove stack.Stack[slab.Key]
		pool   weave.WorkerPool
	}
)

func New(config *Config) func() service.Service {
	return func() service.Service {
		return &Service{
			timers: slab.Init[timer](config.MaxTimers),
			remove: stack.Init[slab.Key](config.MaxTimers),
			pool:   *weave.NewWorkerPool(config.MaxWorkers),
		}
	}
}

func (s *Service) Start(services service.Services, shutdown func()) error {
	return nil
}

func (s *Service) Stop() {
	s.pool.Dispose()
	// TODO: wait for all delayed tasks
}

func (s *Service) Update(_ float64) {
	{
		iter := s.timers.Iter()
		for iter.Next() {
			now := time.Now()
			v := iter.Value()
			if now.Sub(v.time) >= v.interval {
				v.cb()
				v.time = now
			}
		}
	}

	{
		iter := s.remove.Iter()
		for iter.Next() {
			v := iter.Value()
			s.timers.Remove(v)
		}
		s.remove.Clear()
	}

	s.pool.Poll()
}

func (s *Service) Every(d time.Duration, task func()) func() {
	t := timer{interval: d, time: time.Now(), cb: task}
	k := s.timers.Insert(t)
	return func() { s.remove.Push(k) }
}

func (s *Service) Delay(d time.Duration, task func()) {
	e := s.timers.VacantEntry()
	k := e.Key()
	t := timer{interval: d, time: time.Now(), cb: func() {
		task()
		s.remove.Push(k)
	}}
	e.Insert(t)
}

func (s *Service) RepeatEvery(times int, d time.Duration, task func()) {
	e := s.timers.VacantEntry()
	k := e.Key()
	n := 0
	t := timer{interval: d, time: time.Now(), cb: func() {
		n += 1
		task()
		if n == times {
			s.remove.Push(k)
		}
	}}
	e.Insert(t)
}

func (s *Service) Dispatch(task weave.Task) {
	s.pool.Dispatch(task)
}
