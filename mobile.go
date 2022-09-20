package mobile

import (
	"github.com/briansunter/twenty/pkg/ui"
	"github.com/hajimehoshi/ebiten/v2/mobile"
)

var (
	running bool
)

const (
	ScreenWidth  = 320
	ScreenHeight = 240
)

// IsRunning returns a boolean value indicating whether the game is running.
func IsRunning() bool {
	return running
}

// Start starts the game.
func Start(scale float64) error {
	running = true
	game := ui.NewGame()
	game.Initialize()

	// Prepare this function if you also want to make your game run on iOS.
	mobile.SetGame(game)
	// mobile.Start starts the game.
	// In this function, scale is passed from Java/Objecitve-C side
	// and pass it to mobile.Start. You can also receive the screen
	// size from Java/Objetive-C side by adding arguments to this
	// function if you want to make Java/Objective-C side decide the
	// screen size.
	// Note that the screen size unit is dp (device-independent pixel)
	// on Android.
	// if err := mobile.Start(game.Update, ScreenWidth, ScreenHeight, scale, "Your Game's Title"); err != nil {
	// 	return err
	// }
	return nil
}

func Dummy() {}
