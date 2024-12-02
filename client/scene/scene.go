package scene

type Scene interface {
	Load(manager SceneManager)
	Unload()
	Update(dt float64)
	Draw()
}
