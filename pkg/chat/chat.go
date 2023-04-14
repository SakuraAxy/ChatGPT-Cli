package chat

import (
	"context"
	"errors"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"io"
	"net/http"
	"os"
	"time"
)

func GetResponse(c *openai.Client, ctx context.Context, quesiton string) {
	req := openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo,
		MaxTokens: 2000,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: quesiton,
			},
		},
		Stream: true,
	}

	stream, err := c.CreateChatCompletionStream(ctx, req)
	if err != nil {
		fmt.Printf("ChatCompletionStream error: %v\n", err)
		return
	}
	defer stream.Close()

	fmt.Printf("Response: \n")
	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			return
		}

		if err != nil {
			fmt.Printf("\nStream error: %v\n", err)
			return
		}

		fmt.Printf(response.Choices[0].Delta.Content)
	}
}

func GetStaticResponse(c *openai.Client, ctx context.Context, quesiton string) {

	resp, err := c.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: quesiton,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)
}

func CheckChatGPTConnect(timeOut time.Duration) error {
	client := &http.Client{
		Timeout: time.Second * timeOut,
	}

	req, err := http.NewRequest("GET", "https://api.openai.com/", nil)
	if err != nil {
		fmt.Println("Check ChatGPT API Failed, Error request:", err)
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Check ChatGPT API Failed, Error request:", err)
		return err
	}

	// 关闭响应体
	defer resp.Body.Close()
	return nil
}

func SetHTTPProxies(httpProxy, httpsProxy string) error {
	if err := os.Setenv("http_proxy", httpProxy); err != nil {
		return fmt.Errorf("failed to set http_proxy: %v", err)
	}

	if err := os.Setenv("https_proxy", httpsProxy); err != nil {
		return fmt.Errorf("failed to set https_proxy: %v", err)
	}

	return nil
}
