package ui

import (
	"github.com/briansunter/twenty/pkg/game"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	state      *game.State
	input      *Input
	frameCount int
	frameTouch int
}

func NewGame() *Game {
	return &Game{
		state: game.NewState(),
		input: NewInput(),
	}
}
func (g *Game) Initialize() {
	g.state.Initialize()
}

func (g *Game) CanMove() bool {
	return g.frameCount > g.frameTouch+25
}

func (g *Game) Update() error {
	g.frameCount++
	g.input.Update()

	if g.CanMove() {
		if dir, ok := g.input.Dir(); ok {
			switch dir {
			case DirUp:
				g.state.MoveUp()
			case DirRight:
				g.state.MoveRight()
			case DirDown:
				g.state.MoveDown()
			case DirLeft:
				g.state.MoveLeft()
			}
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	DrawBoard(screen, g.state.Board)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 320
}
