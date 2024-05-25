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
		fmt.Fprint(os.Stdout, "$ ")

		commandInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		commandInput = strings.TrimSpace(commandInput)

		switch {
		case commandInput == "exit 0":
			return
		case commandInput == "echo":
			fmt.Println()
		case strings.HasPrefix(commandInput, "echo "):
			fmt.Println(strings.TrimPrefix(commandInput, "echo "))
		case strings.HasPrefix(commandInput, "type "):
			commandType := strings.TrimPrefix(commandInput, "type ")
			switch commandType {
			case "echo", "exit", "type":
				fmt.Printf("%s is a shell builtin\n", commandType)
			default:
				fmt.Printf("%s not found\n", commandType)
			}
		default:
			fmt.Printf("%s: command not found\n", commandInput)
		}
	}
}
