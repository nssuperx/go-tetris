package tetris

import (
	"fmt"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	ScreenWidth  = 1280
	ScreenHeight = 720
)

var Playing = false
var Warm = false

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
	if !Warm && !math.IsInf(1.0/ebiten.ActualTPS(), 0) {
		Playing = true
		Warm = true
		g.minoOperator.SpawnMino(g.minoOperator.bag.GetNextMino())
		return nil
	}
	if Playing {
		g.minoOperator.Update()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%f", 1.0/ebiten.ActualTPS()), 0, 0)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%d", g.minoOperator.mino.direction), 0, 10)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%d", g.minoOperator.hold), 0, 20)
	if !Playing {
		ebitenutil.DebugPrintAt(screen, "Game Over", 0, 30)
	}
	drawField(screen, g.field)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
