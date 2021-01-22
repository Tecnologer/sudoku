package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tecnologer/sudoku/clients/cli/cmd"
	"github.com/tecnologer/sudoku/clients/cli/game"
	sudoku "github.com/tecnologer/sudoku/src"
)

var (
	verbouse   = flag.Bool("v", false, "verbouse")
	version    string
	minversion string
)

func main() {
	flag.Parse()

	if len(os.Args) > 1 && os.Args[1] == "version" {
		fmt.Printf("%s%s\n", version, minversion)
		return
	}

	for {
		cmd.CallCmd("")
	}

	// defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	// defer profile.Start(profile.MemProfile, profile.MemProfileRate(1), profile.ProfilePath(".")).Stop()
	l := sudoku.MasterLevel
	fmt.Printf("creating new game as %s\n", l)
	game.Current = sudoku.NewGame(l)
	// cmd.CallCmd("new", "easy")
	// cmd.CallCmd("solve")
	cmd.CallCmd("print")

	game.Current.Solve()
	fmt.Println("solved")
	cmd.CallCmd("print")

	// err := game.Validate()
	// fmt.Println("validated")

	// if err.Count > 0 {
	// 	for t, errs := range err.Errs {
	// 		for _, e := range errs {
	// 			fmt.Printf("Error type:%s at (%d,%d)\n", t, e.X, e.Y)
	// 		}
	// 	}
	// 	return
	// }

	// fmt.Println("you won!")
}

func printCmds() {
	fmt.Println("Sudoku-CLI provides the following commands to start")
	fmt.Println()
	fmt.Println("  version", "\n\treturns the current version")
}
