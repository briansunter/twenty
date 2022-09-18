package ui

import (
	"fmt"
	"image/color"

	"github.com/briansunter/twenty/pkg/game"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func DrawBoard(screen *ebiten.Image, board *game.Board) {
	for i, row := range board.GameBoard {
		for j, _ := range row {
			image := ebiten.NewImage(32, 32)
			image.Fill(color.RGBA{0, 0, 0, 255})
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(i*32), float64(j*32))
			screen.DrawImage(image, op)
			ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%d", board.GameBoard[j][i]), i*32, j*32)
		}
	}
}