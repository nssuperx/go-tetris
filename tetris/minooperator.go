package tetris

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// 単位は秒
const (
	fallLimit = 60.0 / 60.0
	dasLimit  = 15.0 / 60.0
	arrLimit  = 2.0 / 60.0
	lockLimit = 60.0 / 60.0
)

// 単位は回数。設置して何回操作できるか
const (
	onGroundMoveLimit   = 14
	onGroundRotateLimit = 14
)

type MinoOperator struct {
	fallTime    float64
	dasTime     float64
	arrTime     float64 // 左右で分けてもいいかも
	lockTime    float64
	moveCount   int
	rotateCount int
	mino        Mino
	nowMinoType MinoTypesEnum
	hold        MinoTypesEnum
	holded      bool
	bag         MinoBag
	hardDropPos Vector2
	field       *Field
	ui          *Ui
}

func NewMinoOperator(field *Field, ui *Ui) MinoOperator {
	return MinoOperator{
		fallTime:    0.0,
		dasTime:     0.0,
		arrTime:     0.0,
		lockTime:    0.0,
		moveCount:   0,
		rotateCount: 0,
		nowMinoType: Empty,
		hold:        Empty,
		holded:      false,
		bag:         newMinoBag(),
		hardDropPos: Vector2{0, 0},
		field:       field,
		ui:          ui,
	}
}

func (o *MinoOperator) init() {
	o.fallTime, o.dasTime, o.arrTime, o.lockTime = 0.0, 0.0, 0.0, 0.0
	o.moveCount, o.rotateCount = 0, 0
	o.hold = Empty
	o.holded = false
	o.bag = newMinoBag()
	o.hardDropPos = Vector2{0, 0}
	o.field.clear()
	o.spawnMino(o.bag.getNextMino())
	o.ui.nexts = o.bag.getNextMinos(NextMino)
}

func (o *MinoOperator) Update() {
	if !Playing {
		return
	}
	o.hardDropPos = getHardDropPos(&o.mino, o.field)
	o.field.setGhost(&o.mino, o.hardDropPos)
	inputed := o.input()
	minoFixed := o.fixOrFall()
	if inputed || minoFixed {
		o.field.resetFieldColor()
		o.field.setBlockColor(&o.mino)
	}
	o.fallTime += 1.0 / ebiten.ActualTPS()
}

func (o *MinoOperator) spawnMino(minoType MinoTypesEnum) bool {
	o.nowMinoType = minoType
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

// 入力周りの処理
// 何か入力があったらtrueを返す
func (o *MinoOperator) input() bool {
	switch {
	// ホールド
	case holdPressed():
		if o.holded {
			break
		}
		beforeMinoType := o.nowMinoType
		if o.hold == Empty {
			o.spawnMino(o.bag.getNextMino())
			o.ui.nexts = o.bag.getNextMinos(NextMino)
		} else {
			o.spawnMino(o.hold)
		}
		o.hold = beforeMinoType
		o.ui.hold = beforeMinoType
		o.holded = true
		return true
	// 右回転
	case rotateRightPressed():
		shift, canRotate := o.mino.canRotateRight(o.field)
		if canRotate {
			o.mino.rotateRight(shift)
			if !o.field.canSetBlock(&o.mino, Vector2{0, -1}) {
				o.rotateCount++
			}
			o.lockTime = 0.0
		}
		return true
	// 左回転
	case rotateLeftPressed():
		shift, canRotate := o.mino.canRotateLeft(o.field)
		if canRotate {
			o.mino.rotateLeft(shift)
			if !o.field.canSetBlock(&o.mino, Vector2{0, -1}) {
				o.rotateCount++
			}
			o.lockTime = 0.0
		}
		return true
	// 上入力
	case upPressed():
		o.mino.hardDrop(o.hardDropPos)
		o.fixMino()
		return true
	// 右入力
	case rightJustPressed():
		if o.field.canSetBlock(&o.mino, Vector2{1, 0}) {
			o.mino.moveRight()
			o.lockTime = 0.0
			o.arrTime = 0.0
			o.dasTime = 0.0
			if !o.field.canSetBlock(&o.mino, Vector2{0, -1}) {
				o.moveCount++
			}
		}
		return true
	// 左入力
	case leftJustPressed():
		if o.field.canSetBlock(&o.mino, Vector2{-1, 0}) {
			o.mino.moveLeft()
			o.lockTime = 0.0
			o.arrTime = 0.0
			o.dasTime = 0.0
			if !o.field.canSetBlock(&o.mino, Vector2{0, -1}) {
				o.moveCount++
			}
		}
		return true
	// 右長押し
	case rightPressed():
		o.dasTime += 1.0 / ebiten.ActualTPS()
		o.arrTime += 1.0 / ebiten.ActualTPS()
		if o.field.canSetBlock(&o.mino, Vector2{1, 0}) && o.dasTime > dasLimit && o.arrTime > arrLimit {
			o.mino.moveRight()
			o.lockTime = 0.0
			o.arrTime = 0.0
			if !o.field.canSetBlock(&o.mino, Vector2{0, -1}) {
				o.moveCount++
			}
		}
		return true
	// 左長押し
	case leftPressed():
		o.dasTime += 1.0 / ebiten.ActualTPS()
		o.arrTime += 1.0 / ebiten.ActualTPS()
		if o.field.canSetBlock(&o.mino, Vector2{-1, 0}) && o.dasTime > dasLimit && o.arrTime > arrLimit {
			o.mino.moveLeft()
			o.lockTime = 0.0
			o.arrTime = 0.0
			if !o.field.canSetBlock(&o.mino, Vector2{0, -1}) {
				o.moveCount++
			}
		}
		return true
	// 下入力
	case downPressed():
		o.arrTime += 1.0 / ebiten.ActualTPS()
		if o.field.canSetBlock(&o.mino, Vector2{0, -1}) && o.arrTime > arrLimit {
			o.mino.moveDown()
			o.fallTime = 0.0
			o.lockTime = 0.0
			o.arrTime = 0.0
		}
		return true
	}
	return false
}

// 入力に依存しない処理
// 動きがあったらtrueを返す
func (o *MinoOperator) fixOrFall() bool {
	switch {
	case !o.field.canSetBlock(&o.mino, Vector2{0, -1}) && (o.lockTime > lockLimit || o.moveCount > onGroundMoveLimit || o.rotateCount > onGroundRotateLimit):
		o.fixMino()
		return true
	case !o.field.canSetBlock(&o.mino, Vector2{0, -1}):
		o.lockTime += 1.0 / ebiten.ActualTPS()
		return false
	case o.fallTime > fallLimit:
		o.mino.moveDown()
		o.fallTime = 0.0
		o.lockTime = 0.0
		return true
	}
	return false
}

func (o *MinoOperator) fixMino() {
	o.field.setBlock(&o.mino)
	o.field.setBlockColor(&o.mino)
	o.field.updateMinoFixed()
	o.spawnMino(o.bag.getNextMino())
	o.ui.nexts = o.bag.getNextMinos(NextMino)
	o.holded = false
	o.moveCount, o.rotateCount = 0, 0
	o.fallTime, o.lockTime = 0.0, 0.0
}
