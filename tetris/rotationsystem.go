package tetris

func rotatePosLeft(pos Vector2) Vector2 {
	return Vector2{-pos.y, pos.x}
}

func rotatePosRight(pos Vector2) Vector2 {
	return Vector2{pos.y, -pos.x}
}

func rotateMinoLeft(mino *Mino) {
	for i, s := range mino.shape {
		mino.shape[i] = rotatePosLeft(s)
	}
}

func rotateMinoRight(mino *Mino) {
	for i, s := range mino.shape {
		mino.shape[i] = rotatePosRight(s)
	}
}

func checkSpace(mino *Mino, field *Field) bool {
	for _, s := range mino.shape {
		target := mino.pos.Add(s)
		if target.y < 0 || target.y >= height || target.x < 0 || target.x >= width {
			return false
		}
		if field.blocks[target.y][target.x].exist {
			return false
		}
	}
	return true
}

func copyMino(m Mino) Mino {
	mino := Mino{
		minoType:  m.minoType,
		pos:       m.pos,
		shape:     []Vector2{},
		direction: m.direction,
		color:     m.color,
	}
	mino.shape = append(mino.shape, m.shape...)
	return mino
}

func canRotateRight(mino Mino, field *Field) (Vector2, bool) {
	tmpMino := copyMino(mino)
	beforeDir := tmpMino.direction
	rotateMinoRight(&tmpMino)
	if checkSpace(&tmpMino, field) {
		return Vector2{0, 0}, true
	}
	return srs(&tmpMino, beforeDir, field)
}

func canRotateLeft(mino Mino, field *Field) (Vector2, bool) {
	tmpMino := copyMino(mino)
	beforeDir := tmpMino.direction
	rotateMinoLeft(&tmpMino)
	if checkSpace(&tmpMino, field) {
		return Vector2{0, 0}, true
	}
	return srs(&tmpMino, beforeDir, field)
}

// https://tetrisch.github.io/main/srs.html
func srs(mino *Mino, beforeDir MinoDirEnum, field *Field) (Vector2, bool) {
	defaultPos := mino.pos
	shift := Vector2{0, 0}
	// 軸を左右に動かす
	switch mino.direction {
	case Right:
		shift = shift.Add(Vector2{-1, 0})
	case Left:
		shift = shift.Add(Vector2{1, 0})
	case Up, Down:
		switch beforeDir {
		case Right:
			shift = shift.Add(Vector2{-1, 0})
		case Left:
			shift = shift.Add(Vector2{1, 0})
		}
	}
	mino.pos = mino.pos.Add(shift)
	if checkSpace(mino, field) {
		return shift, true
	}
	// その状態から軸を上下に動かす
	mino.pos = defaultPos
	switch mino.direction {
	case Right, Left:
		shift = shift.Add(Vector2{0, -1})
	case Up, Down:
		shift = shift.Add(Vector2{0, 1})
	}
	mino.pos = mino.pos.Add(shift)
	if checkSpace(mino, field) {
		return shift, true
	}
	// 元に戻し，軸を上下に2マス動かす
	shift = Vector2{0, 0}
	mino.pos = defaultPos
	switch mino.direction {
	case Right, Left:
		shift = shift.Add(Vector2{0, 2})
	case Up, Down:
		shift = shift.Add(Vector2{0, -2})
	}
	mino.pos = mino.pos.Add(shift)
	if checkSpace(mino, field) {
		return shift, true
	}
	// その状態から軸を左右に動かす
	mino.pos = defaultPos
	switch mino.direction {
	case Right:
		shift = shift.Add(Vector2{-1, 0})
	case Left:
		shift = shift.Add(Vector2{1, 0})
	case Up, Down:
		switch beforeDir {
		case Right:
			shift = shift.Add(Vector2{-1, 0})
		case Left:
			shift = shift.Add(Vector2{1, 0})
		}
	}
	mino.pos = mino.pos.Add(shift)
	if checkSpace(mino, field) {
		return shift, true
	}
	return Vector2{0, 0}, false
}
