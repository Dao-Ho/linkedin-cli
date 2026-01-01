package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	displayFeed()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("\nCommands: (q)uit: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		switch input {
		case "q":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Unknown command.")
		}
	}
}
