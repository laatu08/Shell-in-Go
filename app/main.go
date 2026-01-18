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

		command := strings.TrimSpace(line)

		// For now, treat all commands as invalid
		if command != "" {
			fmt.Printf("%s: command not found\n", command)
		}
	}
}
