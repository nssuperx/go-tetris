package tetris

type Vector2 struct {
	x, y int
}

func (v Vector2) Add(v2 Vector2) Vector2 {
	return Vector2{v.x + v2.x, v.y + v2.y}
}

func getHardDropPos(mino *Mino, field *Field) Vector2 {
	y := mino.pos.y
	for {
		y--
		if !field.canSetBlock(mino, Vector2{0, y - mino.pos.y}) {
			y++
			break
		}
	}
	return Vector2{mino.pos.x, y}
}
