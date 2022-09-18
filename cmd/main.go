package main

import (
	"log"

	"github.com/briansunter/twenty/pkg/game"
	"github.com/briansunter/twenty/pkg/ui"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	state      *game.State
	touchState *ui.TouchState
	frameCount int
	frameTouch int
}

func (g *Game) CanMove() bool {
	return g.frameCount > g.frameTouch+25
}

func (g *Game) Update() error {
	g.frameCount++
	ui.HandleInput(g.state)
	g.touchState.HandleTouches()
	if g.touchState.Pan != nil && g.CanMove() {
		g.frameTouch = g.frameCount

		switch g.touchState.Pan.Direction {
		case ui.Left:
			g.state.MoveLeft()
		case ui.Right:
			g.state.MoveRight()
		case ui.Up:
			g.state.MoveUp()
		case ui.Down:
			g.state.MoveDown()

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
	touchState := &ui.TouchState{
		X:        0,
		Y:        0,
		Zoom:     0,
		TouchIDs: []ebiten.TouchID{},
		Touches:  map[ebiten.TouchID]*ui.Touch{},
		Pinch:    &ui.Pinch{},
		Taps:     []ui.Tap{},
	}
	game := &Game{
		touchState: touchState,
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
