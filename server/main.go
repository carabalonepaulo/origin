package main

import (
	c "github.com/carabalonepaulo/origin/server/config"
	"github.com/carabalonepaulo/origin/server/listener"
	"github.com/carabalonepaulo/origin/shared/config"
	"github.com/carabalonepaulo/origin/shared/services"
	"github.com/carabalonepaulo/origin/shared/services/scheduler"
	"github.com/carabalonepaulo/origin/shared/sys"
)

func main() {
	config := config.LoadDefault[c.Config]()
	services.Run(
		sys.New,
		scheduler.New(&config.Scheduler),
		listener.New(&config.Listener),
	)
}
