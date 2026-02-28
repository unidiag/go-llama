package main

import (
	"fmt"
	llama "github.com/unidiag/go-llama"
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
				Content: "Напиши одно слово: какая главная функия в программе на Golang ?",
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