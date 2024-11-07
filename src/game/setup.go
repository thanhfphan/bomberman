package game

import (
	"fmt"

	"thanhfphan.com/bomberman/assets"
	"thanhfphan.com/bomberman/src/engine/spritesheet"
)

func (g *Game) Setup() error {
	player, err := g.entityManager.CreateEntity()
	if err != nil {
		return fmt.Errorf("could not create player entity: %v", err)
	}
	g.player = player
	g.player.Body.AABB.Position.X = 300
	g.player.Body.AABB.Position.Y = 200

	// ********** Add player animations **********
	// Walk right
	ssPlayerWalkRight, err := spritesheet.NewSpriteSheet(assets.PlayerWalkRight, 32, 32)
	if err != nil {
		return fmt.Errorf("could not create player sprite sheet: %v", err)
	}
	walkRightID := g.animationManager.CreateDefinition(ssPlayerWalkRight, []float32{0.1, 0.1, 0.1, 0.1}, []uint8{0, 0, 0, 0}, []uint8{0, 1, 2, 3}, 4)
	g.animationWalkRightID, err = g.animationManager.CreateAnimation(walkRightID, true)
	if err != nil {
		return fmt.Errorf("could not add player walk animation: %v", err)
	}
	player.AnimationID = g.animationWalkRightID
	// Idle
	ssPlayerIdle, err := spritesheet.NewSpriteSheet(assets.PlayerIdleFront, 32, 32)
	if err != nil {
		return fmt.Errorf("could not create player idle sprite sheet: %v", err)
	}
	idleID := g.animationManager.CreateDefinition(ssPlayerIdle, []float32{0}, []uint8{0}, []uint8{0}, 1)
	g.animationIdleID, err = g.animationManager.CreateAnimation(idleID, false)
	if err != nil {
		return fmt.Errorf("could not add player idle animation: %v", err)
	}
	// Walk up
	ssPlayerWalkBack, err := spritesheet.NewSpriteSheet(assets.PlayerWalkBack, 32, 32)
	if err != nil {
		return fmt.Errorf("could not create player walk up sprite sheet: %v", err)
	}
	walkBackID := g.animationManager.CreateDefinition(ssPlayerWalkBack, []float32{0.1, 0.1, 0.1, 0.1}, []uint8{0, 0, 0, 0}, []uint8{0, 1, 2, 3}, 4)
	g.animationWalkBackID, err = g.animationManager.CreateAnimation(walkBackID, true)
	if err != nil {
		return fmt.Errorf("could not add player walk up animation: %v", err)
	}
	// Walk down
	ssPlayerWalkFront, err := spritesheet.NewSpriteSheet(assets.PlayerWalkFront, 32, 32)
	if err != nil {
		return fmt.Errorf("could not create player walk down sprite sheet: %v", err)
	}
	walkFrontID := g.animationManager.CreateDefinition(ssPlayerWalkFront, []float32{0.1, 0.1, 0.1, 0.1}, []uint8{0, 0, 0, 0}, []uint8{0, 1, 2, 3}, 4)
	g.animationWalkFrontID, err = g.animationManager.CreateAnimation(walkFrontID, true)
	if err != nil {
		return fmt.Errorf("could not add player walk down animation: %v", err)
	}

	return nil
}
