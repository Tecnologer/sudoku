package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/tecnologer/sudoku/clients/cli/game"
	sudoku "github.com/tecnologer/sudoku/src"
)

var (
	commandsMap map[string]func()
	commands    []*command
)

func init() {
	commands = []*command{
		{
			cmd:    "help",
			action: printHelp,
			about:  "Prints the list of commands available",
			alias:  []string{"?", "h"},
		},
		{
			cmd:    "exit",
			action: exit,
			about:  "Closes the game",
			alias:  []string{"close"},
		},
		{
			cmd:    "new",
			action: newGame,
			about:  "Starts new game",
			alias:  []string{},
		},
		{
			cmd:    "set",
			action: setValue,
			about:  "Sets the value in the coordinate",
			alias:  []string{},
		},
		{
			cmd:    "display",
			action: printGame,
			about:  "Sets the value in the coordinate",
			alias:  []string{"print", "show"},
		},
	}

	commandsMap = make(map[string]func())
	for _, cmd := range commands {
		commandsMap[cmd.cmd] = cmd.action
		for _, alias := range cmd.alias {
			if _, exists := commandsMap[alias]; exists {
				continue
			}

			commandsMap[alias] = cmd.action
		}
	}
}

func printHelp() {
	fmt.Println("\nAvailable commands: ")
	for _, cmd := range commands {
		fmt.Println("\t", cmd)
	}
	fmt.Println()
}

func exit() {
	i := "y"
	// i := readInput("Are you sure? (Y/n): ")

	if strings.ToLower(i) == "y" {
		os.Exit(0)
	}
}

func newGame() {
	fmt.Println("Difficulties availables: ")
	for _, c := range sudoku.GetComplexities() {
		fmt.Printf("\t- %s\n", c)
	}

	fmt.Println()
	levelStr := readInput("Choose you difficulty: ")

	level := sudoku.StringToComplexity(levelStr)

	game.Current = sudoku.NewGame(level)

	fmt.Printf("new %s game started\n", level)
}

func setValue() {
	input := ""
	for input != "cancel" {
		input = readInput("Type the coordinate and the value separate by commas (x,y,z): ")
		inputs := strings.Split(input, ",")
		if len(inputs) != 3 {
			fmt.Println("Invalid data. Please, try again")
			continue
		}

		row, err := strconv.Atoi(inputs[0])
		if err != nil || row > 9 || row < 1 {
			fmt.Println("The value for the row is invalid. Should be a integer between 1 and 9")
			continue
		}

		col, err := strconv.Atoi(inputs[1])
		if err != nil || col > 9 || col < 1 {
			fmt.Println("The value for the column is invalid. Should be a integer between 1 and 9")
			continue
		}

		val, err := strconv.Atoi(inputs[1])
		if err != nil || val > 9 || val < 1 {
			fmt.Println("The value is invalid. Should be a integer between 1 and 9")
			continue
		}

		game.Current.Set(row-1, col-1, val)
	}

}

func printGame() {
	for x := 0; x < 9; x++ {
		fmt.Print("|")
		for y := 0; y < 9; y++ {
			if game.Current.IsEmpty(x, y) {
				fmt.Print(game.Current.Get(x, y))
			} else {
				fmt.Print("-")
			}
			fmt.Print(" ")

			if (y+1)%3 == 0 {
				fmt.Print("|")
			}
		}
		if (x+1)%3 == 0 && x < 8 {
			fmt.Println("\n|------|------|------|")
		} else {
			fmt.Println()
		}
	}
}
