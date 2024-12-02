package config

import (
	"github.com/carabalonepaulo/origin/server/listener"
	"github.com/carabalonepaulo/origin/shared/services/scheduler"
)

type (
	Config struct {
		Scheduler scheduler.Config `json:"scheduler"`
		Listener  listener.Config  `json:"listener"`
	}
)
