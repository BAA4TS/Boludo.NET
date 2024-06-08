package repl

import (
	"fmt"
)

func Repl() {
	var (
		check = false
		input string
	)
	for !check {
		fmt.Print("\n(Boludo)> ")
		fmt.Scan(&input)
	}
}
