package cleanui

import "fmt"

const (
	CLEAR = "\033[H\033[2J\033[3J"
)

func Run() {
	fmt.Print(CLEAR)
}
