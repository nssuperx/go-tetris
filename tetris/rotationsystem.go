package tetris

func RotatePosLeft(pos Vector2) Vector2 {
	return Vector2{-pos.y, pos.x}
}

func RotatePosRight(pos Vector2) Vector2 {
	return Vector2{pos.y, -pos.x}
}
