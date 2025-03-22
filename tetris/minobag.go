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
)

// ミノ数の2倍の数をバッグに入れておく

const minoTypeCount = 7

type MinoBag struct {
	bag []MinoTypesEnum
}

func NewMinoBag() MinoBag {
	minoBag := MinoBag{}
	minoBag.GenOneLoop()
	minoBag.GenOneLoop() // 14個あればnext7個まで見れる
	return minoBag
}

func (b *MinoBag) GenOneLoop() {
	oneLoop := []MinoTypesEnum{}
	for i := range minoTypeCount {
		oneLoop = append(oneLoop, MinoTypesEnum(i))
	}
	rand.Shuffle(len(oneLoop), func(i, j int) {
		oneLoop[i], oneLoop[j] = oneLoop[j], oneLoop[i]
	})
	b.bag = append(b.bag, oneLoop...)
}

func (b *MinoBag) GetNextMino() MinoTypesEnum {
	nextMino := b.bag[0]
	b.bag = b.bag[1:]
	if len(b.bag) <= minoTypeCount {
		b.GenOneLoop()
	}
	return nextMino
}

func (b *MinoBag) GetNextMinos(num int) []MinoTypesEnum {
	next := []MinoTypesEnum{}
	for i := range num {
		next = append(next, b.bag[i])
	}
	return next
}
