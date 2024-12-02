package boot

import (
	"fmt"

	"github.com/carabalonepaulo/origin/client/scene"
	"github.com/carabalonepaulo/origin/shared/service"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Scene struct{}

func New() *Scene {
	return &Scene{}
}

func (s *Scene) Load(services service.Services, manager scene.SceneManager) {}

func (s *Scene) Unload() {}

func (s *Scene) Draw() {
	rl.DrawText(fmt.Sprintf("dt: %.2f", rl.GetFrameTime()), 10, 10, 30, rl.Black)
}
