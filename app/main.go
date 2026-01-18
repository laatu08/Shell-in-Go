package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
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

		parts := strings.Fields(line)
		cmd := parts[0]
		args := parts[1:]

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
				continue
			}

			// PATH search
			found := false
			pathEnv := os.Getenv("PATH")
			dirs := strings.Split(pathEnv, ":")

			for _, dir := range dirs {
				fullPath := filepath.Join(dir, target)

				info, err := os.Stat(fullPath)
				if err != nil {
					continue
				}

				// Must be a regular file with execute permission
				if info.Mode().IsRegular() && info.Mode()&0111 != 0 {
					fmt.Printf("%s is %s\n", target, fullPath)
					found = true
					break
				}
			}

			// Not found
			if !found {
				fmt.Printf("%s: not found\n", target)
			}

		default:
			fmt.Printf("%s: command not found\n", cmd)
		}
	}
}
