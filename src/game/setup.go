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

	player := &Player{
		Speed: engine.PlayerSpeed,
		Position: math.Vec2{
			X: 300,
			Y: 200,
		},
	}
	player.ID = gs.entityManager.Create(player)
	g.player = player

	return nil
}
