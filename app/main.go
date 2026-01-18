package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")

		// Read input
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println()
			return
		}

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// 4. Tokenize
		parts := strings.Fields(line)
		cmd := parts[0]
		args := parts[1:]

		// 5. Builtins
		switch cmd {
		case "exit":
			return

		case "echo":
			// Print args joined by space + newline
			fmt.Println(strings.Join(args, " "))

		default:
			fmt.Printf("%s: command not found\n", cmd)
		}
	}
}
