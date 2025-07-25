# Multiple Credentials Management System

## 🎯 Objective
Allow each node to use specific credentials for different services, facilitating the management of multiple accounts and projects.

## 🏗️ Architecture

### 1. **CredentialsManager**
Manages multiple credentials organized by service and name:

```go
type CredentialsManager struct {
    credentials map[string]interface{}
    mutex       sync.RWMutex
}
```

### 2. **Credentials Structure**
```json
{
  "openai": {
    "default": "sk-your-default-key",
    "premium": "sk-your-premium-key",
    "test": "sk-your-test-key"
  },
  "telegram": {
    "motivational_bot": {
      "token": "your-motivational-bot-token",
      "channel_id": "@motivational_channel"
    },
    "news_bot": {
      "token": "your-news-bot-token",
      "channel_id": "@news_channel"
    }
  }
}
```

### 3. **Node Configuration**
Each node specifies which credentials to use:

```json
{
  "id": "telegram_publisher",
  "type": "telegram_publisher",
  "name": "Publish to Telegram",
  "credentials": "motivational_bot",  // ✅ Specific credential
  "config": {
    "message_template": "💪 *Daily Motivation*\n\n%s"
  }
}
```

## 🔧 System Usage

### 1. **Load Credentials**
```go
// Create and load credentials
credentialsManager := config.NewCredentialsManager()
err := credentialsManager.LoadCredentials("config/credentials.json")
```

### 2. **Get Specific Credentials**
```go
// OpenAI
apiKey, err := credentialsManager.GetOpenAICredential("premium")

// Telegram
token, channelID, err := credentialsManager.GetTelegramCredential("news_bot")

// Instagram
accessToken, userID, err := credentialsManager.GetInstagramCredential("business_account")

// Google Sheets
serviceAccountFile, spreadsheetID, err := credentialsManager.GetGoogleSheetsCredential("project_1")

// Cloudinary
cloudName, apiKey, apiSecret, err := credentialsManager.GetCloudinaryCredential("main_account")
```

### 3. **Build Pipeline with Credentials**
```go
// Create builder with credentials
builder := pipelinebase.NewPipelineBuilder(credentialsManager)

// Build pipeline (credentials specified per node)
pipeline, err := builder.BuildPipeline(
    pipelineConfig.Name, 
    pipelineConfig.Nodes
)
```

## 📋 Configuration Examples

### **Motivational Pipeline (Default Credentials)**
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
        "prompt_template": "Generate a motivational text in Spanish..."
      }
    },
    {
      "id": "telegram_publisher",
      "type": "telegram_publisher",
      "name": "Publish to Telegram",
      "credentials": "motivational_bot",
      "config": {
        "message_template": "💪 *Daily Motivation*\n\n%s"
      }
    }
  ]
}
```

### **News Pipeline (Premium Credentials)**
```json
{
  "name": "telegram_news_pipeline",
  "description": "Generates and publishes news to Telegram",
  "schedule": "0 12 * * *",
  "nodes": [
    {
      "id": "text_generator",
      "type": "text_generator",
      "name": "Generate News",
      "credentials": "premium",
      "config": {
        "model": "gpt-4",
        "prompt_template": "Generate a brief news article..."
      }
    },
    {
      "id": "telegram_publisher",
      "type": "telegram_publisher",
      "name": "Publish News to Telegram",
      "credentials": "news_bot",
      "config": {
        "message_template": "📰 *News of the Day*\n\n%s"
      }
    }
  ]
}
```

### **Multi-Channel Pipeline**
```json
{
  "name": "multi_telegram_pipeline",
  "description": "Publishes content to multiple Telegram channels",
  "schedule": "0 9 * * *",
  "nodes": [
    {
      "id": "text_generator",
      "type": "text_generator",
      "name": "Generate Content",
      "credentials": "default",
      "config": {
        "model": "gpt-3.5-turbo",
        "prompt_template": "Generate motivational content..."
      }
    },
    {
      "id": "telegram_motivational",
      "type": "telegram_publisher",
      "name": "Publish to Motivational Channel",
      "credentials": "motivational_bot",
      "config": {
        "message_template": "💪 *Daily Motivation*\n\n%s"
      }
    },
    {
      "id": "telegram_news",
      "type": "telegram_publisher",
      "name": "Publish to News Channel",
      "credentials": "news_bot",
      "config": {
        "message_template": "📰 *Daily Inspiration*\n\n%s"
      }
    },
    {
      "id": "telegram_personal",
      "type": "telegram_publisher",
      "name": "Publish to Personal Channel",
      "credentials": "personal_bot",
      "config": {
        "message_template": "💭 *Personal Note*\n\n%s"
      }
    }
  ]
}
```

## 🔐 Supported Services

### 1. **OpenAI**
- Multiple API keys
- Different models (GPT-3.5, GPT-4)
- Usage by project or content type

### 2. **Telegram**
- Multiple bots
- Different channels
- Specific tokens per bot

### 3. **Instagram**
- Personal and business accounts
- Different access tokens
- Specific user IDs

### 4. **Google Services**
- Multiple projects
- Service account files
- Specific spreadsheet IDs

### 5. **Cloudinary**
- Multiple accounts
- API keys and secrets
- Specific cloud names

## 🧪 Testing

### **Test Credentials**
```go
// Create test credentials
testCredentials := map[string]interface{}{
    "openai": map[string]interface{}{
        "default": "sk-test-default-key",
    },
    "telegram": map[string]interface{}{
        "motivational_bot": map[string]interface{}{
            "token":      "test-bot-token",
            "channel_id": "@test_channel",
        },
    },
}

// Set test credentials
credentialsManager.SetCredentials(testCredentials)
```

### **Unit Test**
```go
func TestPipelineBuilder(t *testing.T) {
    // Test credentials
    credentialsManager := config.NewCredentialsManager()
    credentialsManager.SetCredentials(testCredentials)
    
    // Builder with credentials
    builder := pipelinebase.NewPipelineBuilder(credentialsManager)
    
    // Build pipeline
    pipeline, err := builder.BuildPipeline("test", nodes)
    // ... assertions
}
```

## 🔒 Security

### 1. **Credential Files**
- **DO NOT** include in version control
- Use `.gitignore` for `credentials.json`
- Provide `credentials_example.json`

### 2. **Thread-Safe Access**
- Mutex for concurrent access
- Protected read/write operations
- No race conditions

### 3. **Validation**
- Verification of required credentials
- Robust error handling
- Logs without sensitive information

## 📁 File Structure

```
config/
├── credentials.json              # Real credentials (NOT in git)
├── credentials_example.json      # Structure example
├── credentials_test.json         # Test credentials
└── pipelines/
    ├── telegram.json             # Pipeline with default credentials
    ├── telegram_news.json        # Pipeline with premium credentials
    └── multi_telegram.json       # Multi-channel pipeline
```

## 🚀 System Advantages

### 1. **Flexibility**
- ✅ Multiple accounts per service
- ✅ Per-node credential configuration
- ✅ Easy credential changes

### 2. **Scalability**
- ✅ Easy to add new services
- ✅ Multiple simultaneous projects
- ✅ Environment separation

### 3. **Maintainability**
- ✅ Centralized credentials
- ✅ Declarative configuration
- ✅ Independent testing

### 4. **Security**
- ✅ Controlled access
- ✅ No hardcoding
- ✅ Secret management

### 5. **Multi-Channel Support**
- ✅ Multiple Telegram channels in one pipeline
- ✅ Different credentials per node
- ✅ No global credential dependencies

## 📝 Next Steps

1. **Implement more services**: LinkedIn, Twitter, etc.
2. **Encryption**: Encrypt sensitive credentials
3. **Environment variables**: Support for env vars
4. **Automatic rotation**: Renew credentials
5. **Audit**: Credential usage logging

## 🔄 Migration from Old System

### **Before (Global Credentials)**
```json
{
  "credentials": {
    "openai": "default",
    "telegram": "motivational_bot"
  }
}
```

### **After (Per-Node Credentials)**
```json
{
  "nodes": [
    {
      "credentials": "default",
      "type": "text_generator"
    },
    {
      "credentials": "motivational_bot",
      "type": "telegram_publisher"
    }
  ]
}
```

### **Benefits of Migration**
- ✅ **Maximum flexibility**: Each node can use different credentials
- ✅ **Reusability**: A node can be used in different pipelines with different credentials
- ✅ **Scalability**: Easy to add new services and credentials
- ✅ **Clarity**: It's obvious which credential each node uses 