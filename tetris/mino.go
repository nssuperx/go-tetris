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
	shape     []Vector2
	direction MinoDirEnum
	color     color.RGBA
}

type MinoInterface interface {
	GetPos() Vector2
	GetShape() []Vector2
	SetPos(pos Vector2)
	GetColor() color.RGBA
}

func NewOMino() Mino {
	return Mino{
		pos: Vector2{4, 0},
		shape: []Vector2{
			{0, 0}, {1, 0},
			{0, 1}, {1, 1},
		},
		direction: Up,
		color:     color.RGBA{255, 255, 0, 255},
	}
}
