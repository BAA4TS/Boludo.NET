package util

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func LimpiarConsola() {
	switch runtime.GOOS {
	case "linux", "darwin": // Linux y MacOS
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows": // Windows
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Println("Unsupported platform")
	}
}
