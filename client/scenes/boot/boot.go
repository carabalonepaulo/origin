package boot

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Scene struct{}

func New() *Scene {
	return &Scene{}
}

func (s *Scene) Load() {}

func (s *Scene) Unload() {}

func (s *Scene) Update(_ float64) {}

func (s *Scene) Draw() {
	rl.DrawText(fmt.Sprintf("dt: %.2f", rl.GetFrameTime()), 10, 10, 30, rl.Black)
}