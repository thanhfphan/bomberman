package game

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"thanhfphan.com/bomberman/assets"
	"thanhfphan.com/bomberman/src/engine"
	"thanhfphan.com/bomberman/src/engine/animation"
	"thanhfphan.com/bomberman/src/engine/audio"
	"thanhfphan.com/bomberman/src/engine/spritesheet"
)

type GlobalState struct {
	time             *engine.TimeState
	input            *engine.InputSate
	animationManager *animation.Manager
	entityManager    *EntityManager
}

var (
	backgroundMusic  *audio.Player
	soundBombSet     *audio.Player
	soundBombExplode *audio.Player

	animationIdleID      int
	animationWalkRightID int
	animationWalkBackID  int
	animationWalkFrontID int

	bombDefID          int
	bombExplosionDefID int
)

var gs GlobalState

func init() {
	gs.time = engine.NewTimeState()
	gs.input = engine.NewInputState()
	gs.entityManager = NewEntityManager()
	gs.animationManager = animation.NewManager()

	// ********** Audio **********
	var err error
	backgroundMusic, err = audio.LoadWAV(assets.BackGroundMusic, true)
	if err != nil {
		panic(fmt.Errorf("could not load background music: %v", err))
	}
	soundBombSet, err = audio.LoadWAV(assets.BomSetSound, false)
	if err != nil {
		panic(fmt.Errorf("could not load bomb set sound: %v", err))
	}
	soundBombExplode, err = audio.LoadWAV(assets.BomExplodeSound, false)
	if err != nil {
		panic(fmt.Errorf("could not load bomb explode sound: %v", err))
	}

	// ********** Add player animations **********
	ssPlayerWalkRight, err := spritesheet.NewSpriteSheet(assets.PlayerWalkRight, 32, 32)
	if err != nil {
		panic(fmt.Errorf("could not create player sprite sheet: %v", err))
	}
	// Walk right
	walkRightID := gs.animationManager.CreateDefinition(
		ssPlayerWalkRight,
		0.1,
		[]uint8{0, 0, 0, 0},
		[]uint8{0, 1, 2, 3},
		4,
	)
	animationWalkRightID = gs.animationManager.CreateAnimation(walkRightID, true)
	// Idle
	ssPlayerIdle, err := spritesheet.NewSpriteSheet(assets.PlayerIdleFront, 32, 32)
	if err != nil {
		panic(fmt.Errorf("could not create player idle sprite sheet: %v", err))
	}
	idleID := gs.animationManager.CreateDefinition(
		ssPlayerIdle,
		0.1,
		[]uint8{0, 0, 0, 0},
		[]uint8{0, 1, 2, 3},
		4,
	)
	animationIdleID = gs.animationManager.CreateAnimation(idleID, false)
	// Walk up
	ssPlayerWalkBack, err := spritesheet.NewSpriteSheet(assets.PlayerWalkBack, 32, 32)
	if err != nil {
		panic(fmt.Errorf("could not create player walk up sprite sheet: %v", err))
	}
	walkBackID := gs.animationManager.CreateDefinition(
		ssPlayerWalkBack,
		0.1,
		[]uint8{0, 0, 0, 0},
		[]uint8{0, 1, 2, 3},
		4,
	)
	animationWalkBackID = gs.animationManager.CreateAnimation(walkBackID, true)
	// Walk down
	ssPlayerWalkFront, err := spritesheet.NewSpriteSheet(assets.PlayerWalkFront, 32, 32)
	if err != nil {
		panic(fmt.Errorf("could not create player walk down sprite sheet: %v", err))
	}
	walkFrontID := gs.animationManager.CreateDefinition(
		ssPlayerWalkFront,
		0.1,
		[]uint8{0, 0, 0, 0},
		[]uint8{0, 1, 2, 3},
		4,
	)
	animationWalkFrontID = gs.animationManager.CreateAnimation(walkFrontID, true)

	// ********** Add bomb animations **********
	bombSprite, err := spritesheet.NewSpriteSheet(assets.BombSprite, 32, 32)
	if err != nil {
		panic(fmt.Errorf("could not create bomb sprite sheet: %v", err))
	}
	bombDefID = gs.animationManager.CreateDefinition(
		bombSprite,
		0.1,
		[]uint8{0, 0, 0, 1, 1, 1, 2, 2, 2},
		[]uint8{0, 1, 2, 0, 1, 2, 0, 1, 2},
		9,
	)
	bombExplosion, err := spritesheet.NewSpriteSheet(assets.BombExplosion, 32, 32)
	if err != nil {
		panic(fmt.Errorf("could not create bomb explosion sprite sheet: %v", err))
	}
	bombExplosionDefID = gs.animationManager.CreateDefinition(
		bombExplosion,
		0.1,
		[]uint8{0, 0, 0, 1, 1, 1},
		[]uint8{0, 1, 2, 0, 1, 2},
		6,
	)
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

func (g *Game) createBomb() {
	bomb := &Bomb{
		Countdown:          time.Duration(3) * time.Second,
		PlacedAt:           time.Now(),
		Position:           g.player.Position,
		AnimationID:        gs.animationManager.CreateAnimation(bombDefID, true),
		AnimationExploseID: gs.animationManager.CreateAnimation(bombExplosionDefID, false),
	}
	bomb.ID = gs.entityManager.Create(bomb)
}

func (g *Game) Update() error {
	if gs.input.Escape == engine.KeyStatePressed {
		return ebiten.Termination
	}

	gs.time.Update()
	gs.input.Update()

	if gs.input.PlaceBomb == engine.KeyStatePressed {
		audio.Play(soundBombSet)
		g.createBomb()
	}

	gs.animationManager.Update(gs.time.Delta)
	for i := 0; i < gs.entityManager.Size(); i++ {
		entity, err := gs.entityManager.GetEntity(i)
		if err != nil {
			continue
		}
		if !entity.IsActive() {
			continue
		}

		entity.Update(gs.time.Delta)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.render.Begin(screen)

	for i := 0; i < gs.entityManager.Size(); i++ {
		entity, err := gs.entityManager.GetEntity(i)
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
