# Automation Chain Go

A modular Go application that automates content generation and publishing using a flexible n8n-like pipeline architecture. It allows you to create customizable automation chains that connect different services and APIs to automate complex workflows.

## 🏗️ Architecture Overview

The application uses a **modular node-based architecture** similar to n8n, allowing you to create reusable and configurable automation chains:

- **Input Nodes**: Capture data from external sources (APIs, databases, files)
- **Processing Nodes**: Transform and analyze data (AI, formatting, validation)
- **Output Nodes**: Publish content to different platforms (social media, APIs)
- **Pipeline**: Orchestrates sequential node execution
- **PipelineBuilder**: Constructs pipelines from JSON configuration

## 🚀 Quick Start

### 1. Clone and Setup

```bash
git clone <repository-url>
cd automation-chain
go mod tidy
```

### 2. Configure Credentials

Copy the example file and create your own credentials:

```bash
cp config/credentials_example.json config/credentials.json
```

Then edit `config/credentials.json` with your actual API keys:

```json
{
  "openai": {
    "default": {
      "api_key": "sk-your-openai-api-key"
    }
  },
  "telegram": {
    "motivational_bot": {
      "token": "your-telegram-bot-token",
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
  "description": "Generates and publishes motivational texts to Telegram",
  "schedule": "0 9 * * *",
  "nodes": [
    {
      "id": "text_generator",
      "type": "text_generator",
      "name": "Generate Motivational Text",
      "credentials": "default",
      "config": {
        "model": "gpt-3.5-turbo",
        "prompt_template": "Generate a short and powerful motivational text in Spanish...",
        "max_tokens": 300,
        "temperature": 0.8
      }
    },
    {
      "id": "telegram_publisher",
      "type": "telegram_publisher", 
      "name": "Publish to Telegram",
      "credentials": "motivational_bot",
      "config": {
        "message_template": "💪 *Daily Motivational Message*\n\n%s\n\n✨ Have an amazing day!"
      }
    }
  ]
}
```

### 4. Run the Application

```bash
go run main.go
```

The application will automatically load the pipeline configuration and execute it. You can modify the pipeline name in `main.go` to run different pipelines.

### Available Pipelines

- **telegram**: Generates motivational content and publishes to Telegram
- **telegram_news**: Generates news content and publishes to Telegram  
- **multi_telegram**: Publishes to multiple Telegram channels

To run a different pipeline, modify line 20 in `main.go`:
```go
pipelineConfig, err := loadPipelineConfig("telegram") // Change "telegram" to desired pipeline
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
5. Make sure the bot has permission to send messages

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
- Find the `chat.id` in the response (it will be a negative number)

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
│   ├── credentials.json      # API keys and tokens (create this file)
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
      "token": "value",
      "channel_id": "value",
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
      "token": "token1",
      "channel_id": "@motivational"
    },
    "news_bot": {
      "token": "token2", 
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
- Ensure the bot has permission to send messages
- For private channels, make sure the bot is a member of the channel

**"OpenAI API error"**
- Verify your API key is correct and has sufficient credits
- Check that the model name is valid (e.g., "gpt-3.5-turbo", "gpt-4")
- Ensure your OpenAI account has access to the specified model

For more detailed troubleshooting, see [TROUBLESHOOTING.md](TROUBLESHOOTING.md).
