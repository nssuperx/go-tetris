package tetris

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// 入力周りを吸収する

func downPressed() bool {
	return ebiten.IsKeyPressed(ebiten.KeyS)
}

func upPressed() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyW)
}

func leftJustPressed() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyA)
}

func leftPressed() bool {
	return ebiten.IsKeyPressed(ebiten.KeyA)
}

func rightJustPressed() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyD)
}

func rightPressed() bool {
	return ebiten.IsKeyPressed(ebiten.KeyD)
}

func rotateLeftPressed() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyJ)
}

func rotateRightPressed() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyI)
}

func holdPressed() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyO)
}

func startPressed() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyEnter)
}
