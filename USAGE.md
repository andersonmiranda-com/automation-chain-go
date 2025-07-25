# Guía de Uso - Aplicación de Textos Motivacionales

## Configuración Inicial

### 1. Configurar OpenAI

1. Ve a [OpenAI API](https://platform.openai.com/api-keys)
2. Crea una nueva API key
3. Copia la key (comienza con `sk-`)

### 2. Configurar Telegram Bot

1. Habla con [@BotFather](https://t.me/botfather) en Telegram
2. Crea un nuevo bot con `/newbot`
3. Guarda el token del bot
4. Agrega el bot a tu canal como administrador

### 3. Obtener ID del Canal

Para canales públicos:

- Usa el nombre del canal con @ (ej: `@mi_canal`)

Para canales privados:

- Envía un mensaje al canal
- Visita: `https://api.telegram.org/bot<TOKEN>/getUpdates`
- Busca el `chat.id` en la respuesta

### 4. Configurar Variables de Entorno

Crea un archivo `.env` basado en `env.example`:

```env
OPENAI_API_KEY=sk-tu-api-key-aqui
TELEGRAM_BOT_TOKEN=tu-token-del-bot-aqui
TELEGRAM_CHANNEL_ID=@tu_canal_o_id_aqui
```

## Ejecución

### Ejecutar la Aplicación

```bash
# Instalar dependencias
go mod tidy

# Compilar
go build -o motivational-app

# Ejecutar
./motivational-app
```

O ejecutar directamente:

```bash
go run main.go
```

### Programar Ejecución Automática

Para ejecutar diariamente, puedes usar cron:

```bash
# Editar crontab
crontab -e

# Agregar esta línea para ejecutar todos los días a las 8:00 AM
0 8 * * * cd /ruta/a/tu/proyecto && ./motivational-app
```

## Personalización

### Modificar el Prompt de Motivación

Edita `nodes/text_generator.go` y modifica la variable `prompt`:

```go
prompt := `Tu prompt personalizado aquí...`
```

### Agregar Nuevos Nodos

1. Crea un nuevo archivo en `nodes/`
2. Implementa la interfaz `Node`
3. Agrega el nodo en `main.go`

### Ejemplo de Nodo Personalizado

```go
package nodes

import (
    "context"
    "fmt"
    "log"
)

type MyCustomNode struct{}

func NewMyCustomNode() *MyCustomNode {
    return &MyCustomNode{}
}

func (n *MyCustomNode) Name() string {
    return "MyCustomNode"
}

func (n *MyCustomNode) Execute(ctx context.Context, input map[string]interface{}) (map[string]interface{}, error) {
    log.Println("Executing custom node...")

    // Tu lógica aquí

    return map[string]interface{}{
        "custom_result": "valor",
    }, nil
}
```

Luego agrégalo en `main.go`:

```go
myNode := nodes.NewMyCustomNode()
p.AddNode(myNode)
```

## Solución de Problemas

### Error: "OPENAI_API_KEY is required"

- Verifica que el archivo `.env` existe
- Asegúrate de que la variable esté correctamente definida

### Error: "TELEGRAM_BOT_TOKEN is required"

- Verifica que el token del bot sea correcto
- Asegúrate de que el bot esté agregado al canal

### Error: "Failed to send message to Telegram"

- Verifica que el bot sea administrador del canal
- Confirma que el ID del canal sea correcto
- Para canales privados, usa el ID numérico en lugar del nombre

### Error: "Failed to generate text"

- Verifica que tu API key de OpenAI sea válida
- Confirma que tengas créditos disponibles en tu cuenta de OpenAI

## Logs y Debugging

La aplicación muestra logs detallados durante la ejecución:

```
🚀 Starting LangChain Go - Motivational Text Generator
✅ Configuration loaded successfully
Added node: TextGenerator
Added node: TelegramPublisher
📋 Pipeline configured with 2 nodes
Starting pipeline execution...
Executing node 1/2: TextGenerator
Generating motivational text...
Generated text: [texto generado]
Node TextGenerator completed successfully
Executing node 2/2: TelegramPublisher
Publishing to Telegram...
Message published successfully to channel: @tu_canal
Node TelegramPublisher completed successfully
Pipeline execution completed successfully!
🎉 Application completed successfully!
```
