package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"

	"github.com/SakuraAxy/ChatGPT-Cli/pkg/Chat"
	"github.com/sashabaranov/go-openai"

	"net/http"
	"os"
	"time"
)

func checkChatGPTConnect(timeOut time.Duration) error {
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

func main() {

	proxy := flag.String("proxy", "", "http_proxy address")
	apiKey := flag.String("apiKey", "", "Open AI Api Key")
	flag.Parse()

	_ = os.Setenv("http_proxy", *proxy)
	_ = os.Setenv("https_proxy", *proxy)

	if *apiKey == "" {
		apiKey := os.Getenv("OPENAI_API_KEY")
		if apiKey == "" {
			fmt.Println("OPENAI_API_KEY and apikey option are not set.")
			flag.Usage()
			os.Exit(1)
		}

	}
	err := checkChatGPTConnect(2)
	if err != nil {
		fmt.Println("ChatGPT Connect Failed. Please Check your options: ")

		flag.Usage()
		os.Exit(1)
	}
	ctx := context.Background()

	client := openai.NewClient(*apiKey)

	quit := false

	fmt.Println("Hi , Im ChatGPT, You can ask me a question, and I will do my best to answer your question")

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
			Chat.GetResponse(client, ctx, questions)
		}

	}
	fmt.Println("ChatGPT bye.")

}
