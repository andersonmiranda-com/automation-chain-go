# LangChain Go - Motivational Content Pipeline

A modular Go application that generates motivational content using OpenAI and publishes it to Telegram channels using a flexible, n8n-like pipeline architecture.

## 🏗️ Architecture Overview

The application uses a **modular node-based architecture** similar to n8n, allowing you to create reusable and configurable pipelines:

- **TextGeneratorNode**: Generates motivational content using OpenAI
- **TelegramPublisherNode**: Publishes messages to Telegram channels
- **Pipeline**: Orchestrates node execution in sequence
- **PipelineBuilder**: Constructs pipelines from JSON configuration

## 🚀 Quick Start

### 1. Clone and Setup

```bash
git clone <repository-url>
cd automation-chain
go mod tidy
```

### 2. Configure Credentials

Create `config/credentials.json` with your API keys:

```json
{
  "openai": {
    "default": {
      "api_key": "sk-your-openai-api-key"
    }
  },
  "telegram": {
    "motivational_bot": {
      "bot_token": "your-telegram-bot-token",
      "channel_id": "@your_channel_or_channel_id"
    }
  }
}
```

### 3. Configure Pipeline

The pipeline configuration is in `config/pipelines/telegram.json`:

```json
{
  "name": "telegram_pipeline",
  "description": "Generate and publish motivational content to Telegram",
  "nodes": [
    {
      "id": "text_generator",
      "type": "text_generator",
      "name": "Generate Motivational Text",
      "credentials": "default",
      "config": {
        "prompt_template": "Generate a motivational message in Spanish...",
        "max_tokens": 150
      }
    },
    {
      "id": "telegram_publisher",
      "type": "telegram_publisher", 
      "name": "Publish to Telegram",
      "credentials": "motivational_bot",
      "config": {
        "message_format": "text"
      }
    }
  ]
}
```

### 4. Run the Application

```bash
go run main.go
```

## 🔧 Configuration Guide

### OpenAI Setup

1. Go to [OpenAI API](https://platform.openai.com/api-keys)
2. Create a new API key
3. Copy the key (starts with `sk-`)
4. Add it to `config/credentials.json` under `openai.default.api_key`

### Telegram Bot Setup

1. Talk to [@BotFather](https://t.me/botfather) on Telegram
2. Create a new bot with `/newbot`
3. Save the bot token
4. Add the bot to your channel as administrator

### Get Channel ID

**Option 1: Use the included tool**

```bash
go run tools/get_channel_id.go <YOUR_BOT_TOKEN>
```

**Option 2: Manual method**

For public channels:
- Use the channel name with @ (e.g., `@my_channel`)

For private channels:
- Send a message to the channel
- Visit: `https://api.telegram.org/bot<TOKEN>/getUpdates`
- Find the `chat.id` in the response

## 📁 Project Structure

```
├── main.go                    # Application entry point
├── nodes/                     # Node implementations
│   ├── base/                  # Base interfaces and types
│   │   └── node.go           # Node interface definition
│   ├── ai/                   # AI-related nodes
│   │   └── text_generator.go # OpenAI text generation
│   └── publishers/           # Publishing nodes
│       └── telegram_publisher.go # Telegram publishing
├── pipelines/                # Pipeline orchestration
│   └── base/                 # Base pipeline components
│       ├── pipeline.go       # Pipeline execution logic
│       └── builder.go        # Pipeline construction
├── services/                 # External service clients
│   ├── openai.go            # OpenAI API client
│   └── telegram.go          # Telegram Bot API client
├── config/                   # Configuration files
│   ├── credentials.json      # API keys and tokens
│   ├── credentials_example.json # Example credentials structure
│   ├── credentials_test.json # Test credentials
│   └── pipelines/           # Pipeline definitions
│       ├── telegram.json    # Telegram motivational pipeline
│       ├── telegram_news.json # Telegram news pipeline
│       └── multi_telegram.json # Multi-channel example
├── tools/                    # Utility tools
│   └── get_channel_id.go    # Telegram channel ID finder
├── tests/                    # Test files
│   └── pipeline_test.go     # Pipeline unit tests
└── docs/                     # Documentation
    ├── ROADMAP.md           # Development roadmap
    ├── FASE1_COMPLETADA.md  # Phase 1 completion report
    └── NODES_DOCUMENTATION.md # Node configuration guide
```

## 🔐 Credential Management

The application uses a **reference-based credential system**:

### Credential Structure

```json
{
  "service_name": {
    "credential_name": {
      "api_key": "value",
      "other_param": "value"
    }
  }
}
```

### Pipeline References

Each node in a pipeline references credentials by name:

```json
{
  "id": "text_generator",
  "type": "text_generator",
  "credentials": "default",  // References openai.default
  "config": { ... }
}
```

### Multiple Credentials Example

You can have multiple credentials per service:

```json
{
  "telegram": {
    "motivational_bot": {
      "bot_token": "token1",
      "channel_id": "@motivational"
    },
    "news_bot": {
      "bot_token": "token2", 
      "channel_id": "@news"
    }
  }
}
```

## 🧪 Testing

Run the test suite:

```bash
go test ./...
```

Run specific tests:

```bash
go test ./tests/
```

## 📚 Adding New Nodes

To add a new node:

1. **Create the node implementation** in `nodes/` directory
2. **Implement the Node interface**:
   ```go
   type Node interface {
       Execute(ctx context.Context, input map[string]interface{}) (map[string]interface{}, error)
       Validate() error
   }
   ```
3. **Register the node** in `pipelines/base/builder.go`
4. **Add configuration** to your pipeline JSON

### Example: Adding a Text Formatter Node

```go
// In nodes/formatters/text_formatter.go
type TextFormatterNode struct {
    config base.NodeConfig
}

func (n *TextFormatterNode) Execute(ctx context.Context, input map[string]interface{}) (map[string]interface{}, error) {
    text := input["text"].(string)
    formatted := strings.ToUpper(text)
    return map[string]interface{}{"text": formatted}, nil
}
```

Then register it in the builder:

```go
case "text_formatter":
    return formatters.NewTextFormatterNode(nodeConfig)
```

## 🛣️ Development Roadmap

See [ROADMAP.md](docs/ROADMAP.md) for detailed development phases:

- ✅ **Phase 1**: Basic pipeline with OpenAI and Telegram (COMPLETED)
- 🔄 **Phase 2**: Input nodes (TopicSelector, DataFetcher)
- 📋 **Phase 3**: AI nodes (ImageGenerator, ContentAnalyzer)
- 📋 **Phase 4**: Media nodes (ImageUploader, ImageProcessor)
- 📋 **Phase 5**: Publisher nodes (LinkedIn, Instagram, Twitter)
- 📋 **Phase 6**: Google Services (Sheets, Drive)
- 📋 **Phase 7**: Specialized pipelines
- 📋 **Phase 8**: Scheduling system
- 📋 **Phase 9**: Logging and monitoring
- 📋 **Phase 10**: Documentation and optimization

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🆘 Troubleshooting

### Common Issues

**"Failed to load credentials"**
- Ensure `config/credentials.json` exists and is valid JSON
- Check that credential names match between pipeline and credentials files

**"Telegram channel not found"**
- Verify the bot is added to the channel as administrator
- Check that the channel ID is correct (use `tools/get_channel_id.go`)

**"OpenAI API error"**
- Verify your API key is correct and has sufficient credits
- Check that the model name is valid

For more detailed troubleshooting, see [TROUBLESHOOTING.md](TROUBLESHOOTING.md).
