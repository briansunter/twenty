package main

import (
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) initialize() {
	rand.Seed(g.seed)
	g.board.initialize()
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

func (g *Game) Update() error {
	oldBoard := cloneBoard(g.board)
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		//move left
		g.moveLeft()
		g.mergeLeft()
		g.moveLeft()

	} else if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		//move right
		g.moveRight()
		g.mergeRight()
		g.moveRight()
	} else if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		//move up
		g.moveUp()
		g.mergeUp()
		g.moveUp()
	} else if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		//move down
		g.moveDown()
		g.mergeDown()
		g.moveDown()
	} else if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		//add block
	} else if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.initialize()
		//quit
	}

	if !g.board.deepEqual(oldBoard) {
		g.tryAddBlock()
	}
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
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
