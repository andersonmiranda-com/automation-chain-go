package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// Update represents a Telegram update
type Update struct {
	UpdateID int `json:"update_id"`
	Message  struct {
		Chat struct {
			ID       int64  `json:"id"`
			Type     string `json:"type"`
			Title    string `json:"title,omitempty"`
			Username string `json:"username,omitempty"`
		} `json:"chat"`
	} `json:"message,omitempty"`
	ChannelPost struct {
		Chat struct {
			ID       int64  `json:"id"`
			Type     string `json:"type"`
			Title    string `json:"title,omitempty"`
			Username string `json:"username,omitempty"`
		} `json:"chat"`
	} `json:"channel_post,omitempty"`
}

// UpdatesResponse represents the response from getUpdates
type UpdatesResponse struct {
	OK     bool     `json:"ok"`
	Result []Update `json:"result"`
}

func main() {
	// Parse command line arguments
	botToken := flag.String("token", "", "Telegram bot token")
	flag.Parse()

	if *botToken == "" {
		fmt.Println("Error: Bot token is required")
		fmt.Println("Usage: go run get_channel_id.go -token <BOT_TOKEN>")
		fmt.Println("Example: go run get_channel_id.go -token 1234567890:ABCdefGHIjklMNOpqrsTUVwxyz")
		os.Exit(1)
	}
	url := fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates", *botToken)

	fmt.Println(url)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to get updates: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response: %v", err)
	}

	var updates UpdatesResponse
	if err := json.Unmarshal(body, &updates); err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	if !updates.OK {
		fmt.Printf("Error: %s\n", string(body))
		os.Exit(1)
	}

	fmt.Println("📋 Found chats:")
	fmt.Println("================")

	seenChats := make(map[int64]bool)

	for _, update := range updates.Result {
		var chatID int64
		var chatType, title, username string

		// Check if it's a regular message
		if update.Message.Chat.ID != 0 {
			chatID = update.Message.Chat.ID
			chatType = update.Message.Chat.Type
			title = update.Message.Chat.Title
			username = update.Message.Chat.Username
		} else if update.ChannelPost.Chat.ID != 0 {
			// Check if it's a channel post
			chatID = update.ChannelPost.Chat.ID
			chatType = update.ChannelPost.Chat.Type
			title = update.ChannelPost.Chat.Title
			username = update.ChannelPost.Chat.Username
		}

		if chatID != 0 && !seenChats[chatID] {
			seenChats[chatID] = true
			
			fmt.Printf("Chat ID: %d\n", chatID)
			fmt.Printf("Type: %s\n", chatType)
			if title != "" {
				fmt.Printf("Title: %s\n", title)
			}
			if username != "" {
				fmt.Printf("Username: @%s\n", username)
			}
			
			// Provide the correct format for the config
			if chatType == "channel" || chatType == "supergroup" {
				if username != "" {
					fmt.Printf("✅ Use in .env: TELEGRAM_CHANNEL_ID=@%s\n", username)
				} else {
					fmt.Printf("✅ Use in .env: TELEGRAM_CHANNEL_ID=%d\n", chatID)
				}
			}
			fmt.Println("----------------")
		}
	}

	if len(seenChats) == 0 {
		fmt.Println("❌ No chats found. Make sure:")
		fmt.Println("   1. Your bot token is correct")
		fmt.Println("   2. The bot has been added to a channel/group")
		fmt.Println("   3. Someone has sent a message in the channel/group")
		fmt.Println("   4. The bot has permission to read messages")
	}
} 