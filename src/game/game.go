package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"thanhfphan.com/bomberman/src/engine"
	"thanhfphan.com/bomberman/src/engine/animation"
	"thanhfphan.com/bomberman/src/engine/audio"
)

type GlobalState struct {
	time        *engine.TimeState
	input       *engine.InputSate
	animation   *animation.Manager
	entity      *EntityManager
	assetKeeper *AssetKeeper
}

var global GlobalState

func init() {
	global.time = engine.NewTimeState()
	global.input = engine.NewInputState()
	global.entity = NewEntityManager()
	global.animation = animation.NewManager()
	global.assetKeeper = NewAssetKeeper()
}

const (
	TileSize      = 32
	LogicalWidth  = 544
	LogicalHeight = 480
	ScaleFactor   = 2
	WindowWidth   = LogicalWidth * ScaleFactor
	WindowHeight  = LogicalHeight * ScaleFactor
	GridWidth     = LogicalWidth / TileSize  // 17 tiles
	GridHeight    = LogicalHeight / TileSize // 15 tiles
	//A good SnapThreshold typically ranges between 2% to 10% of the TileSize
	// Medium Threshold (5%) 0.05×32=1.6≈2 pixels
	BaseSnapThreshold      = float64(2) // 5% of TileSize
	HighSpeedSnapThreshold = float64(4) // 10-15% of TileSize
	HighSpeedThreshold     = float64(8) // 25% of TileSize
)

type Game struct {
	grid   *Grid
	player *Player
	render *engine.RenderState
}

func New() *Game {
	return &Game{
		grid:   NewGrid(GridWidth, GridHeight),
		render: engine.NewRenderState(WindowWidth, WindowHeight),
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

	g.grid.Render(screen)

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
	return LogicalWidth, LogicalHeight
}
