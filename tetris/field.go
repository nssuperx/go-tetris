package tetris

import (
	"image/color"
)

const (
	// ロジック的なフィールドのサイズ
	height         = 40
	playableHeight = 20
	width          = 10
)

type Block struct {
	exist      bool
	ghostColor color.RGBA
	color      color.RGBA
}

type Field struct {
	blocks [height][width]Block
}

func (f *Field) clear() {
	f.blocks = [height][width]Block{}
}

func (f *Field) clearLine(lineNum int) {
	// slices.Delete() が使えなかった
	copy(f.blocks[lineNum:], f.blocks[lineNum+1:])
	f.blocks[len(f.blocks)-1] = [width]Block{}
}

func (f *Field) judgeLineClear(lineNum int) bool {
	for _, block := range f.blocks[lineNum] {
		if !block.exist {
			return false
		}
	}
	return true
}

func (f *Field) updateMinoFixed() {
	// TODO: ミノ設置
	// ライン消去
	clearedLinesCount := 0
	clearedLinesNum := []int{}
	for i := range playableHeight {
		if f.judgeLineClear(i) {
			// 1行ずつ消す時に既存の配列がずれるので、消した行の数だけ消す対象の行番号をずらす
			clearedLinesNum = append(clearedLinesNum, i-clearedLinesCount)
			clearedLinesCount++
		}
	}
	for _, n := range clearedLinesNum {
		f.clearLine(n)
	}
}

// 色を塗るだけ
func (f *Field) setBlockColor(mino *Mino) {
	shapePos := convertShapeToPos(mino.shape)
	for _, p := range shapePos {
		// ここで範囲外参照したら移動か回転でミスってる
		f.blocks[mino.pos.y+p.y][mino.pos.x+p.x].color = mino.color
	}
}

func (f *Field) setGhost(mino *Mino, hardDropPos Vector2) {
	for y := range playableHeight {
		for x := range width {
			f.blocks[y][x].ghostColor = color.RGBA{0, 0, 0, 0}
		}
	}
	shapePos := convertShapeToPos(mino.shape)
	for _, p := range shapePos {
		// ここで範囲外参照したら移動か回転でミスってる
		f.blocks[hardDropPos.y+p.y][hardDropPos.x+p.x].ghostColor = mino.color
	}
}

func (f *Field) resetFieldColor() {
	for y := range playableHeight {
		for x := range width {
			if !f.blocks[y][x].exist {
				f.blocks[y][x].color = color.RGBA{0, 0, 0, 0}
			}
		}
	}
}

// ブロックを置けるかどうか
func (f *Field) canSetBlock(mino *Mino, wantDir Vector2) bool {
	shapePos := convertShapeToPos(mino.shape)
	for _, p := range shapePos {
		target := Vector2{mino.pos.x + p.x + wantDir.x, mino.pos.y + p.y + wantDir.y}
		if target.y < 0 || target.y >= height || target.x < 0 || target.x >= width {
			return false
		}
		if f.blocks[target.y][target.x].exist {
			return false
		}
	}
	return true
}

// ブロックを置いて確定する
func (f *Field) setBlock(mino *Mino) {
	shapePos := convertShapeToPos(mino.shape)
	for _, p := range shapePos {
		// ここで範囲外参照したら移動か回転でミスってる
		f.blocks[mino.pos.y+p.y][mino.pos.x+p.x].exist = true
	}
}
