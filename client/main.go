package main

import (
	"github.com/carabalonepaulo/origin/client/client"
	sc "github.com/carabalonepaulo/origin/client/config"
	"github.com/carabalonepaulo/origin/client/scenes"
	"github.com/carabalonepaulo/origin/shared/config"
	"github.com/carabalonepaulo/origin/shared/services"
	"github.com/carabalonepaulo/origin/shared/services/scheduler"
	"github.com/carabalonepaulo/origin/shared/sys"
)

func main() {
	config := config.LoadDefault[sc.Config]()
	services.Run(
		sys.New,
		scheduler.New(&config.Scheduler),
		client.New(&config.Client),
		scenes.New,
	)
}
