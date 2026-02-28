package main

import (
	"fmt"

	llama "github.com/unidiag/go-llama"
)

func main() {

	client := llama.New("http://192.168.1.55:80", "MAGICKSECRETKEY")
	//client.SetDefaults(0.5, 100)

	r, err := client.Chat(llama.ChatRequest{

		Messages: []llama.Message{
			// роли:
			// `system` - глобальные инструкции модели.
			// `user` - сообщение пользователя.
			// `assistant` - ответ модели (нужен для истории диалога).
			// `tool` - используется при function-calling / tool-calling.
			{
				Role:    "system",
				Content: "Ты профессиональный редактор EPG. Отвечай кратко. Не выдумывай телеканалы.",
			},
			{
				Role:    "user",
				Content: "Перечисли телеканалы, по которым в Беларуси можно посмотреть спортивные соревнования",
			},
		},

		//Temperature: 0.7,  // Температура управляет креативностью ответа.
		//MaxTokens:   512, // Максимальное количество токенов в ответе.
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(r.Choices[0].Message.Content)

}
