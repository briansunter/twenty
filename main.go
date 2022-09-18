package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Board struct {
	size      int
	gameBoard [][]int
}

func (g *Game) initialize() {
	rand.Seed(g.seed)
	g.board.initialize()
}

func (b *Board) initialize() {
	b.gameBoard = [][]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
}

func (board *Board) drawBoard(screen *ebiten.Image) {
	for i, row := range board.gameBoard {
		for j, _ := range row {
			image := ebiten.NewImage(32, 32)
			image.Fill(color.RGBA{0, 0, 0, 255})
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(i*32), float64(j*32))
			screen.DrawImage(image, op)
			ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%d", board.gameBoard[j][i]), i*32, j*32)
		}
	}
}

type GameState int

const (
	Loading GameState = iota
	New
	Running
	Lose
)

type Game struct {
	board *Board
	seed  int64
	state GameState
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

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.board.drawBoard(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	game := &Game{seed: 1, state: Running}
	game.board = &Board{}
	game.initialize()
	game.tryAddBlock()
	game.moveLeft()
	game.moveRight()
	game.moveDown()
	game.moveUp()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
