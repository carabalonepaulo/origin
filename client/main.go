package main

import (
	"log"

	"github.com/carabalonepaulo/origin/client/client"
	"github.com/carabalonepaulo/origin/client/config"
	"github.com/carabalonepaulo/origin/client/scenes"
	sharedConfig "github.com/carabalonepaulo/origin/shared/config"
	"github.com/carabalonepaulo/origin/shared/services"
	"github.com/carabalonepaulo/origin/shared/services/scheduler"
	"github.com/carabalonepaulo/origin/shared/sys"
)

func main() {
	config, err := sharedConfig.Load[config.Config](sharedConfig.PathOrDefault("./config.json"))
	if err != nil {
		log.Println(err)
		return
	}

	err = services.Run(
		sys.New,
		scheduler.New(&config.Scheduler),
		client.New(&config.Client),
		scenes.New,
	)
	if err != nil {
		log.Println(err)
	}
}
