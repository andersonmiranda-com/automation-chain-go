# Roadmap - n8n-like Pipeline System

## 🎯 Objective
Create a modular pipeline system with reusable and configurable nodes, similar to n8n, to automate content generation and publication across multiple platforms.

## 📋 Implementation Phases

### Phase 1: Base Structure and Existing Nodes ✅ COMPLETED
- [x] Define new folder structure
- [x] Migrate existing nodes to new architecture
- [x] Implement per-pipeline configuration system
- [x] Create pipeline builder
- [x] Test existing nodes
- [x] Implement credentials management system
- [x] Support per-node credentials
- [x] Multi-channel pipeline support
- [x] Clean up obsolete files
- [x] Complete documentation in English

### Phase 2: Input Nodes 🔄
- [ ] TopicSelectorNode (topic selection)
- [ ] DataFetcherNode (fetch external data)
- [ ] FileReaderNode (read files)
- [ ] Individual testing of each node

### Phase 3: AI Nodes ⏳
- [ ] TextGeneratorNode (enhanced with configuration)
- [ ] ImageGeneratorNode (DALL-E/Midjourney)
- [ ] ContentAnalyzerNode (content analysis)
- [ ] Individual testing of each node

### Phase 4: Media Nodes ⏳
- [ ] ImageUploaderNode (Cloudinary/S3)
- [ ] ImageProcessorNode (resize, optimize)
- [ ] MediaValidatorNode (validate formats)
- [ ] Individual testing of each node

### Phase 5: Publisher Nodes ⏳
- [ ] LinkedInPublisherNode (enhanced)
- [ ] InstagramPublisherNode
- [ ] TwitterPublisherNode
- [ ] TelegramPublisherNode (enhanced)
- [ ] Individual testing of each node

### Phase 6: Google Services Nodes ⏳
- [ ] SheetsReaderNode (read Google Sheets)
- [ ] SheetsWriterNode (write Google Sheets)
- [ ] SheetsUpdaterNode (update Google Sheets)
- [ ] DriveUploaderNode (upload to Google Drive)
- [ ] DriveDownloaderNode (download from Google Drive)
- [ ] DriveSearcherNode (search in Google Drive)
- [ ] Individual testing of each node

### Phase 7: Specialized Pipelines ⏳
- [ ] LinkedInPipeline (complete with Google Sheets)
- [ ] InstagramPipeline (complete with Google Drive)
- [ ] TelegramPipeline (enhanced)
- [ ] Complete pipeline testing

### Phase 8: Scheduling System ⏳
- [ ] Scheduler with cron expressions
- [ ] Per-pipeline schedule configuration
- [ ] Automatic execution
- [ ] Scheduling testing

### Phase 9: Logging and Monitoring System ⏳
- [ ] Centralized logs
- [ ] Success/failure metrics
- [ ] Email/Slack alerts
- [ ] Basic dashboard

### Phase 10: Documentation and Optimization ⏳
- [ ] Complete documentation
- [ ] Usage guides
- [ ] Performance optimization
- [ ] Load testing

## 🏗️ Current Folder Structure

```
test-chain-go-cursor/
├── nodes/
│   ├── base/
│   │   └── node.go              # Base Node interface
│   ├── ai/
│   │   └── text_generator.go    # Text generator node
│   └── publishers/
│       └── telegram_publisher.go # Telegram publisher node
├── pipelines/
│   └── base/
│       ├── pipeline.go          # Base pipeline
│       └── builder.go           # Pipeline builder
├── config/
│   ├── credentials.go           # Credentials management
│   ├── credentials_example.json # Credentials example
│   ├── credentials_test.json    # Test credentials
│   └── pipelines/
│       ├── telegram.json        # Pipeline configuration
│       ├── telegram_news.json   # News pipeline
│       └── multi_telegram.json  # Multi-channel pipeline
├── tests/
│   └── pipeline_test.go         # Unit tests
├── docs/
│   ├── NODES_DOCUMENTATION.md   # Complete node documentation
│   ├── PIPELINE_CONFIGURATION.md # Pipeline configuration guide
│   ├── SISTEMA_CREDENCIALES.md  # Credentials system documentation
│   └── FASE1_COMPLETADA.md      # Phase 1 completion report
└── main.go                      # Main entry point
```

## 🧪 Testing Strategy

### For Each Node:
1. **Unit Test** - Basic functionality
2. **Integration Test** - With real data
3. **Configuration Test** - Different configurations
4. **Error Test** - Error handling

### For Each Pipeline:
1. **Build Test** - Pipeline builds correctly
2. **Execution Test** - Pipeline executes all nodes
3. **Data Test** - Data flows correctly between nodes
4. **Error Test** - Error handling in any node

## 📝 Implementation Notes

### Principles:
- ✅ Each node must be individually testable
- ✅ Declarative configuration (JSON)
- ✅ Maximum code reusability
- ✅ Robust error handling
- ✅ Detailed logging
- ✅ Per-node credentials
- ✅ No global dependencies

### Conventions:
- Nodes: `PascalCase` (e.g., `TopicSelectorNode`)
- Configurations: `snake_case` (e.g., `topic_selector.json`)
- Pipelines: `snake_case` (e.g., `linkedin_pipeline.json`)
- Documentation: English only

## 🚀 Current Status

### ✅ Phase 1 Completed:
- **Modular Architecture**: Clean, scalable structure
- **Credential Management**: Per-node credentials with multiple service support
- **Multi-Channel Support**: Multiple Telegram channels in one pipeline
- **Testing**: Unit tests and integration tests working
- **Documentation**: Complete documentation in English
- **Clean Codebase**: No obsolete files

### 🔄 Next Step: Phase 2
Implement **Input Nodes** to expand the system's data input capabilities.

## 📊 Achievements

### Technical Achievements:
- ✅ **Modularity**: Independent and reusable nodes
- ✅ **Configurability**: Pipelines configurable without code
- ✅ **Testability**: Unit and integration tests
- ✅ **Scalability**: Easy to add new nodes and pipelines
- ✅ **Maintainability**: Clean and well-structured code
- ✅ **Credential Management**: Per-node credentials with multiple service support

### Business Value:
- ✅ **Flexibility**: Multiple accounts per service
- ✅ **Efficiency**: Automated content generation and publication
- ✅ **Reliability**: Robust error handling and testing
- ✅ **Scalability**: Easy to add new platforms and services 