package tetris

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	// 描画用の定数
	blockSize = 30
	// フィールドの描画
	fieldX      = ScreenWidth/2 - blockSize*width/2
	fieldY      = ScreenHeight/2 - blockSize*playableHeight/2
	fieldHeight = blockSize * playableHeight
	fieldWidth  = blockSize * width
)

var (
	// 色
	yellow    = color.RGBA{255, 210, 3, 255}
	lightBlue = color.RGBA{0, 174, 237, 255}
	purple    = color.RGBA{147, 39, 140, 255}
	orange    = color.RGBA{251, 151, 40, 255}
	darkBlue  = color.RGBA{1, 119, 193, 255}
	green     = color.RGBA{122, 193, 65, 255}
	red       = color.RGBA{239, 62, 52, 255}
)

func drawField(screen *ebiten.Image, field Field) {
	for i, row := range field.blocks {
		for j, block := range row {
			vector.DrawFilledRect(screen, fieldX+blockSize*float32(j), fieldY+blockSize*float32(i), blockSize, blockSize, block.color, false)
		}
	}

	for i := 1; i < playableHeight; i++ {
		vector.StrokeLine(screen, fieldX, fieldY+blockSize*float32(i), fieldX+fieldWidth, fieldY+blockSize*float32(i), 1, color.Gray{100}, false)
	}
	for i := 1; i < width; i++ {
		vector.StrokeLine(screen, fieldX+blockSize*float32(i), fieldY, fieldX+blockSize*float32(i), fieldY+fieldHeight, 1, color.Gray{100}, false)
	}
	vector.StrokeRect(screen, fieldX, fieldY, fieldWidth, fieldHeight, 4, color.White, false)
}

func drawGhostBlock(screen *ebiten.Image, pos Vector2, block Block) {
	vector.StrokeRect(screen, fieldX+blockSize*float32(pos.x)+2, fieldY+blockSize*float32(pos.y)+2, blockSize-4, blockSize-4, 4, block.color, false)
}
