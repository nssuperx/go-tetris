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
	exist bool
	color color.RGBA
}

type Field struct {
	blocks [height][width]Block
}

func (f *Field) clearLine(lineNum int) {
	// slices.Delete() が使えなかった
	copy(f.blocks[lineNum:], f.blocks[lineNum+1:])
	f.blocks[len(f.blocks)-1] = [10]Block{}
}

func (f *Field) judgeLineClear(lineNum int) bool {
	for _, block := range f.blocks[lineNum] {
		if !block.exist {
			return false
		}
	}
	return true
}

func (f *Field) UpdateMinoFixed() {
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
func (f *Field) SetBlockColor(mino *Mino) {
	for _, s := range mino.shape {
		// ここで範囲外参照したら移動か回転でミスってる
		f.blocks[mino.pos.y+s.y][mino.pos.x+s.x].color = mino.color
	}
}

// ブロックを置けるかどうか
func (f *Field) CanSetBlock(mino *Mino) bool {
	for _, s := range mino.shape {
		// ここで範囲外参照したら移動か回転でミスってる
		if f.blocks[mino.pos.y+s.y][mino.pos.x+s.x].exist {
			return false
		}
	}
	return true
}

// ブロックを置いて確定する
func (f *Field) SetBlock(mino *Mino) {
	for _, s := range mino.shape {
		// ここで範囲外参照したら移動か回転でミスってる
		f.blocks[mino.pos.y+s.y][mino.pos.x+s.x].exist = true
	}
}
