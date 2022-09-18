package main

import "math/rand"

type Board struct {
	size      int
	gameBoard [][]int
}

func (b *Board) initialize() {
	b.gameBoard = [][]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
}

func (g *Game) mergeLeft() {
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			if g.board.gameBoard[i][j] == g.board.gameBoard[i][j+1] {
				g.board.gameBoard[i][j] = g.board.gameBoard[i][j] * 2
				g.board.gameBoard[i][j+1] = 0
				g.board.size--
			}
		}
	}
}

func (g *Game) mergeRight() {
	for i := 0; i < 4; i++ {
		for j := 3; j > 0; j-- {
			if g.board.gameBoard[i][j] == g.board.gameBoard[i][j-1] {
				g.board.gameBoard[i][j] = g.board.gameBoard[i][j] * 2
				g.board.gameBoard[i][j-1] = 0
				g.board.size--
			}
		}
	}
}

func (g *Game) mergeUp() {
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			if g.board.gameBoard[j][i] == g.board.gameBoard[j+1][i] {
				g.board.gameBoard[j][i] = g.board.gameBoard[j][i] * 2
				g.board.gameBoard[j+1][i] = 0
				g.board.size--
			}
		}
	}
}

func (g *Game) mergeDown() {
	for i := 0; i < 4; i++ {
		for j := 3; j > 0; j-- {
			if g.board.gameBoard[j][i] == g.board.gameBoard[j-1][i] {
				g.board.gameBoard[j][i] = g.board.gameBoard[j][i] * 2
				g.board.gameBoard[j-1][i] = 0
				g.board.size--
			}
		}
	}
}

func (g *Game) moveLeft() {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if g.board.gameBoard[i][j] == 0 {
				for k := j; k < 4; k++ {
					if g.board.gameBoard[i][k] != 0 {
						g.board.gameBoard[i][j] = g.board.gameBoard[i][k]
						g.board.gameBoard[i][k] = 0
						break
					}
				}
			}
		}
	}
}

func (g *Game) moveRight() {
	for i := 0; i < 4; i++ {
		for j := 3; j >= 0; j-- {
			if g.board.gameBoard[i][j] == 0 {
				for k := j; k >= 0; k-- {
					if g.board.gameBoard[i][k] != 0 {
						g.board.gameBoard[i][j] = g.board.gameBoard[i][k]
						g.board.gameBoard[i][k] = 0
						break
					}
				}
			}
		}
	}
}
func (g *Game) moveUp() {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if g.board.gameBoard[i][j] == 0 {
				for k := i; k < 4; k++ {
					if g.board.gameBoard[k][j] != 0 {
						g.board.gameBoard[i][j] = g.board.gameBoard[k][j]
						g.board.gameBoard[k][j] = 0
						break
					}
				}
			}
		}
	}
}

func (g *Game) moveDown() {
	for i := 3; i >= 0; i-- {
		for j := 0; j < 4; j++ {
			if g.board.gameBoard[i][j] == 0 {
				for k := i; k >= 0; k-- {
					if g.board.gameBoard[k][j] != 0 {
						g.board.gameBoard[i][j] = g.board.gameBoard[k][j]
						g.board.gameBoard[k][j] = 0
						break
					}
				}
			}
		}
	}
}

func cloneBoard(b *Board) *Board {
	newBoard := &Board{}
	newBoard.initialize()
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			newBoard.gameBoard[i][j] = b.gameBoard[i][j]
		}
	}
	newBoard.size = b.size
	return newBoard
}

func (b *Board) deepEqual(other *Board) bool {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if b.gameBoard[i][j] != other.gameBoard[i][j] {
				return false
			}
		}
	}
	return true
}

func (g *Game) tryAddBlock() {
	x := rand.Intn(4)
	y := rand.Intn(4)
	if g.board.size < 16 {
		if g.board.gameBoard[x][y] == 0 {
			g.board.gameBoard[x][y] = 2
			g.board.size++
		} else {
			g.tryAddBlock()
		}
	} else {
		// Game over
		g.state = Lose
	}
}
