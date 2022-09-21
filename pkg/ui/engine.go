package ui

import (
	"image/color"

	"github.com/briansunter/twenty/pkg/game"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/furex/v2"
	"github.com/yohamta/furex/v2/components"
)

type Game struct {
	init       bool
	state      *game.State
	input      *Input
	gameUI     *furex.View
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
	if !g.init {
		g.init = true
		g.setupUI()
	}
	g.gameUI.Update()

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
	g.gameUI.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 420
}

func (g *Game) setupUI() {
	newButton := func() *furex.View {
		return (&furex.View{
			Width:   75,
			Height:  50,
			Handler: &components.Box{Color: color.Gray{}},
			Justify: furex.JustifyCenter,
		}).AddChild(
			&furex.View{
				Width:  75,
				Height: 50,
				Handler: &components.Button{
					Text:    "New Game",
					OnClick: func() { g.state.Initialize() },
				},
			},
		)
	}

	g.gameUI = (&furex.View{
		Width:      320,
		Height:     420,
		Direction:  furex.Column,
		Justify:    furex.JustifySpaceBetween,
		AlignItems: furex.AlignItemCenter,
	}).AddChild(
		(&furex.View{
			Width:      320 - 20,
			Height:     70,
			Justify:    furex.JustifySpaceBetween,
			AlignItems: furex.AlignItemCenter,
		}).AddChild(
			&furex.View{
				Width:  80,
				Height: 100,
			},
			&furex.View{
				Width:  80,
				Height: 60,
			},
			(&furex.View{
				Width:       60,
				Height:      40,
				MarginRight: 20,
			}).AddChild(newButton()),
		),
	)
}
