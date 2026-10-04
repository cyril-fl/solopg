package cleanui

import (
	"os"
	"os/exec"
)

const (
	// TODO LOW Check ka quelle fonctionne et faire une config gene cleatm mode 1 / 2
	// CLEAR = "\033[H\033[2J"
	CLEAR = "clear"
)

func Run() {
	cmd := exec.Command(CLEAR)
	cmd.Stdout = os.Stdout
	cmd.Run()
}
