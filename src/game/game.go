package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"thanhfphan.com/bomberman/src/engine"
	"thanhfphan.com/bomberman/src/engine/animation"
	"thanhfphan.com/bomberman/src/engine/audio"
)

type GlobalState struct {
	time      *engine.TimeState
	input     *engine.InputSate
	animation *animation.Manager
	entity    *EntityManager
}

var global GlobalState

func init() {
	global.time = engine.NewTimeState()
	global.input = engine.NewInputState()
	global.entity = NewEntityManager()
	global.animation = animation.NewManager()

}

type Game struct {
	player *Player
	render *engine.RenderState
}

func New(w, h int) *Game {
	return &Game{
		render: engine.NewRenderState(w, h),
	}
}

func (g *Game) Update() error {
	if global.input.Escape == engine.KeyStatePressed {
		return ebiten.Termination
	}

	global.time.Update()
	global.input.Update()

	if global.input.PlaceBomb == engine.KeyStatePressed {
		audio.Play(soundBombSet)
		NewBomb(g.player.Position)
	}

	global.animation.Update(global.time.Delta)
	for i := 0; i < global.entity.Size(); i++ {
		entity, err := global.entity.GetEntity(i)
		if err != nil {
			continue
		}
		if !entity.IsActive() {
			continue
		}

		entity.Update(global.time.Delta)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.render.Begin(screen)

	for i := 0; i < global.entity.Size(); i++ {
		entity, err := global.entity.GetEntity(i)
		if err != nil {
			continue
		}

		entity.Render(screen)
	}

	g.render.End(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 544, 480
}
