package llama

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

type Client struct {
	BaseURL     string
	APIKey      string
	HTTP        *http.Client
	Temperature float64
	MaxTokens   int
}

func New(baseURL string, apiKey string) *Client {
	return &Client{
		BaseURL:     baseURL,
		APIKey:      apiKey,
		HTTP:        &http.Client{Timeout: 10 * time.Minute},
		Temperature: 0.7, // default
		MaxTokens:   512, // default
	}
}

func (c *Client) SetDefaults(temp float64, max int) *Client {
	c.Temperature = temp
	c.MaxTokens = max
	return c
}

func (c *Client) Chat(req ChatRequest) (*ChatResponse, error) {

	// Apply defaults if zero values
	if req.Temperature == 0 {
		req.Temperature = c.Temperature
	}

	if req.MaxTokens == 0 {
		req.MaxTokens = c.MaxTokens
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest(
		"POST",
		c.BaseURL+"/v1/chat/completions",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/json")

	// 🔐 Добавляем API ключ если он задан
	if c.APIKey != "" {
		httpReq.Header.Set("Authorization", "Bearer "+c.APIKey)
	}

	start := time.Now()

	resp, err := c.HTTP.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	duration := time.Since(start)

	if resp.StatusCode != http.StatusOK {
		data, _ := io.ReadAll(resp.Body)
		return nil, errors.New(string(data))
	}

	var out ChatResponse
	err = json.NewDecoder(resp.Body).Decode(&out)
	if err != nil {
		return nil, err
	}

	out.GenerationTime = duration

	return &out, nil
}
