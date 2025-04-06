package tetris

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 720
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
	if startPressed() {
		Playing = true
		g.ui.init()
		g.field.clear()
		g.minoOperator.init()
	}
	g.minoOperator.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	drawField(screen, g.field)
	g.ui.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
