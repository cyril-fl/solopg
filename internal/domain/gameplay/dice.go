package gameplay

import "math/rand"

func roll(sides int) int {
	if sides <= 0 {
		return 0
	}
	return 1 + (rand.Intn(sides))
}

func D4() int {
	return roll(4)
}

func D6() int {
	return roll(6)
}

func D8() int {
	return roll(8)
}

func D10() int {
	return roll(10)
}

func D12() int {
	return roll(12)
}

func D20() int {
	return roll(20)
}

func D100() int {
	return roll(100)
}
