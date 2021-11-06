package concepts

// Backtracking uses recursion to
// incrementally build a possible solutions
// and abandons paths that are not
// possible solutions

type NQueenSolver struct {
	board        [][]int
	size         int
	queensPlaced int
}

func NewNQueenSolver(n int) *NQueenSolver {
	solver := &NQueenSolver{
		size: n,
	}

	for i := 0; i < n; i++ {
		solver.board = append(solver.board, make([]int, n))
	}

	return solver
}

func (s *NQueenSolver) Solve() [][]int {
	if s.placeQueen(0) {
		return s.board
	}
	return nil
}

func (s *NQueenSolver) placeQueen(col int) bool {
	// All queens are placed, success!
	if s.queensPlaced == s.size {
		return true
	}

	for i := 0; i < s.size; i++ {
		// check first to see if we can place
		// the queen
		if s.isConflict(col, i) {
			continue
		}

		// place it
		s.board[col][i] = 1
		s.queensPlaced++

		// start placing the next column
		if s.placeQueen(col + 1) {
			return true
		}

		// Failure, remove the piece
		s.board[col][i] = 0
		s.queensPlaced--
	}

	return false
}

func (s *NQueenSolver) isConflict(x, y int) bool {
	// check row
	for i := x - 1; i >= 0; i-- {
		if s.board[i][y] == 1 {
			return true
		}
	}

	// check column
	for i := y - 1; i >= 0; i-- {
		if s.board[x][i] == 1 {
			return true
		}
	}

	// check diag left
	i, j := x-1, y-1
	for {
		if i < 0 || j < 0 {
			break
		}
		if s.board[i][j] == 1 {
			return true
		}
		i--
		j--
	}

	// check diag right
	i, j = x-1, y+1
	for {
		if i < 0 || j > s.size-1 {
			break
		}
		if s.board[i][j] == 1 {
			return true
		}
		i--
		j++
	}

	return false
}
