package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"thanhfphan.com/bomberman/src/engine"
	"thanhfphan.com/bomberman/src/engine/animation"
	"thanhfphan.com/bomberman/src/engine/audio"
)

type GlobalState struct {
	time             *engine.TimeState
	input            *engine.InputSate
	entityManager    *engine.EntityManager
	animationManager *animation.Manager
}

var gs GlobalState

func init() {
	gs.time = engine.NewTimeState()
	gs.input = engine.NewInputState()
	gs.entityManager = engine.NewEntityManager()
	gs.animationManager = animation.NewManager()
}

type Game struct {
	render *engine.RenderState

	player *Player

	animationIdleID      int
	animationWalkRightID int
	animationWalkBackID  int
	animationWalkFrontID int

	backgroundMusic *audio.Player
	playBombSound   *audio.Player
}

func New(w, h int) *Game {
	return &Game{
		render: engine.NewRenderState(w, h),
	}
}

func (g *Game) LoadConfig(file string) error {
	if err := engine.LoadConfig(file); err != nil {
		return fmt.Errorf("could not load config file: %v", err)
	}
	return nil
}

func (g *Game) handlePlayer() {
	g.player.Body.Velocity.X = 0
	g.player.Body.Velocity.Y = 0
	if gs.input.Left == engine.KeyStatePressed || gs.input.Left == engine.KeyStateHeld {
		g.player.Body.Velocity.X -= engine.PlayerSpeed
	}
	if gs.input.Right == engine.KeyStatePressed || gs.input.Right == engine.KeyStateHeld {
		g.player.Body.Velocity.X += engine.PlayerSpeed
	}
	if gs.input.Up == engine.KeyStatePressed || gs.input.Up == engine.KeyStateHeld {
		g.player.Body.Velocity.Y -= engine.PlayerSpeed
	}
	if gs.input.Down == engine.KeyStatePressed || gs.input.Down == engine.KeyStateHeld {
		g.player.Body.Velocity.Y += engine.PlayerSpeed
	}

	if gs.input.PlaceBomb == engine.KeyStatePressed {
		audio.Play(g.playBombSound)
	}
}

func (g *Game) Update() error {
	if gs.input.Escape == engine.KeyStatePressed {
		return ebiten.Termination
	}

	playerv := g.player.Body.Velocity
	if playerv.Y < 0 {
		g.player.AnimationID = g.animationWalkBackID
	} else if playerv.Y > 0 {
		g.player.AnimationID = g.animationWalkFrontID
	} else if playerv.X != 0 {
		g.player.AnimationID = g.animationWalkRightID // Walk left will be flipped
	} else {
		g.player.AnimationID = g.animationIdleID
	}

	gs.time.Update()
	gs.input.Update()

	g.handlePlayer()

	gs.animationManager.Update(gs.time.Delta)
	gs.entityManager.Update(gs.time.Delta)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.render.Begin(screen)

	for i := 0; i < gs.entityManager.Size(); i++ {
		entity, err := gs.entityManager.GetEntity(i)
		if err != nil {
			continue
		}

		// Draw sprite
		if entity.AnimationID >= 0 {
			animation := gs.animationManager.GetAnimation(entity.AnimationID)
			if animation.IsActive {
				if entity.Body.Velocity.Y == 0 {
					if entity.Body.Velocity.X < 0 {
						animation.IsFlipped = true
					} else if entity.Body.Velocity.X > 0 {
						animation.IsFlipped = false
					}
				}
				aframe := animation.Definition.Frames[animation.CurrentFrameIndex]
				animation.Definition.SpriteSheet.DrawFrame(screen, float64(aframe.Row), float64(aframe.Column), entity.Body.Position, animation.IsFlipped)
			}
		}

	}

	g.render.End(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 544, 480
}
