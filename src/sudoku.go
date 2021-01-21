package sudoku

import (
	"time"
)

//Game struct
type Game struct {
	Board      *board
	Level      ComplexityLevel
	StartTime  time.Time
	complexity *complexity
}

//NewGame creates new game with the specific complexity
func NewGame(level ComplexityLevel) *Game {
	return &Game{
		Board: initGame(level),
		Level: level,
	}
}

func (g *Game) isSolved() bool {
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			if g.IsEmpty(x, y) {
				return false
			}
		}
	}
	return true
}

//Solve solves the current board
func (g *Game) Solve() {
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			if !g.IsEmpty(x, y) {
				continue
			}

			for n := 1; n < 10; n++ {
				if g.IsValid(x, y, n) {
					g.Set(x, y, n)
					g.Solve()
					if !g.isSolved() {
						g.Set(x, y, 0)
					}
				}
			}
			return

		}
	}
}

//Validate validates if the solutions is correct
func (g *Game) Validate() *ValidationErrors {
	errs := new(ValidationErrors)
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			switch {
			case g.IsEmpty(x, y):
				errs.appendError(emptyError, ErrorCoordinate{X: x, Y: y})
			case g.IsXValid(x, y, g.Get(x, y)):
				errs.appendError(invalidRow, ErrorCoordinate{X: x, Y: y})
			case g.IsYValid(x, y, g.Get(x, y)):
				errs.appendError(invalidColumn, ErrorCoordinate{X: x, Y: y})
			case g.IsSquareValid(x, y, g.Get(x, y)):
				errs.appendError(invalidSquare, ErrorCoordinate{X: x, Y: y})
			}
		}
	}

	return errs
}

//Set sets the value in the coordinate
func (g *Game) Set(x, y, n int) {
	g.Board.set(x, y, n)
}

//Get return the value in the coordinate
func (g *Game) Get(x, y int) int {
	return g.Board.get(x, y)
}

//IsXValid validate if n is valid in the row
func (g *Game) IsXValid(x, y, n int) bool {
	return g.Board.isXValid(x, y, n)
}

//IsYValid validate if n is valid in the column
func (g *Game) IsYValid(x, y, n int) bool {
	return g.Board.isYValid(x, y, n)
}

//IsSquareValid validate if n is valid in the square
func (g *Game) IsSquareValid(x, y, n int) bool {
	return g.Board.isSquareValid(x, y, n)
}

//IsValid validate if n is valid in the row, column and square
func (g *Game) IsValid(x, y, n int) bool {
	return g.Board.isValid(x, y, n)
}

//IsEmpty validate if the coordinate has value
func (g *Game) IsEmpty(x, y int) bool {
	return g.Board.isEmpty(x, y)
}
