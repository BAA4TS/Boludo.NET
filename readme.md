# Boludo.NET

Este es un proyecto simple, una BOT-NET desarrollada en el lenguaje de programación Go.

[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)
[![Uruguay](https://img.shields.io/badge/Uruguay-blue.svg)](https://es.wikipedia.org/wiki/Uruguay)
[![Malware](https://img.shields.io/badge/Malware-red.svg)](https://en.wikipedia.org/wiki/Malware)
[![Golang](https://img.shields.io/badge/Golang-00ADD8?logo=go&logoColor=white)](https://golang.org/)
[![Botnet](https://img.shields.io/badge/Botnet-FFA500)](https://en.wikipedia.org/wiki/Botnet)
[![Spyware](https://img.shields.io/badge/Spyware-8A2BE2)](https://en.wikipedia.org/wiki/Spyware)

## Cómo Usarlo

Antes de comenzar, configura el archivo `.env` con las siguientes variables:

```plaintext
PORT=8001        # Puerto del servidor
KEY="BAA4TSkey"  # Clave de cifrado de conexión
BOT_LIMIT=5      # Límite de bots que se conectan a la botnet
KEEP_ALIVE=5     # Intervalo en el que se envía un keep alive al bot
```

Para desplegar este proyecto, asegúrate de tener Go instalado en tu sistema.

### Requisitos:
- [Go](https://golang.org/) instalado

### Opción 1: Clonar el proyecto desde Git

1. Clona el proyecto desde Git utilizando el siguiente comando:

```bash
git clone https://github.com/BAA4TS/Boludo.NET
```

2. Ingresa al directorio del proyecto:

```bash
cd Boludo.NET
```

3. Ejecuta el comando `go mod tidy` para garantizar que las dependencias estén actualizadas y limpias:

```bash
go mod tidy
```

4. Después de actualizar las dependencias, puedes ejecutar el proyecto utilizando `go run`:

```bash
go run Server.go
```

### Opción 2: Descargar la versión compilada o compilarla tú mismo

1. Puedes descargar la versión compilada desde las releases del proyecto o compilarla tú mismo descargando el código fuente desde el repositorio Git.

2. Si descargas el código fuente, asegúrate de ingresar al directorio del proyecto.

3. Compila el proyecto utilizando el comando `go build`:

```bash
go build Server.go
```

Esto generará un archivo ejecutable del servidor.

4. Una vez compilado, puedes ejecutar el servidor:

```bash
./Server
```

Con cualquiera de estas opciones, podrás desplegar el proyecto en tu entorno local. Recuerda que si optas por clonar el proyecto, es importante ejecutar `go mod tidy` para asegurarte de tener todas las dependencias necesarias antes de ejecutar el servidor.
