package tetris

import (
	"github.com/hajimehoshi/ebiten/v2"
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
	arrTime  float64
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
	// 左右の入力があったらdas, arr, lockに関わる操作
	// 上が入力されたら
	// ハードドロップ
	// idletimeをリセット
	// 下が入力されたら
	// 1ます下に移動
	// idletimeをリセット
	// 左右が入力されたら
	// 1ます左右に移動
	// dasTimeを足す
	// 左右が入力され続けていたら
	// dasTimeが11/60を超えたら
	// 1ます左右に移動
	// arrTimeを足す
	// arrTimeが2/60を超えたら
	// 1ます左右に移動
	// 左右が入力されていなかったら
	// dasTimeをリセット
	// arrTimeをリセット
	switch {
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
