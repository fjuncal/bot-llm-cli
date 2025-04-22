package llm

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type OllamaRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OllamaResponse struct {
	Message Message `json:"message"`
}

func AskOllama(prompt string) (string, error) {

	url := "http://localhost:11434/api/chat"

	payload := OllamaRequest{
		Model: "mistral",
		Messages: []Message{
			{
				Role:    "user",
				Content: prompt,
			},
		},
	}

	// Convert struct to JSON
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	// Make request http to Ollama local
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	// Read body response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Convert Json response to struct
	var ollamaResponse OllamaResponse
	json.Unmarshal(body, &ollamaResponse)
	if err != nil {
		return "", err
	}

	return ollamaResponse.Message.Content, nil
}
