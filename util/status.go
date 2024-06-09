package util

import (
	"fmt"

	"BOLUDO.NET/config"
)

func Status() {
	LimpiarConsola()
	//
	fmt.Println(" --- status --- ")
	botsStatus := fmt.Sprintf("BOTS: %d", len(config.BOT))
	fmt.Println(botsStatus) // Or use botsStatus in another way
}
