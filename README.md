# LangChain Go - Aplicación de Textos Motivacionales

Una aplicación modular en Go que genera textos motivacionales usando OpenAI y los publica en Telegram.

## Arquitectura de Nodos

La aplicación está diseñada con una arquitectura de nodos que permite agregar fácilmente nuevas funcionalidades:

- **TextGeneratorNode**: Genera textos motivacionales usando OpenAI
- **TelegramPublisherNode**: Publica mensajes en un canal de Telegram
- **Pipeline**: Orquesta la ejecución de los nodos en secuencia

## Configuración

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

**Opción 1: Usar la herramienta incluida**

```bash
go run tools/get_channel_id.go <TU_BOT_TOKEN>
```

**Opción 2: Manual**
Para canales públicos:

- Usa el nombre del canal con @ (ej: `@mi_canal`)

Para canales privados:

- Envía un mensaje al canal
- Visita: `https://api.telegram.org/bot<TOKEN>/getUpdates`
- Busca el `chat.id` en la respuesta

### 4. Configurar Variables de Entorno

Crea un archivo `.env` con las siguientes variables:

```env
OPENAI_API_KEY=tu_api_key_de_openai
TELEGRAM_BOT_TOKEN=tu_token_del_bot
TELEGRAM_CHANNEL_ID=@tu_canal_o_id_del_canal
```

### 5. Instalar y Ejecutar

```bash
# Instalar dependencias
go mod tidy

# Ejecutar la aplicación
go run main.go
```

## Agregar Nuevos Nodos

Para agregar un nuevo nodo:

1. Implementa la interfaz `Node` en un nuevo archivo
2. Registra el nodo en el pipeline en `main.go`
3. El nodo se ejecutará automáticamente en la secuencia

### Ejemplo: Agregar un Nodo de Formateo

Para agregar el nodo de formateo de texto, simplemente agrega estas líneas en `main.go`:

```go
// Add text formatter node
textFormatter := nodes.NewTextFormatterNode()
p.AddNode(textFormatter)
```

El pipeline se ejecutará en este orden:

1. TextGenerator → 2. TextFormatter → 3. TelegramPublisher

## Estructura del Proyecto

```
├── main.go                 # Punto de entrada
├── nodes/                  # Directorio de nodos
│   ├── text_generator.go   # Nodo generador de texto
│   ├── text_formatter.go   # Nodo formateador de texto
│   └── telegram_publisher.go # Nodo publicador de Telegram
├── pipeline/               # Lógica del pipeline
│   └── pipeline.go         # Orquestador de nodos
├── config/                 # Configuración
│   └── config.go           # Carga de variables de entorno
└── tools/                  # Herramientas de utilidad
    └── get_channel_id.go   # Obtener ID del canal de Telegram
```
