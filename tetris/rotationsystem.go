package tetris

type RotationSystem interface {
	canRotate(mino *Mino, field *Field, rotateRight bool) (Vector2, bool)
}

type OMinoRotationSystem struct{}

func (s *OMinoRotationSystem) canRotate(mino *Mino, field *Field, rotateRight bool) (Vector2, bool) {
	return Vector2{0, 0}, true
}

type CommonMinoRotationSystem struct{}

func (s *CommonMinoRotationSystem) canRotate(mino *Mino, field *Field, rotateRight bool) (Vector2, bool) {
	tmpMino := copyMino(mino)
	beforeDir := tmpMino.direction
	if rotateRight {
		rotateMinoRight(&tmpMino)
	} else {
		rotateMinoLeft(&tmpMino)
	}
	if checkSpace(&tmpMino, field) {
		return Vector2{0, 0}, true
	}
	return srs(&tmpMino, beforeDir, field)
}

type IMinoRotationSystem struct{}

func (s *IMinoRotationSystem) canRotate(mino *Mino, field *Field, rotateRight bool) (Vector2, bool) {
	tmpMino := copyMino(mino)
	beforeDir := tmpMino.direction
	if rotateRight {
		rotateMinoRight(&tmpMino)
	} else {
		rotateMinoLeft(&tmpMino)
	}
	if checkSpace(&tmpMino, field) {
		return Vector2{0, 0}, true
	}
	return srsi(&tmpMino, beforeDir, field, rotateRight)
}

// NxNの行列を半時計回り90度回転
func rotateLeftInplace[T any](m [][]T) {
	N := len(m)
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			m[i][j], m[j][i] = m[j][i], m[i][j]
		}
	}
	for j := 0; j < N; j++ {
		for i := 0; i < N/2; i++ {
			m[i][j], m[N-1-i][j] = m[N-1-i][j], m[i][j]
		}
	}
}

// NxNの行列を時計回り90度回転
func rotateRightInplace[T any](m [][]T) {
	N := len(m)
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			m[i][j], m[j][i] = m[j][i], m[i][j]
		}
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N/2; j++ {
			m[i][j], m[i][N-1-j] = m[i][N-1-j], m[i][j]
		}
	}
}

func rotateMinoLeft(mino *Mino) {
	mino.direction = (mino.direction + 3) % 4
	rotateLeftInplace(mino.shape)
}

func rotateMinoRight(mino *Mino) {
	mino.direction = (mino.direction + 1) % 4
	rotateRightInplace(mino.shape)
}

func checkSpace(mino *Mino, field *Field) bool {
	shapePos := convertShapeToPos(mino.shape)
	for _, p := range shapePos {
		target := mino.pos.Add(p)
		if target.y < 0 || target.y >= height || target.x < 0 || target.x >= width {
			return false
		}
		if field.blocks[target.y][target.x].exist {
			return false
		}
	}
	return true
}

func copyMino(m *Mino) Mino {
	mino := Mino{
		pos:       m.pos,
		shape:     [][]bool{},
		direction: m.direction,
	}
	for _, row := range m.shape {
		newRow := make([]bool, len(row))
		copy(newRow, row)
		mino.shape = append(mino.shape, newRow)
	}
	return mino
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
