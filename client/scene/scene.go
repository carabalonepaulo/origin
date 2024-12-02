package scene

type Scene interface {
	Load()
	Unload()
	Update(dt float64)
	Draw()
}
