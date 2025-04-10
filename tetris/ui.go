package tetris

import (
	"bytes"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type DebugPanelEnum int

const (
	Das DebugPanelEnum = iota
	Arr
	Fall
	Lock
	RotateCount
	MoveCount
)

type Ui struct {
	minoPanels     map[MinoTypesEnum]*ebiten.Image
	hold           MinoTypesEnum
	nexts          []MinoTypesEnum
	textFaceSource *text.GoTextFaceSource
}

const (
	uiBlockSize     = 20
	panelSizeX      = uiBlockSize * 5
	panelSizeY      = uiBlockSize * 4
	debugPanelSizeX = uiBlockSize * 3
	debugPanelSizeY = uiBlockSize * 5
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

	// フォント
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}
	ui.textFaceSource = s
	return &ui
}

func (u *Ui) init() {
	u.hold = Empty
	for i := range u.nexts {
		u.nexts[i] = Empty
	}
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

// うまいやり方があれば廃止
type debugUiPos struct {
	descX float64 // descriptionのx
	descY float64 // descriptionのy
	valX  float32 // valueのx
	valY  float32 // valueのy
}

// できれば扇形を書きたい
func (u *Ui) drawDebugUi(screen *ebiten.Image, o *MinoOperator) {
	f := &text.GoTextFace{
		Source: u.textFaceSource,
		Size:   18,
	}

	// imageに描いてscreen(image)に描くとフォントが縁がきれいにならない？
	// 本来はテキストと図形を小さいimageに描いてscreenに描きたかった
	textOp := &text.DrawOptions{}
	textOp.PrimaryAlign = text.AlignCenter

	das := debugUiPos{
		descX: fieldX - debugPanelSizeX*2 - 20,
		descY: fieldY + fieldHeight/3.0 + 10,
		valX:  fieldX - debugPanelSizeX*2 - 20,
		valY:  fieldY + fieldHeight/3.0 + debugPanelSizeY/1.5,
	}
	textOp.GeoM.Translate(das.descX, das.descY)
	text.Draw(screen, "DAS", f, textOp)

	arr := debugUiPos{
		descX: fieldX - debugPanelSizeX - 5,
		descY: fieldY + fieldHeight/3.0 + 10,
		valX:  fieldX - debugPanelSizeX - 5,
		valY:  fieldY + fieldHeight/3.0 + debugPanelSizeY/1.5,
	}
	textOp.GeoM.Reset()
	textOp.GeoM.Translate(arr.descX, arr.descY)
	text.Draw(screen, "ARR", f, textOp)

	fall := debugUiPos{
		descX: fieldX - debugPanelSizeX*2 - 20,
		descY: fieldY + fieldHeight/3.0 + debugPanelSizeY + 10,
		valX:  fieldX - debugPanelSizeX*2 - 20,
		valY:  fieldY + fieldHeight/3 + debugPanelSizeY + debugPanelSizeY/1.5,
	}
	textOp.GeoM.Reset()
	textOp.GeoM.Translate(fall.descX, fall.descY)
	text.Draw(screen, "FALL", f, textOp)

	lock := debugUiPos{
		descX: fieldX - debugPanelSizeX - 5,
		descY: fieldY + fieldHeight/3.0 + debugPanelSizeY + 10,
		valX:  fieldX - debugPanelSizeX - 5,
		valY:  fieldY + fieldHeight/3.0 + debugPanelSizeY + debugPanelSizeY/1.5,
	}
	textOp.GeoM.Reset()
	textOp.GeoM.Translate(lock.descX, lock.descY)
	text.Draw(screen, "LOCK", f, textOp)

	rotate := debugUiPos{
		descX: fieldX - debugPanelSizeX*2 - 20,
		descY: fieldY + fieldHeight/3.0 + debugPanelSizeY*2 + 10,
		valX:  fieldX - debugPanelSizeX*2 - 20,
		valY:  fieldY + fieldHeight/3.0 + debugPanelSizeY*2 + debugPanelSizeY/1.5,
	}
	textOp.GeoM.Reset()
	textOp.GeoM.Translate(rotate.descX, rotate.descY)
	text.Draw(screen, "ROTATE", f, textOp)

	move := debugUiPos{
		descX: fieldX - debugPanelSizeX - 5,
		descY: fieldY + fieldHeight/3.0 + debugPanelSizeY*2 + 10,
		valX:  fieldX - debugPanelSizeX - 5,
		valY:  fieldY + fieldHeight/3.0 + debugPanelSizeY*2 + debugPanelSizeY/1.5,
	}
	textOp.GeoM.Reset()
	textOp.GeoM.Translate(move.descX, move.descY)
	text.Draw(screen, "MOVE", f, textOp)

	r := float32(debugPanelSizeX / 2.1)
	vector.StrokeCircle(screen, das.valX, das.valY, r, 1, color.White, true)
	vector.DrawFilledCircle(screen, das.valX, das.valY, min(float32(o.dasTime/dasLimit), 1)*r, color.White, true)
	vector.StrokeCircle(screen, arr.valX, arr.valY, r, 1, color.White, true)
	vector.DrawFilledCircle(screen, arr.valX, arr.valY, min(float32(o.arrTime/arrLimit), 1)*r, color.White, true)
	vector.StrokeCircle(screen, fall.valX, fall.valY, r, 1, color.White, true)
	vector.DrawFilledCircle(screen, fall.valX, fall.valY, min(float32(o.fallTime/fallLimit), 1)*r, color.White, true)
	vector.StrokeCircle(screen, lock.valX, lock.valY, r, 1, color.White, true)
	vector.DrawFilledCircle(screen, lock.valX, lock.valY, min(float32(o.lockTime/lockLimit), 1)*r, color.White, true)
	vector.StrokeCircle(screen, rotate.valX, rotate.valY, r, 1, color.White, true)
	vector.DrawFilledCircle(screen, rotate.valX, rotate.valY, min(float32(o.rotateCount)/float32(onGroundRotateLimit), 1)*r, color.White, true)
	vector.StrokeCircle(screen, move.valX, move.valY, r, 1, color.White, true)
	vector.DrawFilledCircle(screen, move.valX, move.valY, min(float32(o.moveCount)/float32(onGroundMoveLimit), 1)*r, color.White, true)
}
