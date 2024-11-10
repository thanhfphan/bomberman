package assets

import (
	_ "embed"
)

//go:embed character/walk-right.png
var PlayerWalkRight []byte

//go:embed character/walk-front.png
var PlayerWalkFront []byte

//go:embed character/walk-back.png
var PlayerWalkBack []byte

//go:embed character/idle-right.png
var PlayerIdleRight []byte

//go:embed character/idle-front.png
var PlayerIdleFront []byte

//go:embed character/idle-back.png
var PlayerIdleBack []byte

//go:embed items/dynamite-pack.png
var BombSprite []byte

//go:embed audio/BG_MUSIC.wav
var BackGroundMusic []byte

//go:embed audio/BOM_SET.wav
var BomSetSound []byte

//go:embed audio/BOM_EXPLODE.wav
var BomExplodeSound []byte

//go:embed fxs/explosion.png
var BombExplosion []byte

//go:embed terrain/grass.png
var TerrainGrass []byte
