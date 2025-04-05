package tetris

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// 単位は秒
const (
	fallLimit = 60.0 / 60.0
	dasLimit  = 11.0 / 60.0
	arrLimit  = 2.0 / 60.0
	lockLimit = 40.0 / 60.0
)

// 単位は回数。設置して何回操作できるか
const (
	onGroundMoveLimit   = 15
	onGroundRotateLimit = 15
)

type MinoOperator struct {
	fallTime float64
	dasTime  float64
	arrTime  float64 // 左右で分けてもいいかも
	lockTime float64
	mino     Mino
	hold     MinoTypesEnum
	holded   bool
	bag      MinoBag
	field    *Field
	ui       *Ui
}

func NewMinoOperator(field *Field, ui *Ui) MinoOperator {
	return MinoOperator{
		fallTime: 0.0,
		dasTime:  0.0,
		arrTime:  0.0,
		lockTime: 0.0,
		hold:     Empty,
		holded:   false,
		bag:      newMinoBag(),
		field:    field,
		ui:       ui,
	}
}

func (o *MinoOperator) Update() {
	if !Playing && inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		Playing = true
		o.bag = newMinoBag()
		o.field.clear()
		o.hold = Empty
		o.ui.hold = Empty
		o.spawnMino(o.bag.getNextMino())
		o.ui.nexts = o.bag.getNextMinos(NextMino)
	}
	if !Playing {
		return
	}
	hardDropPos := getHardDropPos(&o.mino, o.field)
	o.field.setGhost(&o.mino, hardDropPos)
	switch {
	// ホールド
	case inpututil.IsKeyJustPressed(ebiten.KeyO):
		if o.holded {
			break
		}
		nowMinoType := o.mino.minoType
		if o.hold == Empty {
			o.spawnMino(o.bag.getNextMino())
			o.ui.nexts = o.bag.getNextMinos(NextMino)
		} else {
			o.spawnMino(o.hold)
		}
		o.hold = nowMinoType
		o.ui.hold = nowMinoType
		o.holded = true
	// 右回転
	case inpututil.IsKeyJustPressed(ebiten.KeyI):
		shift, canRotate := canRotateRight(o.mino, o.field)
		if canRotate {
			o.mino.rotateRight(shift)
		}
	// 左回転
	case inpututil.IsKeyJustPressed(ebiten.KeyJ):
		shift, canRotate := canRotateLeft(o.mino, o.field)
		if canRotate {
			o.mino.rotateLeft(shift)
		}
	// 上入力
	case inpututil.IsKeyJustPressed(ebiten.KeyW):
		o.mino.hardDrop(hardDropPos)
		o.field.setBlock(&o.mino)
		o.field.setBlockColor(&o.mino)
		o.field.updateMinoFixed()
		o.spawnMino(o.bag.getNextMino())
		o.ui.nexts = o.bag.getNextMinos(NextMino)
		o.fallTime = 0.0
		o.holded = false
	// 右入力
	case inpututil.IsKeyJustPressed(ebiten.KeyD):
		if o.field.canSetBlock(&o.mino, Vector2{1, 0}) {
			o.mino.moveRight()
		}
	case ebiten.IsKeyPressed(ebiten.KeyD):
		o.dasTime += 1.0 / ebiten.ActualTPS()
		o.arrTime += 1.0 / ebiten.ActualTPS()
		if o.field.canSetBlock(&o.mino, Vector2{1, 0}) && o.dasTime > dasLimit && o.arrTime > arrLimit {
			o.mino.moveRight()
			o.arrTime = 0.0
		}
	case inpututil.IsKeyJustReleased(ebiten.KeyD):
		o.dasTime = 0.0
	// 左入力
	case inpututil.IsKeyJustPressed(ebiten.KeyA):
		if o.field.canSetBlock(&o.mino, Vector2{-1, 0}) {
			o.mino.moveLeft()
		}
	case ebiten.IsKeyPressed(ebiten.KeyA):
		o.dasTime += 1.0 / ebiten.ActualTPS()
		o.arrTime += 1.0 / ebiten.ActualTPS()
		if o.field.canSetBlock(&o.mino, Vector2{-1, 0}) && o.dasTime > dasLimit && o.arrTime > arrLimit {
			o.mino.moveLeft()
			o.arrTime = 0.0
		}
	case inpututil.IsKeyJustReleased(ebiten.KeyA):
		o.dasTime = 0.0
	// 下入力
	case ebiten.IsKeyPressed(ebiten.KeyS):
		o.arrTime += 1.0 / ebiten.ActualTPS()
		if o.field.canSetBlock(&o.mino, Vector2{0, -1}) && o.arrTime > arrLimit {
			o.mino.moveDown()
			o.fallTime = 0.0
			o.arrTime = 0.0
		}
	}
	o.fallTime += 1.0 / ebiten.ActualTPS()
	switch {
	case !o.field.canSetBlock(&o.mino, Vector2{0, -1}) && o.fallTime > lockLimit:
		o.field.setBlock(&o.mino)
		o.field.updateMinoFixed()
		o.spawnMino(o.bag.getNextMino())
		o.ui.nexts = o.bag.getNextMinos(NextMino)
		o.holded = false
		o.fallTime = 0.0
	case o.fallTime > fallLimit:
		o.mino.moveDown()
		o.fallTime = 0.0
	}
	o.field.resetFieldColor()
	o.field.setBlockColor(&o.mino)
}

func (o *MinoOperator) spawnMino(minoType MinoTypesEnum) bool {
	switch minoType {
	case IMinoType:
		o.mino = newIMino()
	case OMinoType:
		o.mino = newOMino()
	case TMinoType:
		o.mino = newTMino()
	case SMinoType:
		o.mino = newSMino()
	case ZMinoType:
		o.mino = newZMino()
	case JMinoType:
		o.mino = newJMino()
	case LMinoType:
		o.mino = newLMino()
	}

	o.field.setBlockColor(&o.mino)
	if !o.field.canSetBlock(&o.mino, Vector2{0, 0}) {
		Playing = false
		return false
	}
	return true
}
