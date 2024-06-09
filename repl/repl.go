package repl

import (
	"fmt"
	"os"

	"BOLUDO.NET/util"
)

func Repl() {
	var (
		check bool
		input string
	)
	// Imprimir por primera ves el menu
	util.Menu()

	for !check {
		// Input Principal
		fmt.Print("(Boludo)> ")
		fmt.Scan(&input)

		/* Funciones Principales:
		1 - Limpiar consola
		2 - Listar BOT
		3 - Conectarse con un bot
		4 - Menu de ayuda
		*/

		// Función para limpiar la consola
		if input == "cls" || input == "-cl" || input == "clear" {
			util.LimpiarConsola()

			// Siguiente funcion para listar los bots
		} else if input == "-l" || input == "list" || input == "l" {

			// Siguiente funcion para Conectarse a un bot
		} else if input == "-c" || input == "connect" || input == "-con" {

			// Siguiente funcion para Menu
		} else if input == "-h" || input == "help" || input == "ayuda" {
			util.Menu()
			//Siguiente funcion para mostrar status cuantos bots
		} else if input == "-s" || input == "status" || input == "-st" {
			util.Status()
		}
		// Salida del programa
		if input == "exit" {
			os.Exit(1)
		}
	}
}
