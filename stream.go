package llama

import (
	"bufio"
	"encoding/json"
	"net/http"
	"strings"
)

func (c *Client) ChatStream(
	req ChatRequest,
	onToken func(string),
) error {

	req.Stream = true

	body, _ := json.Marshal(req)

	httpReq, _ := http.NewRequest(
		"POST",
		c.BaseURL+"/v1/chat/completions",
		strings.NewReader(string(body)),
	)

	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTP.Do(httpReq)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)

	for scanner.Scan() {
		line := scanner.Text()

		if !strings.HasPrefix(line, "data: ") {
			continue
		}

		data := strings.TrimPrefix(line, "data: ")

		if data == "[DONE]" {
			break
		}

		var chunk map[string]any
		err := json.Unmarshal([]byte(data), &chunk)
		if err != nil {
			continue
		}

		choices := chunk["choices"].([]any)
		if len(choices) > 0 {
			delta := choices[0].(map[string]any)["delta"].(map[string]any)
			if content, ok := delta["content"].(string); ok {
				onToken(content)
			}
		}
	}

	return nil
}