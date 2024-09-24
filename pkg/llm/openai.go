package llm

import (
	"context"
	"fmt"
	"os"

	"github.com/sashabaranov/go-openai"
)

func GenerateCompletion(query string) (string, error) {
	token := os.Getenv("OPENAI_API_KEY")
	if token == "" {
		return "", fmt.Errorf("failed to create the openai client, env OPENAI_API_KEY not defined")
	}

	client := openai.NewClient(token)

	resp, err := client.CreateChatCompletion(context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4o,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: query,
				},
			},
		},
	)

	if err != nil {
		return "", fmt.Errorf("failed to craeate response, %s", err)
	}

	return resp.Choices[0].Message.Content, nil
}
