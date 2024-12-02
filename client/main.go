package main

import (
	"log"

	"github.com/carabalonepaulo/origin/client/client"
	"github.com/carabalonepaulo/origin/client/config"
	"github.com/carabalonepaulo/origin/client/scenes"
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
		client.New("127.0.0.1", 5051),
		scenes.New,
	)
	if err != nil {
		log.Println(err)
	}
}
