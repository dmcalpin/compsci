package concepts

// Backtracking uses recursion to
// incrementally build a possible solutions
// and abandons paths that are not
// possible solutions

type Coord struct {
	X int
	Y int
}

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
	// try to place in this column
	for i := 0; i < s.size; i++ {
		s.board[col][i] = 1
		s.queensPlaced++
		if s.isConflict(&Coord{col, i}) {
			s.board[col][i] = 0
			s.queensPlaced--
		} else {
			// if we place them all, we've succeded
			if s.queensPlaced == s.size {
				return true
			}
			success := s.placeQueen(col + 1)
			if success {
				return true
			}
			s.board[col][i] = 0
			s.queensPlaced--
		}

	}

	return false
}

func (s *NQueenSolver) isConflict(c *Coord) bool {
	// check row
	for i := c.X - 1; i >= 0; i-- {
		if s.board[i][c.Y] == 1 {
			return true
		}
	}

	// check column
	for i := c.Y - 1; i >= 0; i-- {
		if s.board[c.X][i] == 1 {
			return true
		}
	}

	// check diag left
	i, j := c.X-1, c.Y-1
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
	i, j = c.X-1, c.Y+1
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
