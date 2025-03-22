package tetris

import "image/color"

type MinoDirEnum int

const (
	Up MinoDirEnum = iota
	Down
	Left
	Right
)

type Mino struct {
	pos       Vector2
	shape     []Vector2 // 0番目がミノの底で、出現位置とする。一応半時計周りになっている
	direction MinoDirEnum
	color     color.RGBA
}

type MinoInterface interface {
	GetPos() Vector2
	SetPos(pos Vector2)
	GetColor() color.RGBA
}

func NewOMino() Mino {
	// Oミノは回転軸がブロックではなく、格子の位置（回転しない）
	return Mino{
		pos: Vector2{4, 19},
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
	return Mino{
		pos: Vector2{3, 19},
		shape: []Vector2{
			{0, 0}, {1, 0}, {2, 0}, {3, 0},
		},
		direction: Up,
		color:     lightBlue,
	}
}

func NewTMino() Mino {
	return Mino{
		pos: Vector2{4, 19},
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
		pos: Vector2{4, 19},
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
		pos: Vector2{4, 19},
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
		pos: Vector2{4, 19},
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
		pos: Vector2{4, 19},
		shape: []Vector2{
			{0, 0},
			{1, 0}, {-1, 1}, {-1, 0},
		},
		direction: Up,
		color:     darkBlue,
	}
}
