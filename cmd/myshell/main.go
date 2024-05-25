package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func commandRunner(commandParts []string) {
	cmd := exec.Command(commandParts[0], commandParts[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		if _, ok := err.(*exec.Error); ok {
			fmt.Printf("%s: command not found\n", commandParts[0])
		} else {
			fmt.Println(err.Error())
		}
	}
}

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
		commandParts := strings.Split(commandInput, " ")

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
				found := false
				for _, path := range strings.Split(os.Getenv("PATH"), string(os.PathListSeparator)) {
					_, err := os.Stat(filepath.Join(path, commandType))
					if _, ok := err.(*os.PathError); ok {
						continue
					}
					fmt.Printf("%s is %s\n", commandType, filepath.Join(path, commandType))
					found = true
					break
				}
				if !found {
					fmt.Printf("%s not found\n", commandType)
				}
			}
		default:
			commandRunner(commandParts)
		}
	}
}
