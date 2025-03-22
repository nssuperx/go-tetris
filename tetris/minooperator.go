package tetris

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// 単位は秒
const (
	idleLimit = 60 / 60
	dasLimit  = 11 / 60
	arrLimit  = 2 / 60
	lockLimit = 40 / 60
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
	o.idleTime += 1.0 / ebiten.ActualTPS()
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
	if !o.field.CanSetBlock(&o.mino) {
		return false
	}
	return true
}
