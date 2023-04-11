package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	openai "github.com/sashabaranov/go-openai"
	"io"
	"os"
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

func main() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		panic("OPENAI_API_KEY ENV IS NOT FOUND.")
	}
	ctx := context.Background()
	client := openai.NewClient(apiKey)
	quit := false
	for !quit {
		fmt.Print("\ninput your questions(type `q/Q/quit` to exit): ")
		reader := bufio.NewReader(os.Stdin)
		questions, _ := reader.ReadString('\n')
		switch questions {
		case "quit\n", "q\n", "Q\n":
			quit = true
		case "":
			continue
		default:
			GetResponse(client, ctx, questions)
		}

	}
	fmt.Println("Exit ChatGPT.")

}
