package tetris

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	ScreenWidth  = 1280
	ScreenHeight = 720
)

var Playing = false

type Game struct {
	field        *Field
	minoOperator MinoOperator
}

func NewGame() *Game {
	field := Field{}
	minoOperator := NewMinoOperator(&field)
	return &Game{
		field:        &field,
		minoOperator: minoOperator,
	}
}

func (g *Game) Update() error {
	if !Playing {
		Playing = true
		g.minoOperator.SpawnMino()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("%f\n", 1.0/ebiten.ActualTPS()))
	drawField(screen, g.field)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
