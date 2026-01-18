package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Builtin command registry
	builtins := map[string]bool{
		"exit": true,
		"echo": true,
		"type": true,
	}

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
		
		case "type":
			if len(args) == 0 {
				continue
			}

			target := args[0]
			if builtins[target] {
				fmt.Printf("%s is a shell builtin\n", target)
			} else {
				fmt.Printf("%s: not found\n", target)
			}

		default:
			fmt.Printf("%s: command not found\n", cmd)
		}
	}
}
