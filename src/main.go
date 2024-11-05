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

	entityManager *engine.EntityManager
	player        *engine.Entity
}

func NewGame(w, h int) *Game {
	return &Game{
		render:        engine.NewRenderState(w, h),
		time:          engine.NewTimeState(),
		input:         engine.NewInputState(),
		entityManager: engine.NewEntityManager(),
	}
}

func (g *Game) Setup() error {
	player, err := g.entityManager.CreateEntity()
	if err != nil {
		return fmt.Errorf("could not create player entity: %v", err)
	}
	g.player = player

	g.player.Body.AABB.Position.X = 300
	g.player.Body.AABB.Position.Y = 200

	return nil
}

func (g *Game) LoadConfig(file string) error {
	if err := engine.LoadConfig(file); err != nil {
		return fmt.Errorf("could not load config file: %v", err)
	}
	return nil
}

func (g *Game) inputHandle(body *engine.Body) {
	if g.input.Left == engine.KeyStatePressed || g.input.Left == engine.KeyStateHeld {
		fmt.Println(body.AABB.Position.X, body.AABB.Position.Y)
		body.AABB.Position.X -= 500 * g.time.Delta
	}
	if g.input.Right == engine.KeyStatePressed || g.input.Right == engine.KeyStateHeld {
		body.AABB.Position.X += 500 * g.time.Delta
	}
	if g.input.Up == engine.KeyStatePressed || g.input.Up == engine.KeyStateHeld {
		body.AABB.Position.Y -= 500 * g.time.Delta
	}
	if g.input.Down == engine.KeyStatePressed || g.input.Down == engine.KeyStateHeld {
		body.AABB.Position.Y += 500 * g.time.Delta
	}
}

func (g *Game) Update() error {
	if g.input.Escape == engine.KeyStatePressed {
		return ebiten.Termination
	}

	g.time.Update()
	g.input.Update()
	g.inputHandle(g.player.Body)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.render.Begin(screen)

	g.render.RenderQuad(screen, float32(g.player.Body.AABB.Position.X), float32(g.player.Body.AABB.Position.Y), 100, 100, color.White)

	g.render.End(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 544, 480
}

func main() {
	width, height := 1088, 960
	game := NewGame(width, height)
	if err := game.Setup(); err != nil {
		log.Fatalf("could not setup game: %v", err)
	}
	if err := game.LoadConfig("config.ini"); err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	ebiten.SetWindowTitle("Bomberman")
	ebiten.SetWindowSize(width, height)

	if err := ebiten.RunGame(game); err != nil && err != ebiten.Termination {
		log.Fatal(err)
	}
}
