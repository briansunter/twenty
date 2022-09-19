package main

import (
	"log"

	"github.com/briansunter/twenty/pkg/game"
	"github.com/briansunter/twenty/pkg/ui"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	state      *game.State
	input      *ui.Input
	frameCount int
	frameTouch int
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
			case ui.DirUp:
				g.state.MoveUp()
			case ui.DirRight:
				g.state.MoveRight()
			case ui.DirDown:
				g.state.MoveDown()
			case ui.DirLeft:
				g.state.MoveLeft()
			}
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ui.DrawBoard(screen, g.state.Board)
	//print touch state
	// ebitenutil.DebugPrintAt(screen, fmt.Sprintf("panDirection: %v", g.currentMove), 20, 200)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(320, 240)
	ebiten.SetWindowTitle("Hello, World!")
	game := &Game{
		input: ui.NewInput(),
		state: &game.State{
			Board: &game.Board{},
			Seed:  0,
			State: 0,
		},
	}
	game.state.Initialize()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
