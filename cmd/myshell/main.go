package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Uncomment this block to pass the first stage
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	// bufio.NewReader(os.Stdin).ReadString('\n')

	// check if user input command is not reognized, then print error message and wait for user input again

	commandInput, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
	}

	commandInput = strings.TrimRight(commandInput, "\n")

	// Print "command not found"
	fmt.Printf("%s: command not found\n", commandInput)
}
