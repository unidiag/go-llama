package main

import (
	"fmt"
	"github.com/yourname/go-llama-client"
)

func main() {

	client := llama.New("http://192.168.1.55:80")

	resp, err := client.Chat(llama.ChatRequest{
		Model: "qwen",
		Messages: []llama.Message{
			{
				Role:    "system",
				Content: "You are a professional Golang assistant.",
			},
			{
				Role:    "user",
				Content: "Explain goroutines briefly",
			},
		},
		Temperature: 0.7,
		MaxTokens:   512,
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Choices[0].Message.Content)
}