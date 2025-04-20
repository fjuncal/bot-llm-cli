package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("🤖 Welcome to llm-cli!")
	fmt.Println("Type your question and press ENTER:")

	reader := bufio.NewReader(os.Stdin)
	question, _ := reader.ReadString('\n')

	fmt.Println("You asked: ", question)
	fmt.Println("Answer: (here would be the AI response)")
}
