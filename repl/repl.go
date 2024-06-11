package repl

import (
	"fmt"
	"os"

	"BOLUDO.NET/nets"
	"BOLUDO.NET/util"
)

func Repl() {
	var (
		check bool
		input string
	)
	// Imprimir por primera vez el menú
	util.Menu()

	for !check {
		// Entrada principal
		fmt.Print("(Boludo)> ")
		fmt.Scan(&input)

		/* Funciones Principales:
		1 - Limpiar consola
		2 - Listar BOT
		3 - Conectarse con un bot
		4 - Menú de ayuda
		5 - Mostrar estado
		*/

		switch input {
		// Función para limpiar la consola
		case "-cl":
			util.LimpiarConsola()

		// Función para listar los bots
		case "-l":
			nets.Listar()

		// Función para conectarse a un bot
		case "-c":
			nets.ReplBOT()

		// Función para mostrar el menú de ayuda
		case "-h":
			util.Menu()

		// Función para mostrar el estado
		case "-s":
			util.Status()

		// Salida del programa
		case "-e":
			os.Exit(1)
		}
	}
}
