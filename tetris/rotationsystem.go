package tetris

func rotatePosLeft(pos Vector2) Vector2 {
	return Vector2{-pos.y, pos.x}
}

func rotatePosRight(pos Vector2) Vector2 {
	return Vector2{pos.y, -pos.x}
}

func rotateMinoLeft(mino *Mino) {
	mino.direction = (mino.direction + 3) % 4
	for i, s := range mino.shape {
		mino.shape[i] = rotatePosLeft(s)
	}
}

func rotateMinoRight(mino *Mino) {
	mino.direction = (mino.direction + 1) % 4
	for i, s := range mino.shape {
		mino.shape[i] = rotatePosRight(s)
	}
}

func rotateIMinoLeft(imino *Mino) {
	imino.direction = (imino.direction + 3) % 4
	rotateIMino(imino)
}

func rotateIMinoRight(imino *Mino) {
	imino.direction = (imino.direction + 1) % 4
	rotateIMino(imino)
}

func rotateIMino(imino *Mino) {
	switch imino.direction {
	case Up:
		imino.shape = []Vector2{{-1, 1}, {0, 1}, {1, 1}, {2, 1}}
	case Right:
		imino.shape = []Vector2{{1, 2}, {1, 1}, {1, 0}, {1, -1}}
	case Down:
		imino.shape = []Vector2{{-1, 0}, {0, 0}, {1, 0}, {2, 0}}
	case Left:
		imino.shape = []Vector2{{0, 2}, {0, 1}, {0, 0}, {0, -1}}
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
	if mino.minoType == OMinoType {
		return Vector2{0, 0}, true
	}
	tmpMino := copyMino(mino)
	beforeDir := tmpMino.direction
	if mino.minoType == IMinoType {
		rotateIMinoRight(&tmpMino)
		if checkSpace(&tmpMino, field) {
			return Vector2{0, 0}, true
		}
		return srsi(&tmpMino, beforeDir, field, true)
	}
	rotateMinoRight(&tmpMino)
	if checkSpace(&tmpMino, field) {
		return Vector2{0, 0}, true
	}
	return srs(&tmpMino, beforeDir, field)
}

func canRotateLeft(mino Mino, field *Field) (Vector2, bool) {
	if mino.minoType == OMinoType {
		return Vector2{0, 0}, true
	}
	tmpMino := copyMino(mino)
	beforeDir := tmpMino.direction
	if mino.minoType == IMinoType {
		rotateIMinoLeft(&tmpMino)
		if checkSpace(&tmpMino, field) {
			return Vector2{0, 0}, true
		}
		return srsi(&tmpMino, beforeDir, field, false)
	}
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
			shift = shift.Add(Vector2{1, 0})
		case Left:
			shift = shift.Add(Vector2{-1, 0})
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
		shift = shift.Add(Vector2{0, 1})
	case Up, Down:
		shift = shift.Add(Vector2{0, -1})
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
		shift = shift.Add(Vector2{0, -2})
	case Up, Down:
		shift = shift.Add(Vector2{0, 2})
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
			shift = shift.Add(Vector2{1, 0})
		case Left:
			shift = shift.Add(Vector2{-1, 0})
		}
	}
	mino.pos = mino.pos.Add(shift)
	if checkSpace(mino, field) {
		return shift, true
	}
	return Vector2{0, 0}, false
}

func srsi(imino *Mino, beforeDir MinoDirEnum, field *Field, rotateRight bool) (Vector2, bool) {
	// なんか違うけど回る。要確認
	// 1. 軸を左右に動かす
	defaultPos := imino.pos
	shift1 := Vector2{0, 0}
	switch imino.direction {
	case Right:
		shift1 = shift1.Add(Vector2{1, 0})
	case Left:
		shift1 = shift1.Add(Vector2{-1, 0})
	case Up:
		switch beforeDir {
		case Right:
			shift1 = shift1.Add(Vector2{-1, 0})
		case Left:
			shift1 = shift1.Add(Vector2{1, 0})
		}
	case Down:
		switch beforeDir {
		case Right:
			shift1 = shift1.Add(Vector2{-1, 0})
		case Left:
			shift1 = shift1.Add(Vector2{1, 0})
		}
	}
	imino.pos = imino.pos.Add(shift1)
	if checkSpace(imino, field) {
		return shift1, true
	}
	// 2. 軸を左右に動かす
	shift2 := Vector2{0, 0}
	imino.pos = defaultPos
	switch imino.direction {
	case Right:
		shift2 = shift2.Add(Vector2{-1, 0})
	case Left:
		shift2 = shift2.Add(Vector2{1, 0})
	case Up:
		switch beforeDir {
		case Right:
			shift2 = shift2.Add(Vector2{2, 0})
		case Left:
			shift2 = shift2.Add(Vector2{-2, 0})
		}
	case Down:
		switch beforeDir {
		case Right:
			shift2 = shift2.Add(Vector2{2, 0})
		case Left:
			shift2 = shift2.Add(Vector2{-2, 0})
		}
	}
	imino.pos = imino.pos.Add(shift2)
	if checkSpace(imino, field) {
		return shift2, true
	}
	// 3. 軸を上下に動かす
	shift3 := Vector2{0, 0}
	imino.pos = defaultPos
	amount := 1
	if !rotateRight {
		amount = 2
	}
	switch imino.direction {
	case Right:
		shift3 = shift1.Add(Vector2{0, -amount})
	case Left:
		shift3 = shift1.Add(Vector2{0, amount})
	case Up, Down:
		switch beforeDir {
		case Right:
			shift3 = shift1.Add(Vector2{0, amount})
		case Left:
			shift3 = shift2.Add(Vector2{0, -amount})
		}

	}
	imino.pos = imino.pos.Add(shift3)
	if checkSpace(imino, field) {
		return shift3, true
	}
	// 4. 軸を上下に動かす
	shift4 := Vector2{0, 0}
	imino.pos = defaultPos
	amount = 1
	if rotateRight {
		amount = 2
	}
	switch imino.direction {
	case Right:
		shift4 = shift2.Add(Vector2{0, amount})
	case Left:
		shift4 = shift2.Add(Vector2{0, -amount})
	case Up, Down:
		switch beforeDir {
		case Right:
			shift4 = shift2.Add(Vector2{0, -amount})
		case Left:
			shift4 = shift1.Add(Vector2{0, amount})
		}

	}
	imino.pos = imino.pos.Add(shift4)
	if checkSpace(imino, field) {
		return shift4, true
	}
	return Vector2{0, 0}, false
}
