package engine

import "github.com/hajimehoshi/ebiten/v2"

type Entity interface {
	GetID() int
	Update(deltaTime float64)
	IsActive() bool
	Render(screen *ebiten.Image)
	Destroy()
}
