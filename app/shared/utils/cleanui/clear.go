package cleanui

import (
	"os"
	"os/exec"
)

func Run() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
