package core

import (
	"os/exec"
)

// CMD ejecuta un comando CMD dado y retorna su salida como una cadena.
func CMD(Command string) (string, error) {
	cmd := exec.Command("cmd", "/C", Command)
	Salida, err := cmd.Output()
	if err != nil {
		return "", err
	} else {
		return string(Salida), nil
	}
}
