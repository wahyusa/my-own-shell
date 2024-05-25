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
		default:
			fmt.Printf("%s: command not found\n", commandInput)
		}
	}
}
