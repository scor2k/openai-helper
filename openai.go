package main

import (
	"context"
	"fmt"
	"os/exec"

	openai "github.com/sashabaranov/go-openai"

	"encoding/json"
	"os"
)

type Prompts struct {
	Rewrite string `json:"rewrite"`
}

func loadPrompts(path string) (Prompts, error) {
	filePath := os.Getenv("HOME") + "/.config/prompts.json"
	if path != "" {
		filePath = path
	}

	// Read the JSON file
	file, err := os.ReadFile(filePath)
	if err != nil {
		return Prompts{}, err
	}

	// Unmarshal the JSON into a map[string]string
	var prompts Prompts
	err = json.Unmarshal(file, &prompts)
	if err != nil {
		return Prompts{}, err
	}

	return prompts, nil
}

func sendOpenAIRequest(prompt string, request string, model string) (string, error) {
	client := openai.NewClient(os.Getenv("OPENAI_HELPER"))

	fullContent := fmt.Sprintf("%s\n\n%s", prompt, request)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: fullContent,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}

func getContentViaVim() (string, error) {
	tmpFile, err := os.CreateTemp("", "content")
	if err != nil {
		return "", err
	}
	defer os.Remove(tmpFile.Name())

	cmd := exec.Command("vim", tmpFile.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return "", err
	}

	content, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		return "", err
	}

	return string(content), nil
}
