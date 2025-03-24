package tetris

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Ui struct {
	minoPanels map[MinoTypesEnum]*ebiten.Image
	hold       MinoTypesEnum
	nexts      []MinoTypesEnum
}

const (
	panelSizeX  = 100
	uiBlockSize = 20
	panelSizeY  = panelSizeX - uiBlockSize
)

func NewUi() *Ui {
	var ui Ui
	ui.hold = Empty
	ui.nexts = make([]MinoTypesEnum, NextMino, NextMino)
	for i := range NextMino {
		ui.nexts[i] = Empty
	}
	ui.minoPanels = make(map[MinoTypesEnum]*ebiten.Image, 8)
	// Iミノ
	panel := ebiten.NewImage(panelSizeX, panelSizeY)
	vector.StrokeRect(panel, 0, 0, panelSizeX, panelSizeY, 4, color.White, false)
	mino := ebiten.NewImage(uiBlockSize*4, uiBlockSize)
	vector.DrawFilledRect(mino, 0, 0, float32(uiBlockSize*4), float32(uiBlockSize), lightBlue, false)
	for i := 1; i <= 3; i++ {
		vector.StrokeLine(mino, float32(uiBlockSize*i), 0, float32(uiBlockSize*i), float32(uiBlockSize), 1, color.Black, false)
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(panelSizeX-mino.Bounds().Dx())/2.0, float64(panelSizeY-mino.Bounds().Dy())/2.0)
	panel.DrawImage(mino, op)
	ui.minoPanels[IMinoType] = panel
	// Oミノ
	panel = ebiten.NewImage(panelSizeX, panelSizeY)
	vector.StrokeRect(panel, 0, 0, panelSizeX, panelSizeY, 4, color.White, false)
	mino = ebiten.NewImage(uiBlockSize*2, uiBlockSize*2)
	vector.DrawFilledRect(mino, 0, 0, float32(uiBlockSize*2), float32(uiBlockSize*2), yellow, false)
	vector.StrokeLine(mino, 0, float32(uiBlockSize), float32(uiBlockSize*2), float32(uiBlockSize), 1, color.Black, false)
	vector.StrokeLine(mino, float32(uiBlockSize), 0, float32(uiBlockSize), float32(uiBlockSize*2), 1, color.Black, false)
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(panelSizeX-mino.Bounds().Dx())/2.0, float64(panelSizeY-mino.Bounds().Dy())/2.0)
	panel.DrawImage(mino, op)
	ui.minoPanels[OMinoType] = panel
	// Tミノ
	panel = ebiten.NewImage(panelSizeX, panelSizeY)
	vector.StrokeRect(panel, 0, 0, panelSizeX, panelSizeY, 4, color.White, false)
	mino = ebiten.NewImage(uiBlockSize*3, uiBlockSize*2)
	vector.DrawFilledRect(mino, float32(uiBlockSize), 0, float32(uiBlockSize), float32(uiBlockSize), purple, false)
	vector.DrawFilledRect(mino, 0, float32(uiBlockSize), float32(uiBlockSize*3), float32(uiBlockSize*2), purple, false)
	vector.StrokeLine(mino, float32(uiBlockSize), float32(uiBlockSize), float32(uiBlockSize*2), float32(uiBlockSize), 1, color.Black, false)
	vector.StrokeLine(mino, float32(uiBlockSize), float32(uiBlockSize), float32(uiBlockSize), float32(uiBlockSize*2), 1, color.Black, false)
	vector.StrokeLine(mino, float32(uiBlockSize*2), float32(uiBlockSize), float32(uiBlockSize*2), float32(uiBlockSize*2), 1, color.Black, false)
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(panelSizeX-mino.Bounds().Dx())/2.0, float64(panelSizeY-mino.Bounds().Dy())/2.0)
	panel.DrawImage(mino, op)
	ui.minoPanels[TMinoType] = panel
	// Sミノ
	panel = ebiten.NewImage(panelSizeX, panelSizeY)
	vector.StrokeRect(panel, 0, 0, panelSizeX, panelSizeY, 4, color.White, false)
	mino = ebiten.NewImage(uiBlockSize*3, uiBlockSize*2)
	vector.DrawFilledRect(mino, float32(uiBlockSize), 0, float32(uiBlockSize*3), float32(uiBlockSize), green, false)
	vector.DrawFilledRect(mino, 0, float32(uiBlockSize), float32(uiBlockSize*2), float32(uiBlockSize*2), green, false)
	vector.StrokeLine(mino, float32(uiBlockSize*2), 0, float32(uiBlockSize*2), float32(uiBlockSize), 1, color.Black, false)
	vector.StrokeLine(mino, float32(uiBlockSize), float32(uiBlockSize), float32(uiBlockSize*2), float32(uiBlockSize), 1, color.Black, false)
	vector.StrokeLine(mino, float32(uiBlockSize), float32(uiBlockSize), float32(uiBlockSize), float32(uiBlockSize*2), 1, color.Black, false)
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(panelSizeX-mino.Bounds().Dx())/2.0, float64(panelSizeY-mino.Bounds().Dy())/2.0)
	panel.DrawImage(mino, op)
	ui.minoPanels[SMinoType] = panel
	// Zミノ
	panel = ebiten.NewImage(panelSizeX, panelSizeY)
	vector.StrokeRect(panel, 0, 0, panelSizeX, panelSizeY, 4, color.White, false)
	mino = ebiten.NewImage(uiBlockSize*3, uiBlockSize*2)
	vector.DrawFilledRect(mino, 0, 0, float32(uiBlockSize*2), float32(uiBlockSize), red, false)
	vector.DrawFilledRect(mino, float32(uiBlockSize), float32(uiBlockSize), float32(uiBlockSize*3), float32(uiBlockSize*2), red, false)
	vector.StrokeLine(mino, float32(uiBlockSize), 0, float32(uiBlockSize), float32(uiBlockSize), 1, color.Black, false)
	vector.StrokeLine(mino, float32(uiBlockSize), float32(uiBlockSize), float32(uiBlockSize*2), float32(uiBlockSize), 1, color.Black, false)
	vector.StrokeLine(mino, float32(uiBlockSize*2), float32(uiBlockSize), float32(uiBlockSize*2), float32(uiBlockSize*2), 1, color.Black, false)
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(panelSizeX-mino.Bounds().Dx())/2.0, float64(panelSizeY-mino.Bounds().Dy())/2.0)
	panel.DrawImage(mino, op)
	ui.minoPanels[ZMinoType] = panel
	// Jミノ
	panel = ebiten.NewImage(panelSizeX, panelSizeY)
	vector.StrokeRect(panel, 0, 0, panelSizeX, panelSizeY, 4, color.White, false)
	mino = ebiten.NewImage(uiBlockSize*3, uiBlockSize*2)
	vector.DrawFilledRect(mino, 0, 0, float32(uiBlockSize), float32(uiBlockSize), darkBlue, false)
	vector.DrawFilledRect(mino, 0, float32(uiBlockSize), float32(uiBlockSize*3), float32(uiBlockSize*2), darkBlue, false)
	vector.StrokeLine(mino, 0, float32(uiBlockSize), float32(uiBlockSize), float32(uiBlockSize), 1, color.Black, false)
	vector.StrokeLine(mino, float32(uiBlockSize), float32(uiBlockSize), float32(uiBlockSize), float32(uiBlockSize*2), 1, color.Black, false)
	vector.StrokeLine(mino, float32(uiBlockSize*2), float32(uiBlockSize), float32(uiBlockSize*2), float32(uiBlockSize*2), 1, color.Black, false)
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(panelSizeX-mino.Bounds().Dx())/2.0, float64(panelSizeY-mino.Bounds().Dy())/2.0)
	panel.DrawImage(mino, op)
	ui.minoPanels[JMinoType] = panel
	// Lミノ
	panel = ebiten.NewImage(panelSizeX, panelSizeY)
	vector.StrokeRect(panel, 0, 0, panelSizeX, panelSizeY, 4, color.White, false)
	mino = ebiten.NewImage(uiBlockSize*3, uiBlockSize*2)
	vector.DrawFilledRect(mino, float32(uiBlockSize*2), 0, float32(uiBlockSize*3), float32(uiBlockSize*2), orange, false)
	vector.DrawFilledRect(mino, 0, float32(uiBlockSize), float32(uiBlockSize*3), float32(uiBlockSize*2), orange, false)
	vector.StrokeLine(mino, float32(uiBlockSize*2), float32(uiBlockSize), float32(uiBlockSize*3), float32(uiBlockSize), 1, color.Black, false)
	vector.StrokeLine(mino, float32(uiBlockSize), float32(uiBlockSize), float32(uiBlockSize), float32(uiBlockSize*2), 1, color.Black, false)
	vector.StrokeLine(mino, float32(uiBlockSize*2), float32(uiBlockSize), float32(uiBlockSize*2), float32(uiBlockSize*2), 1, color.Black, false)
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(panelSizeX-mino.Bounds().Dx())/2.0, float64(panelSizeY-mino.Bounds().Dy())/2.0)
	panel.DrawImage(mino, op)
	ui.minoPanels[LMinoType] = panel
	// 空
	panel = ebiten.NewImage(panelSizeX, panelSizeY)
	vector.StrokeRect(panel, 0, 0, panelSizeX, panelSizeY, 4, color.White, false)
	ui.minoPanels[Empty] = panel
	return &ui
}

func (u *Ui) Draw(screen *ebiten.Image) {
	u.drawHoldMino(screen)
	u.drawNextMinos(screen)
}

func (u *Ui) drawHoldMino(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(fieldX-panelSizeX-20, fieldY)
	screen.DrawImage(u.minoPanels[u.hold], op)
}

func (u *Ui) drawNextMinos(screen *ebiten.Image) {
	for i, minoType := range u.nexts {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(fieldX+fieldWidth+20, float64(fieldY+i*panelSizeX))
		screen.DrawImage(u.minoPanels[minoType], op)
	}
}
