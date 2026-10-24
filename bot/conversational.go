package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms/googleai"
	"github.com/tmc/langchaingo/memory"
)

func conversationalChat() {
	ctx := context.Background()
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Failed to load .env: %v", err)
	}
	apiKey := os.Getenv("GEMINI_API_KEY")
	llm, _ := googleai.New(ctx, googleai.WithAPIKey(apiKey),
		googleai.WithDefaultModel("gemini-3.8-flash"))
	buffer := memory.NewConversationBuffer()
	conversationchain := chains.NewConversation(llm, buffer)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("\nYou:")
		userInput, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Error reading input: %v\n", err)
			continue
		}
		userInput = strings.TrimSpace(userInput)
		if userInput == "" {
			fmt.Println("Please Enter a message!!")
			continue
		}
		fmt.Print("Agent: ")
		fmt.Print("Calculating response...")
		response, err := chains.Run(ctx, conversationchain, "Reply concisely: "+userInput)
		//clear the line after calculating response
		fmt.Print("\r Agent: ")
		if err != nil {
			fmt.Printf("Error generating response: %v\n", err)
			continue
		}
		fmt.Println(response)

	}

}
