package main

import (
	"log"

	"github.com/briansunter/twenty/pkg/game"
	"github.com/briansunter/twenty/pkg/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	state *game.State
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		//move left
		g.state.MoveLeft()

	} else if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		//move right
		g.state.MoveRight()
	} else if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		//move up
		g.state.MoveUp()
	} else if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		//move down
		g.state.MoveDown()
	} else if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		//add block
	} else if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.state.Initialize()
		//quit
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ui.DrawBoard(screen, g.state.Board)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	game := &Game{
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
