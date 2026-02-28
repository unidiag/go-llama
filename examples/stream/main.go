package main

import (
	"fmt"
	"log"

	llama "github.com/unidiag/go-llama"
)

func main() {

	client := llama.New("http://192.168.1.55:80", "MAGICKSECRETKEY")
	//client.SetDefaults(0.5, 100)

	req := llama.ChatRequest{
		Messages: []llama.Message{
			{
				Role:    "system",
				Content: "Ты профессиональный редактор EPG. Отвечай кратко. Не выдумывай телеканалы.",
			},
			{
				Role:    "user",
				Content: "Перечисли телеканалы, по которым в Беларуси можно посмотреть спортивные соревнования",
			},
		},

		// Temperature: 0.7,
		// MaxTokens:   512,
	}

	err := client.ChatStream(req, func(token string) {
		fmt.Print(token) // потоковый вывод
	})

	fmt.Println()

	if err != nil {
		log.Fatal(err)
	}
}
