package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"thanhfphan.com/bomberman/src/engine"
)

type Game struct {
	render *engine.RenderState
	time   *engine.TimeState
	input  *engine.InputSate

	entityManager *engine.EntityManager
	player        *engine.Entity
}

func New(w, h int) *Game {
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
