# go-llama

Minimal and clean Go client for OpenAI-compatible LLM servers
(llama.cpp, Mistral, Qwen, etc.).

Supports:

-   Chat completion
-   Streaming responses (SSE)
-   Simple configuration
-   Examples included

------------------------------------------------------------------------

## Installation

``` bash
go get github.com/unidiag/go-llama@latest
```

------------------------------------------------------------------------

## Quick Start

### Simple Chat Request

``` go
package main

import (
    "fmt"
    llama "github.com/unidiag/go-llama"
)

func main() {

    client := llama.New("http://localhost:80", "YOUR_API_KEY")

    resp, err := client.Chat(llama.ChatRequest{
        Messages: []llama.Message{
            {Role: "system", Content: "You are a helpful assistant."},
            {Role: "user", Content: "Hello!"},
        },
        Temperature: 0.7,
        MaxTokens:   512,
    })

    if err != nil {
        panic(err)
    }

    fmt.Println(resp.Choices[0].Message.Content)
}
```

------------------------------------------------------------------------

## Streaming Example

``` go
err := client.ChatStream(req, func(token string) {
    fmt.Print(token)
})
```

`ChatStream` automatically:

-   Enables `stream=true`
-   Parses SSE chunks
-   Filters `[DONE]`
-   Calls your callback for each token


------------------------------------------------------------------------

# Examples

The repository contains working examples:

examples/
├── simple/
│    └── main.go 
├── stream/
│    └── main.go 
└── server/
     └── main.go

------------------------------------------------------------------------

## 1️⃣ Streaming CLI Example

Run:

``` bash
cd examples/stream
go run main.go
```

------------------------------------------------------------------------

## 2️⃣ Web Streaming Server Example

Run:

``` bash
cd examples/server
go run main.go
```

Open in browser:

http://localhost:8080

Features:

-   HTML form for system/user messages
-   Server-Sent Events streaming
-   Real-time model output in browser

------------------------------------------------------------------------

# API Reference

## Create Client

``` go
client := llama.New(baseURL, apiKey)
```

------------------------------------------------------------------------

## Chat Request

``` go
type ChatRequest struct {
    Messages    []Message
    Temperature float32
    MaxTokens   int
    Stream      bool
}
```

------------------------------------------------------------------------

## Message Roles

  Role        Description
  ----------- ------------------------------
  system      Global instruction
  user        User message
  assistant   Model response (for history)
  tool        Tool/function calling

------------------------------------------------------------------------

# Requirements

-   LLM server compatible with `/v1/chat/completions`
-   Go 1.20+

------------------------------------------------------------------------

# License

MIT
