Claro, puedo ayudarte a mejorar el orden y agregar un índice para que la información sea más accesible. Aquí tienes una versión mejorada:

## Índice

1. [Instalación y Uso](#instalación-y-uso)
2. [Configuración](#configuración)
3. [Features](#features)

## Instalación y Uso <a name="instalación-y-uso"></a>

Para comenzar a usar Boludo.NET, sigue estos pasos:

1. **Descargar el proyecto:**
   ```bash
   git clone https://github.com/BAA4TS/Boludo.NET
   ```

2. **Entrar a la carpeta del proyecto:**
   ```bash
   cd Boludo.NET
   ```

3. **Instalar dependencias:**
   ```bash
   go mod tidy
   ```

4. **Ejecutar el servidor:**
   ```bash
   go run .\Server.go
   ```

## Configuración <a name="configuración"></a>

Para configurar Boludo.NET, modifica el archivo `.env` según tus necesidades. Aquí hay un ejemplo de configuración:

```go
// Puerto en el que correrá el servidor
PORT=8001

// Clave que se utilizará para cifrar los mensajes (comandos)
KEY="BAA4TSkey"

// Definición de límites de bots (máximo: 9999)
BOT_LIMIT=5

// Intervalo en segundos para enviar un keep alive a un bot
KEEP_ALIVE=5

// Mensaje de keep alive
KEEP_ALIVE_MSG="check"

// Notificaciones toast del servidor (notifica mediante un toast cuando se conecta y desconecta un bot)
NOTIFY=no
```

Recuerda que los cambios realizados en el archivo `.env` requieren reiniciar el servidor para que surtan efecto.

## Features <a name="features"></a>

Cosas a agregar a futuro

| Server       | BOT                   |
|--------------|-----------------------|
| Menu fix     | Arreglar Receptor     |
| Menu Color   | Screenshot            |
| Cliente Builder | Transferencia de archivos |
