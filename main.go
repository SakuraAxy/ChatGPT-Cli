package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"

	"github.com/SakurAxy/ChatGPT-Cli/pkg/chat"
	"github.com/sashabaranov/go-openai"

	"os"
)

func main() {

	proxy := flag.String("proxy", "", "http_proxy address")
	apiKey := flag.String("apiKey", "", "Open AI Api Key")
	questions := flag.String("q", "", "Your questions")
	flag.Parse()
	chat.SetHTTPProxies(*proxy, *proxy)

	if *apiKey == "" {
		apiKey := os.Getenv("OPENAI_API_KEY")
		if apiKey == "" {
			fmt.Println("OPENAI_API_KEY and apikey option are not set.")
			flag.Usage()
			os.Exit(1)
		}

	}
	err := chat.CheckChatGPTConnect(2)
	if err != nil {
		fmt.Println("ChatGPT Connect Failed. Please Check your options: ")

		flag.Usage()
		os.Exit(1)
	}
	ctx := context.Background()

	client := openai.NewClient(*apiKey)

	if *questions != "" {
		chat.GetStaticResponse(client, ctx, *questions)

	} else {
		fmt.Println("Hi , Im ChatGPT, You can ask me a question, and I will do my best to answer your question")

		quit := false

		for !quit {
			fmt.Print("\nTyping your questions(use `q | quit` to exit console) : ")
			reader := bufio.NewReader(os.Stdin)
			questions, _ := reader.ReadString('\n')
			switch questions {
			case "quit\n", "q\n":
				quit = true
			case "":
				continue
			default:
				chat.GetResponse(client, ctx, questions)
			}

		}

	}

	fmt.Println("ChatGPT bye.")

}
