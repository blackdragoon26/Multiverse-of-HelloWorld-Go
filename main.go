package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var todos []string
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter todo (or type 'list' or 'quit'): ")

		if !scanner.Scan() {
			break
		}
		cmd := strings.TrimSpace(scanner.Text())

		switch cmd {
		case "quit":
			fmt.Println("Goodbye!")
			return
		case "list":
			fmt.Println("Your todos:")
			for i, todo := range todos {
				fmt.Printf("%d. %s\n", i+1, todo)
			}
		default:
			todos = append(todos, cmd)
			fmt.Printf("Added: %s\n", cmd)
		}
	}
}

