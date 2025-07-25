# Solución de Problemas - LangChain Go

## Error: "chat not found (400)" en Telegram

### Problema

```
Error in node TelegramPublisher: failed to send message to Telegram: telegram: chat not found (400)
```

### Soluciones

#### 1. Verificar que el bot sea administrador del canal

- Ve a tu canal en Telegram
- Ve a Configuración del canal → Administradores
- Agrega tu bot como administrador
- Dale permisos para enviar mensajes

#### 2. Obtener el ID correcto del canal

**Usar la herramienta incluida:**

```bash
go run tools/get_channel_id.go <TU_BOT_TOKEN>
```

**Manual:**

1. Envía un mensaje al canal
2. Visita: `https://api.telegram.org/bot<TU_TOKEN>/getUpdates`
3. Busca el `chat.id` en la respuesta JSON

#### 3. Formato correcto del ID del canal

**Para canales públicos:**

```env
TELEGRAM_CHANNEL_ID=@nombre_del_canal
```

**Para canales privados:**

```env
TELEGRAM_CHANNEL_ID=-1001234567890
```

**Para supergrupos:**

```env
TELEGRAM_CHANNEL_ID=-1001234567890
```

#### 4. Verificar permisos del bot

El bot necesita estos permisos:

- ✅ Enviar mensajes
- ✅ Leer mensajes (para obtener updates)
- ✅ Administrar canal (opcional, pero recomendado)

### Pasos de Verificación

1. **Verificar token del bot:**

```bash
curl "https://api.telegram.org/bot<TU_TOKEN>/getMe"
```

2. **Verificar acceso al canal:**

```bash
curl "https://api.telegram.org/bot<TU_TOKEN>/getChat?chat_id=@tu_canal"
```

3. **Probar envío manual:**

```bash
curl -X POST "https://api.telegram.org/bot<TU_TOKEN>/sendMessage" \
  -H "Content-Type: application/json" \
  -d '{"chat_id":"@tu_canal","text":"Test message"}'
```

## Error: "OPENAI_API_KEY is required"

### Solución

1. Verifica que el archivo `.env` existe en el directorio raíz
2. Asegúrate de que la variable esté correctamente definida:

```env
OPENAI_API_KEY=sk-tu-api-key-aqui
```

3. Verifica que no haya espacios extra o caracteres ocultos

## Error: "Failed to generate text"

### Soluciones

1. **Verificar API key de OpenAI:**

   - Ve a [OpenAI Platform](https://platform.openai.com/api-keys)
   - Confirma que la key sea válida y activa

2. **Verificar créditos:**

   - Ve a [OpenAI Usage](https://platform.openai.com/usage)
   - Confirma que tengas créditos disponibles

3. **Verificar límites de rate:**
   - Espera unos minutos antes de intentar nuevamente
   - Considera usar un modelo diferente

## Error: "No .env file found"

### Solución

1. Crea el archivo `.env` en el directorio raíz del proyecto
2. Copia el contenido de `env.example`
3. Llena con tus valores reales

## Error: "Failed to create bot"

### Solución

1. Verifica que el token del bot sea correcto
2. Confirma que el bot no haya sido eliminado
3. Crea un nuevo bot con @BotFather si es necesario

## Logs de Debug

Para obtener más información de debug, puedes modificar temporalmente el nivel de log en `main.go`:

```go
log.SetFlags(log.LstdFlags | log.Lshortfile)
```

## Verificación Completa

Ejecuta este script para verificar toda la configuración:

```bash
#!/bin/bash
echo "🔍 Verificando configuración..."

# Verificar archivo .env
if [ -f ".env" ]; then
    echo "✅ Archivo .env encontrado"
else
    echo "❌ Archivo .env no encontrado"
    exit 1
fi

# Verificar OpenAI API key
if grep -q "OPENAI_API_KEY=sk-" .env; then
    echo "✅ OpenAI API key configurada"
else
    echo "❌ OpenAI API key no configurada correctamente"
fi

# Verificar Telegram bot token
if grep -q "TELEGRAM_BOT_TOKEN=" .env; then
    echo "✅ Telegram bot token configurado"
else
    echo "❌ Telegram bot token no configurado"
fi

# Verificar canal ID
if grep -q "TELEGRAM_CHANNEL_ID=" .env; then
    echo "✅ Canal ID configurado"
else
    echo "❌ Canal ID no configurado"
fi

echo "🎯 Configuración verificada"
```
