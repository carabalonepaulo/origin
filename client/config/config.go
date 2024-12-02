package config

import (
	"github.com/carabalonepaulo/origin/client/client"
	"github.com/carabalonepaulo/origin/shared/services/scheduler"
)

type (
	Config struct {
		Scheduler scheduler.Config `json:"scheduler"`
		Client    client.Config    `json:"config"`
	}
)
