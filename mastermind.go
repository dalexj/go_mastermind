package main

import (
	"bufio"
	"strings"
	"fmt"
	"os"
)

func getInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func main() {
	mastermind := game{code: "rrrr", guesses: 0}
	for {
		input := getInput("Enter a 4 color guess: ")
		fmt.Println("guesses:", mastermind.guesses)
		numCorrect := mastermind.guess(input)
		fmt.Println("You got ", numCorrect, " correct colors")
		if numCorrect == 4 {
			fmt.Println("You win, go home.")
			break
		}
	}
}