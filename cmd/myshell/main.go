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

		if commandInput == "exit 0" {
			break
		}

		// If the command is not "exit 0", print "command not found"
		fmt.Printf("%s: command not found\n", commandInput)
	}
}
