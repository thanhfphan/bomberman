package engine

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type RenderState struct {
	ScreenWidth  int
	ScreenHeight int
}

func NewRenderState(w, h int) *RenderState {
	return &RenderState{w, h}
}

func (rs *RenderState) Begin(screen *ebiten.Image) {
	screen.Clear()
}

func (rs *RenderState) End(screen *ebiten.Image) {
}

func (rs *RenderState) RenderQuad(screen *ebiten.Image, x, y, w, h float32, c color.Color) {
	vector.DrawFilledRect(screen, x-w/2, y-h/2, w, h, c, false)
}
