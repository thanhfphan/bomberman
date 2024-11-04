package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"thanhfphan.com/bomberman/src/engine"
)

type Player struct {
	x, y  float64
	speed float64 // units per second
}

func NewPlayer() *Player {
	return &Player{
		x:     400,
		y:     400,
		speed: 500,
	}
}

type Game struct {
	render *engine.RenderState
	time   *engine.TimeState
	input  *engine.InputSate

	player *Player
}

func (p *Player) Update(delta float64, input *engine.InputSate) {
	if input.Left == engine.KeyStatePressed || input.Left == engine.KeyStateHeld {
		p.x -= p.speed * delta
	}
	if input.Right == engine.KeyStatePressed || input.Right == engine.KeyStateHeld {
		p.x += p.speed * delta
	}
	if input.Up == engine.KeyStatePressed || input.Up == engine.KeyStateHeld {
		p.y -= p.speed * delta
	}
	if input.Down == engine.KeyStatePressed || input.Down == engine.KeyStateHeld {
		p.y += p.speed * delta
	}
}

func NewGame(w, h int) *Game {
	return &Game{
		render: engine.NewRenderState(w, h),
		time:   engine.NewTimeState(),
		input:  engine.NewInputState(),
		player: NewPlayer(),
	}
}

func (g *Game) LoadConfig(file string) error {
	if err := engine.LoadConfig(file); err != nil {
		return fmt.Errorf("could not load config file: %v", err)
	}
	return nil
}

func (g *Game) Update() error {
	if g.input.Escape == engine.KeyStatePressed {
		return ebiten.Termination
	}

	g.time.Update()
	g.input.Update()
	g.player.Update(g.time.Delta(), g.input)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.render.Begin(screen)

	g.render.RenderQuad(screen, float32(g.player.x), float32(g.player.y), 100, 100, color.White)

	g.render.End(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	width, height := 640, 480
	game := NewGame(width, height)
	if err := game.LoadConfig("config.ini"); err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	ebiten.SetWindowTitle("Bomberman")
	ebiten.SetWindowSize(width, height)

	if err := ebiten.RunGame(game); err != nil && err != ebiten.Termination {
		log.Fatal(err)
	}
}
