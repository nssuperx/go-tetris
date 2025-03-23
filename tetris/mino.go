package tetris

import "image/color"

type MinoDirEnum int

const (
	Up MinoDirEnum = iota
	Right
	Down
	Left
)

type Mino struct {
	minoType  MinoTypesEnum
	pos       Vector2
	shape     []Vector2 // 0番目がミノの底で、出現位置とする。一応半時計周りになっている
	direction MinoDirEnum
	color     color.RGBA
}

func NewOMino() Mino {
	// Oミノは回転軸がブロックではなく、格子の位置（回転しない）
	return Mino{
		minoType: OMinoType,
		pos:      Vector2{4, 19},
		shape: []Vector2{
			{0, 0}, {1, 0},
			{1, 1}, {0, 1},
		},
		direction: Up,
		color:     yellow,
	}
}

func NewIMino() Mino {
	// Iミノは回転軸がブロックではなく、格子の位置
	// いい感じに回転させる方法を思いつかないので、とりあえず4方向手で書く
	// 4x4で回せる関数を作るのがいいかもしれない
	// shapeの順番はIミノだけ例外
	return Mino{
		minoType: IMinoType,
		pos:      Vector2{4, 18},
		shape: []Vector2{
			{-1, 1}, {0, 1}, {1, 1}, {2, 1},
		},
		direction: Up,
		color:     lightBlue,
	}
}

func NewTMino() Mino {
	return Mino{
		minoType: TMinoType,
		pos:      Vector2{4, 19},
		shape: []Vector2{
			{0, 0},
			{1, 0}, {0, 1}, {-1, 0},
		},
		direction: Up,
		color:     purple,
	}
}

func NewSMino() Mino {
	return Mino{
		minoType: SMinoType,
		pos:      Vector2{4, 19},
		shape: []Vector2{
			{0, 0},
			{1, 1}, {0, 1}, {-1, 0},
		},
		direction: Up,
		color:     green,
	}
}

func NewZMino() Mino {
	return Mino{
		minoType: ZMinoType,
		pos:      Vector2{4, 19},
		shape: []Vector2{
			{0, 0},
			{1, 0}, {0, 1}, {-1, 1},
		},
		direction: Up,
		color:     red,
	}
}

func NewLMino() Mino {
	return Mino{
		minoType: LMinoType,
		pos:      Vector2{4, 19},
		shape: []Vector2{
			{0, 0},
			{1, 0}, {1, 1}, {-1, 0},
		},
		direction: Up,
		color:     orange,
	}
}

func NewJMino() Mino {
	return Mino{
		minoType: JMinoType,
		pos:      Vector2{4, 19},
		shape: []Vector2{
			{0, 0},
			{1, 0}, {-1, 1}, {-1, 0},
		},
		direction: Up,
		color:     darkBlue,
	}
}

func (m *Mino) MoveDown() {
	m.pos.y--
}

func (m *Mino) MoveLeft() {
	m.pos.x--
}

func (m *Mino) MoveRight() {
	m.pos.x++
}

func (m *Mino) HardDrop(pos Vector2) {
	m.pos = pos
}

func (m *Mino) RotateRight(shiftPos Vector2) {
	// TODO: Iミノはあとで
	if m.minoType == OMinoType {
		return
	}
	m.pos = m.pos.Add(shiftPos)
	if m.minoType == IMinoType {
		rotateIMinoRight(m)
		return
	}
	rotateMinoRight(m)
}

func (m *Mino) RotateLeft(shiftPos Vector2) {
	if m.minoType == OMinoType {
		return
	}
	m.pos = m.pos.Add(shiftPos)
	if m.minoType == IMinoType {
		rotateIMinoLeft(m)
		return
	}
	rotateMinoLeft(m)
}
