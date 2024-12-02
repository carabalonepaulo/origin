package scenes

import (
	"github.com/carabalonepaulo/origin/client/scene"
	"github.com/carabalonepaulo/origin/client/scenes/boot"
	"github.com/carabalonepaulo/origin/shared/service"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var ratio = rl.Vector2{X: 16, Y: 9}
var width int32 = 1024
var height int32 = int32(float32(width) / ratio.X * ratio.Y)

const title = "Origin"
const fps = 60

type Service struct {
	manager  *scene.Manager
	shutdown func()
}

func New() service.Service {
	return &Service{}
}

func (s *Service) Start(services service.Services, shutdown func()) error {
	s.manager = scene.NewManager(services, boot.New())
	s.shutdown = shutdown

	rl.SetTraceLogLevel(rl.LogNone)
	rl.InitWindow(width, height, title)
	rl.SetTargetFPS(fps)

	return nil
}

func (s *Service) Stop() {
	s.manager.Unload()
	rl.CloseWindow()
}

func (s *Service) Update(dt float64) {
	if rl.WindowShouldClose() {
		s.shutdown()
		return
	}

	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)
	s.manager.Draw()
	rl.EndDrawing()
}
