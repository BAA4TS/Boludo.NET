package nets

import "BOLUDO.NET/config"

// CheckBOT verifica si existe un bot con el ID dado en la lista de bots conectados.
func CheckBOT(ID int) bool {
	return ID < len(config.BOT) && len(config.BOT) > 0
}
