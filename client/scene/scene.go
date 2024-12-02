package scene

import "github.com/carabalonepaulo/origin/shared/service"

type Scene interface {
	Load(services service.Services, manager SceneManager)
	Unload()
	Draw()
}
