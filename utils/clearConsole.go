package utils

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func ClearConsole() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func ClearLastLine() {
	fmt.Print("\033[1A") // Move cursor up one line
	fmt.Print("\033[2K") // Clear entire line
}
