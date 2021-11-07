package concepts

import "fmt"

// Sudoku is a puzzle where a 9x9
// grid is pre-filled with some amount
// of numbers and the blank cells are
// to be filled in by the solver. This
// grid is further divided into 3x3
// sections.
// The rules are that no column or row can
// contain a duplicate number, and also
// no 3x3 section of the grid can contain
// a duplicate

type SudokuSolver struct {
	board [][]int
}

// Creates a new Sudoku solver
func NewSudokuSolver(startingBoard [][]int) *SudokuSolver {
	return &SudokuSolver{
		board: startingBoard,
	}
}

// Begins solving the puzzle. Will return
// the solved board if successful, otherwise
// nil
func (s *SudokuSolver) Solve() [][]int {
	if s.placeNumber(0, 0) {
		return s.board
	}
	return nil
}

// Returns a string representation of the
// board
func (s *SudokuSolver) String() string {
	return fmt.Sprintf("%v", s.board)
}

// attempts to place a number in a given
// cell, returns false if it does not lead
// to a correct solution
func (s *SudokuSolver) placeNumber(col, row int) bool {
	// made it to the end, success
	if row == 9 {
		return true
	}

	// cell not empty
	if s.board[col][row] != 0 {
		// place a number
		// in the next column or next row
		if col == 8 {
			return s.placeNumber(0, row+1)
		}
		return s.placeNumber(col+1, row)
	}

	// cell is empty, try and place a number
	for n := 1; n < 10; n++ {
		// check if it conflicts
		if s.isConflict(col, row, n) {
			continue
		}

		// put it on the board
		s.board[col][row] = n

		// place a number
		// in the next column or next row
		if col == 8 {
			if s.placeNumber(0, row+1) {
				return true
			}
		} else {
			if s.placeNumber(col+1, row) {
				return true
			}
		}
	}

	s.board[col][row] = 0
	return false
}

func (s *SudokuSolver) isConflict(col, row, num int) bool {
	// check row
	for i := 0; i < 9; i++ {
		if s.board[col][i] == num {
			return true
		}
	}

	// check column
	for i := 0; i < 9; i++ {
		if s.board[i][row] == num {
			return true
		}
	}

	// check 3x3 group
	colStart := col / 3
	rowStart := row / 3
	for i := colStart * 3; i < (colStart+1)*3; i++ {
		for j := rowStart * 3; j < (rowStart+1)*3; j++ {
			if s.board[i][j] == num {
				return true
			}
		}
	}

	return false
}
