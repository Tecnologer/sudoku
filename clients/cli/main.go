package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tecnologer/sudoku/clients/cli/cmd"
)

var (
	verbouse   = flag.Bool("v", false, "verbouse")
	version    string
	minversion string
)

func main() {
	flag.Parse()

	if len(os.Args) < 2 {
		printCmds()
		return
	}

	if os.Args[1] == "version" {
		fmt.Printf("%s%s\n", version, minversion)
		return
	}

	if os.Args[1] != "new" {
		printCmds()
		return
	}

	for {
		cmd.CallCmd()
	}

	// defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	// defer profile.Start(profile.MemProfile, profile.MemProfileRate(1), profile.ProfilePath(".")).Stop()
	// l := sudoku.MasterLevel
	// fmt.Printf("creating new game as %s\n", l)
	// game := sudoku.NewGame(l)
	// printGame(game)

	// game.Solve()
	// fmt.Println("solved")
	// printGame(game)

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
	fmt.Println("  new", "\n\tcreates a new game")
	fmt.Println("  version", "\n\treturns the current version")
}
