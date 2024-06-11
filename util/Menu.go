package util

import "fmt"

func Menu() {
	LimpiarConsola()
	fmt.Println(" --- Menú --- Github: BAA4TS --- ")
	fmt.Println("(-l) Listar los bots conectados")
	fmt.Println("(-c) Establecer una sesión PowerShell con el bot indicado")
	fmt.Println("(-s) Mostrar un estado de cuántos bots conectados hay, tiempo de conexión, etc.")
	fmt.Println("(-cl) Limpiar la consola")
	fmt.Println("(-h) Menú de ayuda")
	fmt.Println("(-e) Salir del programa")
	fmt.Print("\n\n")
}
