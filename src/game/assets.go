package game

import (
	"fmt"

	"thanhfphan.com/bomberman/assets"
	"thanhfphan.com/bomberman/src/engine/audio"
	"thanhfphan.com/bomberman/src/engine/spritesheet"
)

var (
	backgroundMusic  *audio.Player
	soundBombSet     *audio.Player
	soundBombExplode *audio.Player

	animationIdleRightID int
	animationIdleBackID  int
	animationIdleFrontID int

	animationWalkRightID int
	animationWalkBackID  int
	animationWalkFrontID int

	bombDefID          int
	bombExplosionDefID int
)

func initAssets() {
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
	walkRightID := global.animation.CreateDefinition(
		ssPlayerWalkRight,
		0.1,
		[]uint8{0, 0, 0, 0},
		[]uint8{0, 1, 2, 3},
		4,
	)
	animationWalkRightID = global.animation.CreateAnimation(walkRightID, true)
	// Idle front
	idleFront, err := spritesheet.NewSpriteSheet(assets.PlayerIdleFront, 32, 32)
	if err != nil {
		panic(fmt.Errorf("could not create player idle sprite sheet: %v", err))
	}
	animationIdleFrontID = global.animation.CreateAnimation(
		global.animation.CreateDefinition(
			idleFront,
			0.3,
			[]uint8{0, 0, 0, 0},
			[]uint8{0, 1, 2, 3},
			4,
		),
		true)
	// Idle right
	idleRight, err := spritesheet.NewSpriteSheet(assets.PlayerIdleRight, 32, 32)
	if err != nil {
		panic(fmt.Errorf("could not create player idle sprite sheet: %v", err))
	}
	animationIdleRightID = global.animation.CreateAnimation(
		global.animation.CreateDefinition(
			idleRight,
			0.3,
			[]uint8{0, 0, 0, 0},
			[]uint8{0, 1, 2, 3},
			4,
		),
		true)
	idleBack, err := spritesheet.NewSpriteSheet(assets.PlayerIdleBack, 32, 32)
	if err != nil {
		panic(fmt.Errorf("could not create player idle sprite sheet: %v", err))
	}
	animationIdleBackID = global.animation.CreateAnimation(
		global.animation.CreateDefinition(
			idleBack,
			0.3,
			[]uint8{0, 0, 0, 0},
			[]uint8{0, 1, 2, 3},
			4,
		),
		true)
	// Walk up
	ssPlayerWalkBack, err := spritesheet.NewSpriteSheet(assets.PlayerWalkBack, 32, 32)
	if err != nil {
		panic(fmt.Errorf("could not create player walk up sprite sheet: %v", err))
	}
	walkBackID := global.animation.CreateDefinition(
		ssPlayerWalkBack,
		0.1,
		[]uint8{0, 0, 0, 0},
		[]uint8{0, 1, 2, 3},
		4,
	)
	animationWalkBackID = global.animation.CreateAnimation(walkBackID, true)
	// Walk down
	ssPlayerWalkFront, err := spritesheet.NewSpriteSheet(assets.PlayerWalkFront, 32, 32)
	if err != nil {
		panic(fmt.Errorf("could not create player walk down sprite sheet: %v", err))
	}
	walkFrontID := global.animation.CreateDefinition(
		ssPlayerWalkFront,
		0.1,
		[]uint8{0, 0, 0, 0},
		[]uint8{0, 1, 2, 3},
		4,
	)
	animationWalkFrontID = global.animation.CreateAnimation(walkFrontID, true)

	// ********** Add bomb animations **********
	bombSprite, err := spritesheet.NewSpriteSheet(assets.BombSprite, 32, 32)
	if err != nil {
		panic(fmt.Errorf("could not create bomb sprite sheet: %v", err))
	}
	bombDefID = global.animation.CreateDefinition(
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
	bombExplosionDefID = global.animation.CreateDefinition(
		bombExplosion,
		0.1,
		[]uint8{0, 0, 0, 1, 1, 1},
		[]uint8{0, 1, 2, 0, 1, 2},
		6,
	)
}