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
	pos       Vector2
	shape     [][]bool
	direction MinoDirEnum
	color     color.RGBA
	rs        RotationSystem
}

func newOMino() Mino {
	return Mino{
		pos: Vector2{4, 20},
		shape: [][]bool{
			{true, true},
			{true, true},
		},
		direction: Up,
		color:     yellow,
		rs:        &OMinoRotationSystem{},
	}
}

func newIMino() Mino {
	return Mino{
		pos: Vector2{3, 20},
		shape: [][]bool{
			{false, false, false, false},
			{true, true, true, true},
			{false, false, false, false},
			{false, false, false, false},
		},
		direction: Up,
		color:     lightBlue,
		rs:        &IMinoRotationSystem{},
	}
}

func newTMino() Mino {
	return Mino{
		pos: Vector2{3, 20},
		shape: [][]bool{
			{false, true, false},
			{true, true, true},
			{false, false, false},
		},
		direction: Up,
		color:     purple,
		rs:        &CommonMinoRotationSystem{},
	}
}

func newSMino() Mino {
	return Mino{
		pos: Vector2{3, 20},
		shape: [][]bool{
			{false, true, true},
			{true, true, false},
			{false, false, false},
		},
		direction: Up,
		color:     green,
		rs:        &CommonMinoRotationSystem{},
	}
}

func newZMino() Mino {
	return Mino{
		pos: Vector2{3, 20},
		shape: [][]bool{
			{true, true, false},
			{false, true, true},
			{false, false, false},
		},
		direction: Up,
		color:     red,
		rs:        &CommonMinoRotationSystem{},
	}
}

func newLMino() Mino {
	return Mino{
		pos: Vector2{3, 20},
		shape: [][]bool{
			{false, false, true},
			{true, true, true},
			{false, false, false},
		},
		direction: Up,
		color:     orange,
		rs:        &CommonMinoRotationSystem{},
	}
}

func newJMino() Mino {
	return Mino{
		pos: Vector2{3, 20},
		shape: [][]bool{
			{true, false, false},
			{true, true, true},
			{false, false, false},
		},
		direction: Up,
		color:     darkBlue,
		rs:        &CommonMinoRotationSystem{},
	}
}

func (m *Mino) moveDown() {
	m.pos.y--
}

func (m *Mino) moveLeft() {
	m.pos.x--
}

func (m *Mino) moveRight() {
	m.pos.x++
}

func (m *Mino) hardDrop(pos Vector2) {
	m.pos = pos
}

func (m *Mino) canRotateRight(field *Field) (Vector2, bool) {
	shift, canRotate := m.rs.canRotateRight(m, field)
	if canRotate {
		return shift, true
	}
	return Vector2{0, 0}, false
}

func (m *Mino) canRotateLeft(field *Field) (Vector2, bool) {
	shift, canRotate := m.rs.canRotateLeft(m, field)
	if canRotate {
		return shift, true
	}
	return Vector2{0, 0}, false
}

func (m *Mino) rotateRight(shiftPos Vector2) {
	m.pos = m.pos.Add(shiftPos)
	rotateMinoRight(m)
}

func (m *Mino) rotateLeft(shiftPos Vector2) {
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
