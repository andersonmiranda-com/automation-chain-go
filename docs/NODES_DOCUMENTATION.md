# Nodes Documentation

## 🎯 Overview
This document describes all available nodes in the pipeline system, their configuration parameters, and usage examples.

## 📋 Node Categories

### 🤖 AI Nodes (`nodes/ai/`)
Nodes that interact with AI services like OpenAI.

### 📤 Publisher Nodes (`nodes/publishers/`)
Nodes that publish content to external platforms.

### 📥 Input Nodes (`nodes/input/`)
Nodes that fetch or generate input data.

### 🎨 Media Nodes (`nodes/media/`)
Nodes that handle media processing and generation.

### 🔧 Utility Nodes (`nodes/utility/`)
Nodes that perform data transformation and formatting.

---

## 🤖 AI Nodes

### TextGeneratorNode

**Purpose**: Generates text content using OpenAI's language models.

**Type**: `text_generator`

**Location**: `nodes/ai/text_generator.go`

#### Configuration Parameters

| Parameter | Type | Required | Default | Description |
|-----------|------|----------|---------|-------------|
| `model` | string | Yes | - | OpenAI model to use (e.g., "gpt-3.5-turbo", "gpt-4") |
| `prompt_template` | string | Yes | - | Template for the prompt to send to OpenAI |
| `max_tokens` | int | No | 300 | Maximum number of tokens to generate |
| `temperature` | float | No | 0.7 | Controls randomness (0.0 = deterministic, 1.0 = very random) |
| `top_p` | float | No | 1.0 | Controls diversity via nucleus sampling |
| `frequency_penalty` | float | No | 0.0 | Reduces repetition of similar content |
| `presence_penalty` | float | No | 0.0 | Reduces repetition of topics |

#### Input
- `topic` (string, optional): Topic to include in the prompt
- `style` (string, optional): Writing style to apply
- `length` (string, optional): Desired length ("short", "medium", "long")

#### Output
- `generated_text` (string): The generated text content
- `tokens_used` (int): Number of tokens consumed
- `model_used` (string): Model that was used for generation

#### Example Configuration
```json
{
  "id": "text_generator",
  "type": "text_generator",
  "name": "Generate Motivational Text",
  "config": {
    "model": "gpt-3.5-turbo",
    "prompt_template": "Generate a motivational text in Spanish about {topic}. Style: {style}. Length: {length}.",
    "max_tokens": 300,
    "temperature": 0.8
  }
}
```

#### Example Usage
```go
// Input to the node
input := map[string]interface{}{
    "topic": "success",
    "style": "inspirational",
    "length": "medium",
}

// Output from the node
output := map[string]interface{}{
    "generated_text": "Success is not final, failure is not fatal...",
    "tokens_used": 150,
    "model_used": "gpt-3.5-turbo",
}
```

---

## 📤 Publisher Nodes

### TelegramPublisherNode

**Purpose**: Publishes messages to Telegram channels or groups.

**Type**: `telegram_publisher`

**Location**: `nodes/publishers/telegram_publisher.go`

#### Configuration Parameters

| Parameter | Type | Required | Default | Description |
|-----------|------|----------|---------|-------------|
| `message_template` | string | Yes | - | Template for the message with placeholders |
| `parse_mode` | string | No | "Markdown" | Message parsing mode ("Markdown", "HTML", "None") |
| `disable_web_page_preview` | bool | No | false | Disable link previews |
| `disable_notification` | bool | No | false | Send silently |
| `reply_to_message_id` | int | No | - | Reply to specific message |

#### Input
- `text` (string, required): Text content to publish
- `image_url` (string, optional): URL of image to include
- `caption` (string, optional): Caption for the image

#### Output
- `message_id` (int): ID of the sent message
| `chat_id` (string): ID of the target chat/channel
- `sent_at` (string): Timestamp when message was sent
- `success` (bool): Whether the message was sent successfully

#### Example Configuration
```json
{
  "id": "telegram_publisher",
  "type": "telegram_publisher",
  "name": "Publish to Telegram",
  "config": {
    "message_template": "💪 *Daily Motivation*\n\n{text}\n\n✨ Have an amazing day!",
    "parse_mode": "Markdown",
    "disable_web_page_preview": true
  }
}
```

#### Example Usage
```go
// Input to the node
input := map[string]interface{}{
    "text": "Success is not final, failure is not fatal...",
}

// Output from the node
output := map[string]interface{}{
    "message_id": 12345,
    "chat_id": "-1002876849256",
    "sent_at": "2025-07-25T20:03:48Z",
    "success": true,
}
```

---

## 📥 Input Nodes (Planned)

### TopicSelectorNode

**Purpose**: Selects topics from predefined lists or external sources.

**Type**: `topic_selector`

**Location**: `nodes/input/topic_selector.go` (to be implemented)

#### Configuration Parameters

| Parameter | Type | Required | Default | Description |
|-----------|------|----------|---------|-------------|
| `topic_list` | array | Yes | - | List of available topics |
| `selection_method` | string | No | "random" | How to select ("random", "sequential", "weighted") |
| `exclude_used` | bool | No | true | Exclude recently used topics |
| `max_history` | int | No | 10 | Number of recent topics to remember |

#### Input
- `category` (string, optional): Topic category to filter by
- `mood` (string, optional): Mood to match topics with

#### Output
- `selected_topic` (string): The selected topic
- `category` (string): Topic category
- `tags` (array): Associated tags

### DataFetcherNode

**Purpose**: Fetches data from external APIs or databases.

**Type**: `data_fetcher`

**Location**: `nodes/input/data_fetcher.go` (to be implemented)

#### Configuration Parameters

| Parameter | Type | Required | Default | Description |
|-----------|------|----------|---------|-------------|
| `api_url` | string | Yes | - | URL to fetch data from |
| `method` | string | No | "GET" | HTTP method to use |
| `headers` | object | No | {} | HTTP headers to include |
| `timeout` | int | No | 30 | Request timeout in seconds |
| `retry_attempts` | int | No | 3 | Number of retry attempts |

#### Input
- `query_params` (object, optional): Query parameters to include
- `body` (object, optional): Request body for POST/PUT requests

#### Output
- `data` (object): Fetched data
- `status_code` (int): HTTP status code
- `response_time` (float): Response time in seconds

### FileReaderNode

**Purpose**: Reads data from files (JSON, CSV, TXT).

**Type**: `file_reader`

**Location**: `nodes/input/file_reader.go` (to be implemented)

#### Configuration Parameters

| Parameter | Type | Required | Default | Description |
|-----------|------|----------|---------|-------------|
| `file_path` | string | Yes | - | Path to the file to read |
| `file_type` | string | Yes | - | Type of file ("json", "csv", "txt") |
| `encoding` | string | No | "utf-8" | File encoding |
| `delimiter` | string | No | "," | CSV delimiter (for CSV files) |

#### Input
- `line_number` (int, optional): Specific line to read (for TXT files)
- `key_path` (string, optional): JSON path to extract specific data

#### Output
- `content` (object/string): File content
- `file_size` (int): File size in bytes
- `lines_count` (int): Number of lines (for text files)

---

## 🎨 Media Nodes (Planned)

### ImageGeneratorNode

**Purpose**: Generates images using AI services like DALL-E or Stable Diffusion.

**Type**: `image_generator`

**Location**: `nodes/media/image_generator.go` (to be implemented)

#### Configuration Parameters

| Parameter | Type | Required | Default | Description |
|-----------|------|----------|---------|-------------|
| `model` | string | Yes | - | AI model to use ("dall-e-3", "stable-diffusion") |
| `size` | string | No | "1024x1024" | Image dimensions |
| `quality` | string | No | "standard" | Image quality ("standard", "hd") |
| `style` | string | No | "natural" | Artistic style |
| `format` | string | No | "png" | Output format |

#### Input
- `prompt` (string, required): Description of the image to generate
- `negative_prompt` (string, optional): What to avoid in the image

#### Output
- `image_url` (string): URL of the generated image
- `image_path` (string): Local path to saved image
- `generation_time` (float): Time taken to generate

### ImageUploaderNode

**Purpose**: Uploads images to cloud storage services.

**Type**: `image_uploader`

**Location**: `nodes/media/image_uploader.go` (to be implemented)

#### Configuration Parameters

| Parameter | Type | Required | Default | Description |
|-----------|------|----------|---------|-------------|
| `service` | string | Yes | - | Service to use ("cloudinary", "aws-s3", "google-cloud") |
| `folder` | string | No | "uploads" | Target folder |
| `public_id` | string | No | - | Custom public ID |
| `transformation` | object | No | {} | Image transformations |

#### Input
- `image_path` (string, required): Path to the image file
- `tags` (array, optional): Tags for the uploaded image

#### Output
- `public_url` (string): Public URL of uploaded image
- `asset_id` (string): Asset identifier
- `file_size` (int): File size in bytes

---

## 🔧 Utility Nodes

### TextFormatterNode

**Purpose**: Formats and processes text content.

**Type**: `text_formatter`

**Location**: `nodes/utility/text_formatter.go` (to be implemented)

#### Configuration Parameters

| Parameter | Type | Required | Default | Description |
|-----------|------|----------|---------|-------------|
| `trim_spaces` | bool | No | true | Remove extra spaces |
| `normalize_line_breaks` | bool | No | true | Standardize line breaks |
| `add_punctuation` | bool | No | false | Add missing punctuation |
| `max_length` | int | No | - | Maximum text length |
| `case` | string | No | "preserve" | Text case ("lower", "upper", "title", "preserve") |

#### Input
- `text` (string, required): Text to format
- `style` (string, optional): Formatting style

#### Output
- `formatted_text` (string): Processed text
- `changes_made` (array): List of changes applied
- `original_length` (int): Original text length
- `final_length` (int): Final text length

---

## 🔐 Credential Requirements

Each node type requires specific credentials to be configured in the pipeline:

### AI Nodes
- **TextGeneratorNode**: `openai.{credential_name}`

### Publisher Nodes
- **TelegramPublisherNode**: `telegram.{bot_name}`
- **InstagramPublisherNode**: `instagram.{account_name}` (planned)
- **LinkedInPublisherNode**: `linkedin.{account_name}` (planned)

### Media Nodes
- **ImageGeneratorNode**: `openai.{credential_name}` or `stability.{credential_name}`
- **ImageUploaderNode**: `cloudinary.{account_name}` or `aws.{account_name}`

### Input Nodes
- **DataFetcherNode**: May require API keys depending on the service
- **GoogleSheetsReaderNode**: `google.sheets.{project_name}` (planned)

---

## 📝 Best Practices

### 1. **Configuration Validation**
- Always validate required parameters
- Provide sensible defaults for optional parameters
- Use type checking for configuration values

### 2. **Error Handling**
- Return descriptive error messages
- Log errors with appropriate levels
- Handle timeouts and retries gracefully

### 3. **Performance**
- Use connection pooling for external services
- Implement caching where appropriate
- Monitor execution times

### 4. **Security**
- Never log sensitive data
- Validate input data
- Use secure credential management

### 5. **Testing**
- Write unit tests for each node
- Mock external dependencies
- Test error scenarios

---

## 🧪 Testing Nodes

### Example Test Structure
```go
func TestTextGeneratorNode(t *testing.T) {
    // Setup
    config := base.NodeConfig{
        ID:   "test_generator",
        Type: "text_generator",
        Name: "Test Generator",
        Parameters: map[string]interface{}{
            "model":           "gpt-3.5-turbo",
            "prompt_template": "Generate a test message: {topic}",
            "max_tokens":      100,
            "temperature":     0.7,
        },
    }
    
    node := ai.NewTextGeneratorNode("test-api-key", config)
    
    // Test validation
    err := node.Validate()
    if err != nil {
        t.Fatalf("Validation failed: %v", err)
    }
    
    // Test execution
    input := map[string]interface{}{
        "topic": "success",
    }
    
    output, err := node.Execute(context.Background(), input)
    if err != nil {
        t.Fatalf("Execution failed: %v", err)
    }
    
    // Assertions
    if output["generated_text"] == "" {
        t.Error("Expected generated text, got empty string")
    }
}
```

---

## 📚 Additional Resources

- [Pipeline Configuration Guide](../docs/PIPELINE_CONFIGURATION.md)
- [Credentials Management](../docs/SISTEMA_CREDENCIALES.md)
- [Testing Guide](../docs/TESTING_GUIDE.md)
- [API Reference](../docs/API_REFERENCE.md) 