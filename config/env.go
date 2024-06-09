package config

import (
	"log"
	"net"
	"os"
	"strconv"
	"sync"

	env "github.com/joho/godotenv"
)

var (
	PORT         string
	KEY          []byte
	BOT_LIMIT    int
	KEEP_ALIVE   int
	BOT          []net.Conn
	Mutex        sync.Mutex
	KeepAliveMSG string
	Notify       bool
)

// init carga las variables de entorno y las configura en las variables globales.
func init() {
	err := env.Load()
	if err != nil {
		log.Fatalf(" [!] Error: %s", err.Error())
	}
	PORT = os.Getenv("PORT")
	KEY = []byte(os.Getenv("KEY"))
	BOT_LIMIT = StringToInt(os.Getenv("BOT_LIMIT"))
	KEEP_ALIVE = StringToInt(os.Getenv("KEEP_ALIVE"))
	KeepAliveMSG = os.Getenv("KEEP_ALIVE_MSG")
	Notify = getNotify()
}

// StringToInt convierte una cadena a un entero y maneja errores.
func StringToInt(input string) int {
	numero, err := strconv.Atoi(input)
	if err != nil {
		log.Fatalf(" [!] Error: %s", err.Error())
		return 0
	}
	return numero
}

// getNotify convierte la variable de entorno NOTIFY a un booleano.
func getNotify() bool {
	check := os.Getenv("NOTIFY")
	return check == "si" || check == "true" || check == "yes" || check == "SI"
}
