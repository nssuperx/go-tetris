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

const (
	NextMino = 4
)

var Playing = false
var Warm = false

type Game struct {
	field        *Field
	minoOperator MinoOperator
	ui           *Ui
}

func NewGame() *Game {
	field := Field{}
	ui := NewUi()
	minoOperator := NewMinoOperator(&field, ui)
	return &Game{
		field:        &field,
		minoOperator: minoOperator,
		ui:           ui,
	}
}

func (g *Game) Update() error {
	g.minoOperator.Update()
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
	g.ui.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
