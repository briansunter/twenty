package game

import "math/rand"

func (g *State) Initialize() {
	rand.Seed(g.Seed)
	g.Board.initialize()
	g.Board.TryAddBlock()
}

type Status int

const (
	Loading Status = iota
	New
	Running
	Lose
)

type State struct {
	Board *Board
	Seed  int64
	State Status
}

func (g *State) MoveUp() {
	oldBoard := g.Board.CloneBoard()
	g.Board.moveUp()
	g.Board.mergeUp()
	g.Board.moveUp()

	if !g.Board.DeepEqual(oldBoard) {
		g.Board.TryAddBlock()
	}

}

func (g *State) MoveDown() {
	oldBoard := g.Board.CloneBoard()
	g.Board.moveDown()
	g.Board.mergeDown()
	g.Board.moveDown()
	if !g.Board.DeepEqual(oldBoard) {
		g.Board.TryAddBlock()
	}
}

func (g *State) MoveLeft() {
	oldBoard := g.Board.CloneBoard()
	g.Board.moveLeft()
	g.Board.mergeLeft()
	g.Board.moveLeft()
	if !g.Board.DeepEqual(oldBoard) {
		g.Board.TryAddBlock()
	}
}

func (g *State) MoveRight() {
	oldBoard := g.Board.CloneBoard()
	g.Board.moveRight()
	g.Board.mergeRight()
	g.Board.moveRight()
	if !g.Board.DeepEqual(oldBoard) {
		g.Board.TryAddBlock()
	}
}
