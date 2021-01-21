package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CallCmd() {
	cmd := readInput("Type your next action (help for help): ")

	if action, ok := commandsMap[cmd]; ok {
		action()
		return
	}

	fmt.Println("Invalid option. Please, try again.")
}

func readInput(msg string) string {
	for {
		fmt.Print(msg)
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')

		if err == nil {
			return strings.ReplaceAll(strings.Trim(input, ""), "\n", "")
		}

		fmt.Println("incorrect input. try again.")
	}
}
