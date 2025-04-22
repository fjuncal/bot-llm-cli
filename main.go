package main

import (
	"bot-llm-cli/internal/llm"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("🤖 Welcome to llm-cli!")
	fmt.Println("Type your question and press ENTER:")

	reader := bufio.NewReader(os.Stdin)
	question, _ := reader.ReadString('\n')
	question = strings.TrimSpace(question)

	response, err := llm.AskOllama(question)
	if err != nil {
		fmt.Println("Error talking to the AI:", err)
		return
	}

	fmt.Println("\n💬 Response from AI:\n", response)
}
