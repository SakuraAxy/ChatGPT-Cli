package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/SakuraAxy/ChatGPT-Cli/pkg/Chat"
	"github.com/sashabaranov/go-openai"
	"os"
)

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
			Chat.GetResponse(client, ctx, questions)
		}

	}
	fmt.Println("Exit ChatGPT.")

}
