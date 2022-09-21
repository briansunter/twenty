package game

import "math/rand"

type Board struct {
	Size      int
	GameBoard [][]int
}

func (b *Board) initialize() {
	b.GameBoard = [][]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
}

func (b *Board) mergeLeft() {
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			if b.GameBoard[i][j] == b.GameBoard[i][j+1] {
				b.GameBoard[i][j] = b.GameBoard[i][j] * 2
				b.GameBoard[i][j+1] = 0
				b.Size--
			}
		}
	}
}

func (b *Board) mergeRight() {
	for i := 0; i < 4; i++ {
		for j := 3; j > 0; j-- {
			if b.GameBoard[i][j] == b.GameBoard[i][j-1] {
				b.GameBoard[i][j] = b.GameBoard[i][j] * 2
				b.GameBoard[i][j-1] = 0
				b.Size--
			}
		}
	}
}

func (b *Board) mergeUp() {
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			if b.GameBoard[j][i] == b.GameBoard[j+1][i] {
				b.GameBoard[j][i] = b.GameBoard[j][i] * 2
				b.GameBoard[j+1][i] = 0
				b.Size--
			}
		}
	}
}

func (b *Board) mergeDown() {
	for i := 0; i < 4; i++ {
		for j := 3; j > 0; j-- {
			if b.GameBoard[j][i] == b.GameBoard[j-1][i] {
				b.GameBoard[j][i] = b.GameBoard[j][i] * 2
				b.GameBoard[j-1][i] = 0
				b.Size--
			}
		}
	}
}

func (b *Board) moveLeft() {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if b.GameBoard[i][j] == 0 {
				for k := j; k < 4; k++ {
					if b.GameBoard[i][k] != 0 {
						b.GameBoard[i][j] = b.GameBoard[i][k]
						b.GameBoard[i][k] = 0
						break
					}
				}
			}
		}
	}
}

func (b *Board) moveRight() {
	for i := 0; i < 4; i++ {
		for j := 3; j >= 0; j-- {
			if b.GameBoard[i][j] == 0 {
				for k := j; k >= 0; k-- {
					if b.GameBoard[i][k] != 0 {
						b.GameBoard[i][j] = b.GameBoard[i][k]
						b.GameBoard[i][k] = 0
						break
					}
				}
			}
		}
	}
}

func (b *Board) moveUp() {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if b.GameBoard[i][j] == 0 {
				for k := i; k < 4; k++ {
					if b.GameBoard[k][j] != 0 {
						b.GameBoard[i][j] = b.GameBoard[k][j]
						b.GameBoard[k][j] = 0
						break
					}
				}
			}
		}
	}
}

func (b *Board) moveDown() {
	for i := 3; i >= 0; i-- {
		for j := 0; j < 4; j++ {
			if b.GameBoard[i][j] == 0 {
				for k := i; k >= 0; k-- {
					if b.GameBoard[k][j] != 0 {
						b.GameBoard[i][j] = b.GameBoard[k][j]
						b.GameBoard[k][j] = 0
						break
					}
				}
			}
		}
	}
}

func (b *Board) CloneBoard() *Board {
	newBoard := &Board{}
	newBoard.initialize()
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			newBoard.GameBoard[i][j] = b.GameBoard[i][j]
		}
	}
	newBoard.Size = b.Size
	return newBoard
}

func (b *Board) DeepEqual(other *Board) bool {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if b.GameBoard[i][j] != other.GameBoard[i][j] {
				return false
			}
		}
	}
	return true
}

func (b *Board) TryAddBlock() {
	x := rand.Intn(4)
	y := rand.Intn(4)
	if b.Size < 16 {
		if b.GameBoard[x][y] == 0 {
			b.GameBoard[x][y] = 2
			b.Size++
		} else {
			b.TryAddBlock()
		}
	} else {
		// Game over
	}
}

func (b *Board) CalculateScore() int {
	score := 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			score += b.GameBoard[i][j]
		}
	}
	return score
}
