package assignmentrecognizeprovider

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
)

const (
	apiKey      = "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"
	apiEndpoint = "https://api.openai.com/v1/chat/completions"
)

type chatGptProvider struct {
	apiKey string
}

func NewChatGptProvider(apiKey string) *chatGptProvider {
	return &chatGptProvider{
		apiKey: apiKey,
	}
}

func (gpt *chatGptProvider) RecognizeAssignment(ctx context.Context, data map[string]interface{}) (map[string]interface{}, error) {
	client := resty.New()

	fmt.Println("running chat gpt provider")

	response, err := client.R().
		SetAuthToken(apiKey).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"model":      "gpt-3.5-turbo",
			"messages":   []interface{}{map[string]interface{}{"role": "system", "content": "Hi can you tell me what is the factorial of 10?"}},
			"max_tokens": 100,
		}).
		Post(apiEndpoint)
	if err != nil {
		return nil, err
	}
	fmt.Println(response.String())
	return nil, nil
}
