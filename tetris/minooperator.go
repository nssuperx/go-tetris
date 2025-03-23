package tetris

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// 単位は秒
const (
	idleLimit = 60.0 / 60.0
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
	idleTime float64
	dasTime  float64
	arrTime  float64 // 左右で分けてもいいかも
	lockTime float64
	mino     Mino
	bag      MinoBag
	field    *Field
}

func NewMinoOperator(field *Field) MinoOperator {
	return MinoOperator{
		idleTime: 0.0,
		dasTime:  0.0,
		arrTime:  0.0,
		lockTime: 0.0,
		bag:      NewMinoBag(),
		field:    field,
	}
}

func (o *MinoOperator) Update() {
	hardDropPos := getHardDropPos(&o.mino, o.field)
	o.field.SetGhost(&o.mino, hardDropPos)
	switch {
	// 右回転
	case inpututil.IsKeyJustPressed(ebiten.KeyI):
		shift, canRotate := canRotateRight(o.mino, o.field)
		if canRotate {
			o.mino.RotateRight(shift)
		}
		o.idleTime = 0.0
	// 左回転
	case inpututil.IsKeyJustPressed(ebiten.KeyJ):
		shift, canRotate := canRotateLeft(o.mino, o.field)
		if canRotate {
			o.mino.RotateLeft(shift)
		}
		o.idleTime = 0.0
	// 上入力
	case inpututil.IsKeyJustPressed(ebiten.KeyW):
		o.mino.HardDrop(hardDropPos)
		o.field.SetBlock(&o.mino)
		o.field.SetBlockColor(&o.mino)
		o.field.UpdateMinoFixed()
		o.SpawnMino()
		o.idleTime = 0.0
	// 右入力
	case inpututil.IsKeyJustPressed(ebiten.KeyD):
		if o.field.CanSetBlock(&o.mino, Vector2{1, 0}) {
			o.mino.MoveRight()
		}
	case ebiten.IsKeyPressed(ebiten.KeyD):
		o.dasTime += 1.0 / ebiten.ActualTPS()
		o.arrTime += 1.0 / ebiten.ActualTPS()
		if o.field.CanSetBlock(&o.mino, Vector2{1, 0}) && o.dasTime > dasLimit && o.arrTime > arrLimit {
			o.mino.MoveRight()
			o.arrTime = 0.0
		}
	case inpututil.IsKeyJustReleased(ebiten.KeyD):
		o.dasTime = 0.0
	// 左入力
	case inpututil.IsKeyJustPressed(ebiten.KeyA):
		if o.field.CanSetBlock(&o.mino, Vector2{-1, 0}) {
			o.mino.MoveLeft()
		}
	case ebiten.IsKeyPressed(ebiten.KeyA):
		o.dasTime += 1.0 / ebiten.ActualTPS()
		o.arrTime += 1.0 / ebiten.ActualTPS()
		if o.field.CanSetBlock(&o.mino, Vector2{-1, 0}) && o.dasTime > dasLimit && o.arrTime > arrLimit {
			o.mino.MoveLeft()
			o.arrTime = 0.0
		}
	case inpututil.IsKeyJustReleased(ebiten.KeyA):
		o.dasTime = 0.0
	// 下入力
	case ebiten.IsKeyPressed(ebiten.KeyS):
		o.arrTime += 1.0 / ebiten.ActualTPS()
		if o.field.CanSetBlock(&o.mino, Vector2{0, -1}) && o.arrTime > arrLimit {
			o.mino.MoveDown()
			o.idleTime = 0.0
			o.arrTime = 0.0
		}
	}
	o.idleTime += 1.0 / ebiten.ActualTPS()
	switch {
	case !o.field.CanSetBlock(&o.mino, Vector2{0, -1}) && o.idleTime > lockLimit:
		o.field.SetBlock(&o.mino)
		o.field.UpdateMinoFixed()
		o.SpawnMino()
		o.idleTime = 0.0
	case o.idleTime > idleLimit:
		o.mino.MoveDown()
		o.idleTime = 0.0
	}
	o.field.ResetFieldColor()
	o.field.SetBlockColor(&o.mino)
}

func (o *MinoOperator) SpawnMino() bool {
	minoType := o.bag.GetNextMino()
	switch minoType {
	case IMinoType:
		o.mino = NewIMino()
	case OMinoType:
		o.mino = NewOMino()
	case TMinoType:
		o.mino = NewTMino()
	case SMinoType:
		o.mino = NewSMino()
	case ZMinoType:
		o.mino = NewZMino()
	case JMinoType:
		o.mino = NewJMino()
	case LMinoType:
		o.mino = NewLMino()
	}

	o.field.SetBlockColor(&o.mino)
	if !o.field.CanSetBlock(&o.mino, Vector2{0, 0}) {
		Playing = false
		return false
	}
	return true
}
