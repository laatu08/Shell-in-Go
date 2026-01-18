package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
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
			pathEnv := os.Getenv("PATH")
			dirs := strings.Split(pathEnv, string(os.PathListSeparator))

			exts := []string{""} // Unix-style: no extension
			if pe := os.Getenv("PATHEXT"); pe != "" {
				exts = strings.Split(pe, string(os.PathListSeparator))
			}

			found := false

			for _, dir := range dirs {
				for _, ext := range exts {
					fullPath := filepath.Join(dir, target+ext)

					info, err := os.Stat(fullPath)
					if err != nil {
						continue
					}

					if info.Mode().IsRegular() {
						if runtime.GOOS == "windows" {
							// On Windows, existence + extension is enough
							fmt.Printf("%s is %s\n", target, fullPath)
							found = true
							break
						}

						// Unix: must have execute permission
						if info.Mode()&0111 != 0 {
							fmt.Printf("%s is %s\n", target, fullPath)
							found = true
							break
						}
					}

				}
				if found {
					break
				}
			}

			if !found {
				fmt.Printf("%s: not found\n", target)
			}

		default:
			fmt.Printf("%s: command not found\n", cmd)
		}
	}
}
