package main

import (
	"flag"
	"fmt"

	sudoku "github.com/tecnologer/sudoku/src"
)

var (
	reqVersion = flag.Bool("version", false, "returns the current version")
	version    string
	minversion string
)

func init() {
	flag.Parse()
}

func main() {
	if *reqVersion {
		fmt.Printf("%s.%s\n", version, minversion)
		return
	}
	// defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	// defer profile.Start(profile.MemProfile, profile.MemProfileRate(1), profile.ProfilePath(".")).Stop()
	l := sudoku.MasterLevel
	fmt.Printf("creating new game as %s\n", l)
	game := sudoku.NewGame(l)
	printGame(game)

	game.Solve()
	fmt.Println("solved")
	printGame(game)

	err := game.Validate()
	fmt.Println("validated")

	if err.Count > 0 {
		for t, errs := range err.Errs {
			for _, e := range errs {
				fmt.Printf("Error type:%s at (%d,%d)\n", t, e.X, e.Y)
			}
		}
		return
	}

	fmt.Println("you won!")
}

func printGame(game *sudoku.Game) {
	for x := 0; x < 9; x++ {
		fmt.Print("|")
		for y := 0; y < 9; y++ {
			if game.Board[x][y] != 0 {
				fmt.Print(game.Board[x][y])
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
