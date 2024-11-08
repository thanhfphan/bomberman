package game

import (
	"fmt"

	"thanhfphan.com/bomberman/assets"
	"thanhfphan.com/bomberman/src/engine/audio"
	"thanhfphan.com/bomberman/src/engine/spritesheet"
)

func (g *Game) Setup() error {
	player, err := g.entityManager.CreateEntity()
	if err != nil {
		return fmt.Errorf("could not create player entity: %v", err)
	}
	g.player = player
	g.player.Body.Position.X = 300
	g.player.Body.Position.Y = 200

	// ********** Add player animations **********
	// Walk right
	ssPlayerWalkRight, err := spritesheet.NewSpriteSheet(assets.PlayerWalkRight, 32, 32)
	if err != nil {
		return fmt.Errorf("could not create player sprite sheet: %v", err)
	}
	walkRightID := g.animationManager.CreateDefinition(ssPlayerWalkRight, 0.1, 0, []uint8{0, 1, 2, 3}, 4)
	g.animationWalkRightID = g.animationManager.CreateAnimation(walkRightID, true)
	player.AnimationID = g.animationWalkRightID
	// Idle
	ssPlayerIdle, err := spritesheet.NewSpriteSheet(assets.PlayerIdleFront, 32, 32)
	if err != nil {
		return fmt.Errorf("could not create player idle sprite sheet: %v", err)
	}
	idleID := g.animationManager.CreateDefinition(ssPlayerIdle, 0, 0, []uint8{0}, 1)
	g.animationIdleID = g.animationManager.CreateAnimation(idleID, false)
	// Walk up
	ssPlayerWalkBack, err := spritesheet.NewSpriteSheet(assets.PlayerWalkBack, 32, 32)
	if err != nil {
		return fmt.Errorf("could not create player walk up sprite sheet: %v", err)
	}
	walkBackID := g.animationManager.CreateDefinition(ssPlayerWalkBack, 0.1, 0, []uint8{0, 1, 2, 3}, 4)
	g.animationWalkBackID = g.animationManager.CreateAnimation(walkBackID, true)
	// Walk down
	ssPlayerWalkFront, err := spritesheet.NewSpriteSheet(assets.PlayerWalkFront, 32, 32)
	if err != nil {
		return fmt.Errorf("could not create player walk down sprite sheet: %v", err)
	}
	walkFrontID := g.animationManager.CreateDefinition(ssPlayerWalkFront, 0.1, 0, []uint8{0, 1, 2, 3}, 4)
	g.animationWalkFrontID = g.animationManager.CreateAnimation(walkFrontID, true)

	// ********** Audio **********
	g.backgroundMusic, err = audio.LoadWAV(assets.BackGroundMusic, true)
	if err != nil {
		return fmt.Errorf("could not load background music: %v", err)
	}
	g.playBombSound, err = audio.LoadWAV(assets.BomSetSound, false)
	if err != nil {
		return fmt.Errorf("could not load bomb set sound: %v", err)
	}

	return nil
}
