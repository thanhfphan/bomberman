package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"thanhfphan.com/bomberman/src/game"
)

func main() {
	width, height := 1088, 960
	game := game.New(width, height)
	if err := game.Setup(); err != nil {
		log.Fatalf("could not setup game: %v", err)
	}
	if err := game.LoadConfig("config.ini"); err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	ebiten.SetWindowTitle("Bomberman")
	ebiten.SetWindowSize(width, height)

	if err := ebiten.RunGame(game); err != nil && err != ebiten.Termination {
		log.Fatal(err)
	}
}
