package tetris

import "math/rand/v2"

type MinoTypesEnum int

const (
	IMinoType MinoTypesEnum = iota
	OMinoType
	TMinoType
	SMinoType
	ZMinoType
	JMinoType
	LMinoType
	Empty
)

// ミノ数の2倍の数をバッグに入れておく

const minoTypeCount = 7

type MinoBag struct {
	bag []MinoTypesEnum
}

func newMinoBag() MinoBag {
	minoBag := MinoBag{}
	minoBag.genOneLoop()
	minoBag.genOneLoop() // 14個あればnext7個まで見れる
	return minoBag
}

func (b *MinoBag) genOneLoop() {
	oneLoop := []MinoTypesEnum{}
	for i := range minoTypeCount {
		oneLoop = append(oneLoop, MinoTypesEnum(i))
	}
	rand.Shuffle(len(oneLoop), func(i, j int) {
		oneLoop[i], oneLoop[j] = oneLoop[j], oneLoop[i]
	})
	b.bag = append(b.bag, oneLoop...)
}

func (b *MinoBag) getNextMino() MinoTypesEnum {
	nextMino := b.bag[0]
	b.bag = b.bag[1:]
	if len(b.bag) <= minoTypeCount {
		b.genOneLoop()
	}
	return nextMino
}

func (b *MinoBag) getNextMinos(num int) []MinoTypesEnum {
	next := []MinoTypesEnum{}
	for i := range num {
		next = append(next, b.bag[i])
	}
	return next
}
