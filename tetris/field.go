package tetris

import (
	"image/color"
	"slices"
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

func (f *Field) UpdateMinoFixed() bool {
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
	// ゲームオーバー判定
	return slices.ContainsFunc(f.blocks[playableHeight][3:7], func(b Block) bool { return b.exist })
}
