package scene

type Manager struct {
	current Scene
}

func NewManager(firstScene Scene) *Manager {
	manager := &Manager{}
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
		m.current.Load()
	}
}

func (m *Manager) Unload() {
	if m.current == nil {
		return
	}
	m.current.Unload()
	m.current = nil
}

func (m *Manager) Update(dt float64) {
	if m.current != nil {
		m.current.Update(dt)
	}
}

func (m *Manager) Draw() {
	if m.current != nil {
		m.current.Draw()
	}
}
