package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"thanhfphan.com/bomberman/src/engine"
)

type Game struct {
	renderState *engine.RenderState
}

func NewGame(w, h int) *Game {
	return &Game{
		renderState: engine.NewRenderState(w, h),
	}
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return ebiten.Termination
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.renderState.Begin(screen)

	screenW, screenH := g.renderState.ScreenWidth(), g.renderState.ScreenHeight()
	g.renderState.RenderQuad(screen, float32(screenW/2), float32(screenH/2), 100, 100, color.White)

	g.renderState.End(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	width, height := 640, 480
	game := NewGame(width, height)

	ebiten.SetWindowTitle("Bomberman")
	ebiten.SetWindowSize(width, height)

	if err := ebiten.RunGame(game); err != nil && err != ebiten.Termination {
		log.Fatal(err)
	}
}
