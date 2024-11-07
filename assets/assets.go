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
