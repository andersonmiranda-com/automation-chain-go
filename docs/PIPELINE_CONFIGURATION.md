# Pipeline Configuration Guide

## 🎯 Overview
This guide explains how to configure pipelines in the n8n-like system, including node definitions, credentials, and scheduling.

## 📋 Pipeline Structure

### Basic Pipeline Configuration
```json
{
  "name": "pipeline_name",
  "description": "Pipeline description",
  "schedule": "0 9 * * *",
  "credentials": {
    "openai": "default",
    "telegram": "motivational_bot"
  },
  "nodes": [
    {
      "id": "node_id",
      "type": "node_type",
      "name": "Node Display Name",
      "config": {
        "parameter1": "value1",
        "parameter2": "value2"
      }
    }
  ]
}
```

### Configuration Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `name` | string | Yes | Unique identifier for the pipeline |
| `description` | string | No | Human-readable description |
| `schedule` | string | No | Cron expression for scheduling |
| `credentials` | object | Yes | Service credentials to use |
| `nodes` | array | Yes | Array of node definitions |

## 🔐 Credentials Configuration

### Credentials Object Structure
```json
{
  "credentials": {
    "openai": "credential_name",
    "telegram": "bot_name",
    "instagram": "account_name",
    "google": "project_name",
    "cloudinary": "account_name"
  }
}
```

### Available Credential Types

#### OpenAI Credentials
- `"openai": "default"` - Default API key
- `"openai": "premium"` - Premium API key (GPT-4)
- `"openai": "test"` - Test environment API key

#### Telegram Credentials
- `"telegram": "motivational_bot"` - Motivational content bot
- `"telegram": "news_bot"` - News content bot
- `"telegram": "personal_bot"` - Personal bot

#### Instagram Credentials
- `"instagram": "business_account"` - Business Instagram account
- `"instagram": "personal_account"` - Personal Instagram account

#### Google Services
- `"google": "project_1"` - Google Sheets/Drive project 1
- `"google": "project_2"` - Google Sheets/Drive project 2

#### Cloudinary Credentials
- `"cloudinary": "main_account"` - Primary Cloudinary account
- `"cloudinary": "backup_account"` - Backup Cloudinary account

## 📅 Scheduling

### Cron Expression Format
```
┌───────────── minute (0 - 59)
│ ┌───────────── hour (0 - 23)
│ │ ┌───────────── day of the month (1 - 31)
│ │ │ ┌───────────── month (1 - 12)
│ │ │ │ ┌───────────── day of the week (0 - 6) (Sunday to Saturday)
│ │ │ │ │
* * * * *
```

### Common Schedule Examples

| Schedule | Description |
|----------|-------------|
| `"0 9 * * *"` | Daily at 9:00 AM |
| `"0 12 * * *"` | Daily at 12:00 PM |
| `"0 9 * * 1"` | Every Monday at 9:00 AM |
| `"0 9 1 * *"` | First day of each month at 9:00 AM |
| `"*/30 * * * *"` | Every 30 minutes |
| `"0 */2 * * *"` | Every 2 hours |

## 🧩 Node Configuration

### Node Definition Structure
```json
{
  "id": "unique_node_id",
  "type": "node_type",
  "name": "Human Readable Name",
  "config": {
    "parameter1": "value1",
    "parameter2": "value2"
  }
}
```

### Node Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `id` | string | Yes | Unique identifier within the pipeline |
| `type` | string | Yes | Node type (see Node Types below) |
| `name` | string | Yes | Display name for the node |
| `config` | object | Yes | Node-specific configuration |

## 📋 Available Node Types

### AI Nodes
- `"text_generator"` - Generate text using OpenAI
- `"image_generator"` - Generate images using AI (planned)

### Publisher Nodes
- `"telegram_publisher"` - Publish to Telegram
- `"instagram_publisher"` - Publish to Instagram (planned)
- `"linkedin_publisher"` - Publish to LinkedIn (planned)

### Input Nodes
- `"topic_selector"` - Select topics from lists (planned)
- `"data_fetcher"` - Fetch data from APIs (planned)
- `"file_reader"` - Read data from files (planned)

### Media Nodes
- `"image_uploader"` - Upload images to cloud storage (planned)

### Utility Nodes
- `"text_formatter"` - Format and process text (planned)

## 📝 Configuration Examples

### 1. Motivational Telegram Pipeline
```json
{
  "name": "telegram_pipeline",
  "description": "Generates and publishes motivational texts to Telegram",
  "schedule": "0 9 * * *",
  "credentials": {
    "openai": "default",
    "telegram": "motivational_bot"
  },
  "nodes": [
    {
      "id": "text_generator",
      "type": "text_generator",
      "name": "Generate Motivational Text",
      "config": {
        "model": "gpt-3.5-turbo",
        "prompt_template": "Generate a short and powerful motivational text in Spanish. The text should be inspiring, positive, and motivate people to achieve their goals. It should be between 100-150 words and appropriate for sharing on social media. Do not include hashtags or emojis, just pure text.",
        "max_tokens": 300,
        "temperature": 0.8
      }
    },
    {
      "id": "telegram_publisher",
      "type": "telegram_publisher",
      "name": "Publish to Telegram",
      "config": {
        "message_template": "💪 *Daily Motivational Message*\n\n%s\n\n✨ Have an amazing day!"
      }
    }
  ]
}
```

### 2. News Telegram Pipeline
```json
{
  "name": "telegram_news_pipeline",
  "description": "Generates and publishes news to Telegram",
  "schedule": "0 12 * * *",
  "credentials": {
    "openai": "premium",
    "telegram": "news_bot"
  },
  "nodes": [
    {
      "id": "text_generator",
      "type": "text_generator",
      "name": "Generate News",
      "config": {
        "model": "gpt-4",
        "prompt_template": "Generate a brief and objective news article about technology. It should be informative and neutral, with a maximum of 200 words.",
        "max_tokens": 400,
        "temperature": 0.3
      }
    },
    {
      "id": "telegram_publisher",
      "type": "telegram_publisher",
      "name": "Publish News to Telegram",
      "config": {
        "message_template": "📰 *News of the Day*\n\n%s\n\n📅 %s"
      }
    }
  ]
}
```

### 3. Instagram Pipeline (Planned)
```json
{
  "name": "instagram_pipeline",
  "description": "Generates and publishes content to Instagram",
  "schedule": "0 18 * * *",
  "credentials": {
    "openai": "default",
    "instagram": "business_account",
    "cloudinary": "main_account"
  },
  "nodes": [
    {
      "id": "topic_selector",
      "type": "topic_selector",
      "name": "Select Topic",
      "config": {
        "topic_list": ["success", "motivation", "leadership", "growth"],
        "selection_method": "random"
      }
    },
    {
      "id": "text_generator",
      "type": "text_generator",
      "name": "Generate Post Text",
      "config": {
        "model": "gpt-3.5-turbo",
        "prompt_template": "Generate an Instagram post about {topic}. Include relevant hashtags.",
        "max_tokens": 200,
        "temperature": 0.7
      }
    },
    {
      "id": "image_generator",
      "type": "image_generator",
      "name": "Generate Image",
      "config": {
        "model": "dall-e-3",
        "size": "1080x1080",
        "quality": "standard"
      }
    },
    {
      "id": "image_uploader",
      "type": "image_uploader",
      "name": "Upload to Cloudinary",
      "config": {
        "service": "cloudinary",
        "folder": "instagram_posts"
      }
    },
    {
      "id": "instagram_publisher",
      "type": "instagram_publisher",
      "name": "Publish to Instagram",
      "config": {
        "caption_template": "{text}\n\n{hashtags}"
      }
    }
  ]
}
```

## 🔧 Advanced Configuration

### Environment Variables
You can use environment variables in your configuration:
```json
{
  "config": {
    "api_url": "${API_BASE_URL}/endpoint",
    "timeout": "${REQUEST_TIMEOUT:-30}"
  }
}
```

### Conditional Configuration
```json
{
  "config": {
    "model": "${ENVIRONMENT:-production}" === "production" ? "gpt-4" : "gpt-3.5-turbo"
  }
}
```

### Template Variables
Use template variables in prompts and messages:
```json
{
  "config": {
    "prompt_template": "Generate content about {topic} in {language}",
    "message_template": "📅 {date}\n\n{content}\n\n👤 {author}"
  }
}
```

## ✅ Validation Rules

### Required Fields
- `name` must be unique across all pipelines
- `credentials` must reference valid credential names
- `nodes` array must not be empty
- Each node must have valid `id`, `type`, and `config`

### Node Type Validation
- `type` must match an available node type
- `config` must contain required parameters for the node type
- Parameter types must match expected types

### Credential Validation
- Referenced credentials must exist in `credentials.json`
- Credential types must match node requirements
- Service-specific credentials must be properly configured

## 🚨 Common Errors

### Configuration Errors
```json
{
  "error": "Missing required field 'name'",
  "error": "Invalid node type 'unknown_type'",
  "error": "Credential 'invalid_credential' not found",
  "error": "Missing required parameter 'model' for text_generator"
}
```

### Runtime Errors
```json
{
  "error": "OpenAI API key is invalid",
  "error": "Telegram bot token is invalid",
  "error": "Channel ID not found",
  "error": "Rate limit exceeded"
}
```

## 📚 Best Practices

### 1. **Naming Conventions**
- Use descriptive pipeline names: `telegram_motivational_daily`
- Use clear node names: `Generate Motivational Text`
- Use consistent ID patterns: `text_generator_1`

### 2. **Configuration Organization**
- Group related parameters logically
- Use consistent formatting and indentation
- Add comments for complex configurations

### 3. **Error Handling**
- Provide fallback values for optional parameters
- Use appropriate timeouts and retry settings
- Log errors with sufficient detail

### 4. **Security**
- Never hardcode sensitive values
- Use environment variables for secrets
- Validate all input parameters

### 5. **Performance**
- Use appropriate model sizes for the task
- Set reasonable token limits
- Consider rate limits and quotas

## 📁 File Organization

### Recommended Structure
```
config/
├── credentials.json              # Service credentials
├── credentials_example.json      # Example credentials
├── pipelines/
│   ├── telegram.json             # Telegram pipeline
│   ├── telegram_news.json        # News pipeline
│   ├── instagram.json            # Instagram pipeline (planned)
│   └── linkedin.json             # LinkedIn pipeline (planned)
└── templates/
    ├── motivational.json         # Template configurations
    └── news.json                 # News templates
```

## 🔄 Pipeline Execution

### Execution Flow
1. **Load Configuration**: Read pipeline JSON file
2. **Validate Configuration**: Check all required fields
3. **Load Credentials**: Get service credentials
4. **Build Pipeline**: Create node instances
5. **Execute Nodes**: Run nodes in sequence
6. **Handle Results**: Process outputs and errors

### Execution Logs
```
2025-07-25 20:03:46 Building pipeline: telegram_pipeline
2025-07-25 20:03:46 Added node: Generate Motivational Text to pipeline: telegram_pipeline
2025-07-25 20:03:46 Added node: Publish to Telegram to pipeline: telegram_pipeline
2025-07-25 20:03:46 Pipeline telegram_pipeline built successfully with 2 nodes
2025-07-25 20:03:46 Starting pipeline execution: telegram_pipeline
2025-07-25 20:03:46 Executing node 1/2: Generate Motivational Text
2025-07-25 20:03:47 Node Generate Motivational Text completed successfully
2025-07-25 20:03:47 Executing node 2/2: Publish to Telegram
2025-07-25 20:03:48 Node Publish to Telegram completed successfully
2025-07-25 20:03:48 Pipeline telegram_pipeline execution completed successfully!
```

## 📚 Additional Resources

- [Nodes Documentation](./NODES_DOCUMENTATION.md)
- [Credentials Management](./SISTEMA_CREDENCIALES.md)
- [Testing Guide](./TESTING_GUIDE.md)
- [API Reference](./API_REFERENCE.md) 