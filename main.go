package main

import (
	"bot-llm-cli/internal/llm"
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"os"
	"strings"
)

func main() {
	title := color.New(color.FgCyan, color.Bold).SprintFunc()
	prompt := color.New(color.FgYellow).SprintFunc()
	answer := color.New(color.FgGreen).SprintFunc()
	errorText := color.New(color.FgRed).SprintFunc()

	fmt.Println(title("🤖 Welcome to llm-cli!"))
	fmt.Println(prompt("Type your question and press ENTER:"))

	reader := bufio.NewReader(os.Stdin)
	question, _ := reader.ReadString('\n')
	question = strings.TrimSpace(question)

	response, err := llm.AskOllama(question)
	if err != nil {
		fmt.Println(errorText("Error talking to the AI:"), err)
		return
	}

	fmt.Println(answer("\n💬 Response from AI:\n"))
	fmt.Println(response)
}
