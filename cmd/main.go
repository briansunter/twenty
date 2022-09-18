package main

import (
	"fmt"
	"log"

	"github.com/briansunter/twenty/pkg/game"
	"github.com/briansunter/twenty/pkg/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	state      *game.State
	touchState *ui.TouchState
}

func (g *Game) Update() error {
	ui.HandleInput(g.state)
	g.touchState.HandleTouches()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ui.DrawBoard(screen, g.state.Board)
	//print touch state
	ebitenutil.DebugPrint(screen, fmt.Sprintf("x: %f, y: %f, zoom: %f", g.touchState.X, g.touchState.Y, g.touchState.Zoom))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	touchState := &ui.TouchState{
		X:        0,
		Y:        0,
		Zoom:     0,
		TouchIDs: []ebiten.TouchID{},
		Touches:  map[ebiten.TouchID]*ui.Touch{},
		Pinch:    &ui.Pinch{},
		Pan:      &ui.Pan{},
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
