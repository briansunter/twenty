package ui

import (
	"github.com/briansunter/twenty/pkg/game"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func HandleInput(g *game.State) error {
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		//move left
		g.MoveLeft()
	} else if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		//move right
		g.MoveRight()
	} else if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		//move up
		g.MoveUp()
	} else if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		//move down
		g.MoveDown()
	} else if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		//add block
	} else if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.Initialize()
		//quit
	}

	return nil
}
