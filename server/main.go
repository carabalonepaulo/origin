package main

import (
	"log"

	"github.com/carabalonepaulo/origin/server/config"
	"github.com/carabalonepaulo/origin/server/listener"
	"github.com/carabalonepaulo/origin/shared/services"
	"github.com/carabalonepaulo/origin/shared/services/scheduler"
	"github.com/carabalonepaulo/origin/shared/sys"
)

func main() {
	config, err := config.Load(config.PathOrDefault("./config.json"))
	if err != nil {
		log.Println(err)
		return
	}

	err = services.Run(
		sys.New,
		scheduler.New(&config.Scheduler),
		listener.New(&config.Listener),
	)
	if err != nil {
		log.Println(err)
	}
}
