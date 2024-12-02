package scene

import (
	"github.com/carabalonepaulo/origin/shared/service"
)

type Manager struct {
	current  Scene
	services service.Services
}

type SceneManager interface {
	ChangeTo(scene Scene)
}

func NewManager(services service.Services, firstScene Scene) *Manager {
	manager := &Manager{}
	manager.services = services
	manager.ChangeTo(firstScene)
	return manager
}

func (m *Manager) ChangeTo(scene Scene) {
	m.Unload()
	m.current = scene
	m.Load()
}

func (m *Manager) Load() {
	if m.current != nil {
		m.current.Load(m.services, m)
	}
}

func (m *Manager) Unload() {
	if m.current == nil {
		return
	}
	m.current.Unload()
	m.current = nil
}

func (m *Manager) Draw() {
	if m.current != nil {
		m.current.Draw()
	}
}
