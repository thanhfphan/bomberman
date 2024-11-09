package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"thanhfphan.com/bomberman/src/game"
)

func main() {
	width, height := 1088, 960
	game := game.New(width, height)
	if err := game.Setup("config.ini"); err != nil {
		log.Fatalf("could not setup game: %v", err)
	}
	game.Init()

	ebiten.SetWindowTitle("Bomberman")
	ebiten.SetWindowSize(width, height)

	if err := ebiten.RunGame(game); err != nil && err != ebiten.Termination {
		log.Fatal(err)
	}
}
