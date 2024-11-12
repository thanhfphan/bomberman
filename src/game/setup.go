package game

import (
	"fmt"

	"thanhfphan.com/bomberman/src/engine"
	"thanhfphan.com/bomberman/src/engine/math"
)

func (g *Game) Setup(configFile string) error {
	if err := engine.LoadConfig(configFile); err != nil {
		return fmt.Errorf("could not load config file: %w", err)
	}

	if err := global.assetKeeper.Load(); err != nil {
		return fmt.Errorf("could not load assets: %w", err)
	}

	g.player = NewPlayer(
		engine.PlayerSpeed,
		math.Vec2{X: 300, Y: 200},
	)

	// Create a new bat for testing
	NewBat(math.Vec2{X: 100, Y: 100})

	return nil
}
