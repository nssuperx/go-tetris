package main

import (
	"log"

	"github.com/nssuperx/go-tetris/tetris"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(tetris.ScreenWidth, tetris.ScreenHeight)
	ebiten.SetWindowTitle("Tetris")
	game := tetris.NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
