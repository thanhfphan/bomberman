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

	initAssets()
	global.assetKeeper.Load()

	g.player = NewPlayer(
		engine.PlayerSpeed,
		math.Vec2{X: 300, Y: 200},
	)

	return nil
}
