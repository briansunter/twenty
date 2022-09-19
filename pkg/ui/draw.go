package ui

import (
	"fmt"
	"image/color"
	"log"

	"github.com/briansunter/twenty/pkg/game"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var (
	mplusNormalFont     font.Face
	mplusSmallFont      font.Face
	mplusExtraSmallFont font.Face
	mplusSmallestFont   font.Face
)

func init() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 400
	mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    8,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	mplusSmallFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    6,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	mplusExtraSmallFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    4,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	mplusSmallestFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    3,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func DrawBoard(screen *ebiten.Image, board *game.Board) {
	for i, row := range board.GameBoard {
		for j, _ := range row {
			image := ebiten.NewImage(64, 64)
			image.Fill(color.RGBA{100, 0, 0, 255})
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(i*76), float64(j*76))
			// ebitenutil.DebugPrintAt(image, fmt.Sprintf("%d", board.GameBoard[j][i]), 8, 8)
			value := board.GameBoard[j][i]
			numX, numY := 16, 46
			font := mplusNormalFont
			if value > 1000 {
				numX = 10
				numY = 32
				font = mplusSmallestFont
			} else if value > 100 {
				numX = 12
				numY = 38
				font = mplusExtraSmallFont
			} else if value > 10 {
				numX = 10
				numY = 42
				font = mplusSmallFont
			}
			text.Draw(image, fmt.Sprintf("%d", board.GameBoard[j][i]), font, numX, numY, color.White)
			screen.DrawImage(image, op)
		}
	}
}
