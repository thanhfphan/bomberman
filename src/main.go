package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"thanhfphan.com/bomberman/src/game"
)

func main() {
	g := game.New()
	if err := g.Setup("config.ini"); err != nil {
		log.Fatalf("could not setup game: %v", err)
	}
	g.Init()

	ebiten.SetWindowTitle("Bomberman")
	ebiten.SetWindowSize(game.WindowWidth, game.WindowHeight)

	if err := ebiten.RunGame(g); err != nil && err != ebiten.Termination {
		log.Fatal(err)
	}
}
