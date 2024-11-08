package game

import (
	"thanhfphan.com/bomberman/src/engine/audio"
)

func (g *Game) Init() {
	audio.Play(g.backgroundMusic)
}
