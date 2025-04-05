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
	shape     [][]bool
	direction MinoDirEnum
	color     color.RGBA
}

func NewOMino() Mino {
	return Mino{
		minoType: OMinoType,
		pos:      Vector2{4, 20},
		shape: [][]bool{
			{true, true},
			{true, true},
		},
		direction: Up,
		color:     yellow,
	}
}

func NewIMino() Mino {
	return Mino{
		minoType: IMinoType,
		pos:      Vector2{3, 20},
		shape: [][]bool{
			{false, false, false, false},
			{true, true, true, true},
			{false, false, false, false},
			{false, false, false, false},
		},
		direction: Up,
		color:     lightBlue,
	}
}

func NewTMino() Mino {
	return Mino{
		minoType: TMinoType,
		pos:      Vector2{3, 20},
		shape: [][]bool{
			{false, true, false},
			{true, true, true},
			{false, false, false},
		},
		direction: Up,
		color:     purple,
	}
}

func NewSMino() Mino {
	return Mino{
		minoType: SMinoType,
		pos:      Vector2{3, 20},
		shape: [][]bool{
			{false, true, true},
			{true, true, false},
			{false, false, false},
		},
		direction: Up,
		color:     green,
	}
}

func NewZMino() Mino {
	return Mino{
		minoType: ZMinoType,
		pos:      Vector2{3, 20},
		shape: [][]bool{
			{true, true, false},
			{false, true, true},
			{false, false, false},
		},
		direction: Up,
		color:     red,
	}
}

func NewLMino() Mino {
	return Mino{
		minoType: LMinoType,
		pos:      Vector2{3, 20},
		shape: [][]bool{
			{false, false, true},
			{true, true, true},
			{false, false, false},
		},
		direction: Up,
		color:     orange,
	}
}

func NewJMino() Mino {
	return Mino{
		minoType: JMinoType,
		pos:      Vector2{3, 20},
		shape: [][]bool{
			{true, false, false},
			{true, true, true},
			{false, false, false},
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
	if m.minoType == OMinoType {
		return
	}
	m.pos = m.pos.Add(shiftPos)
	rotateMinoRight(m)
}

func (m *Mino) RotateLeft(shiftPos Vector2) {
	if m.minoType == OMinoType {
		return
	}
	m.pos = m.pos.Add(shiftPos)
	rotateMinoLeft(m)
}

func convertShapeToPos(shape [][]bool) []Vector2 {
	var positions []Vector2
	for y, row := range shape {
		for x, block := range row {
			if block {
				positions = append(positions, Vector2{x, -y})
			}
		}
	}
	return positions
}
