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
	offset              = 12
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
	background := ebiten.NewImage(320, 320)
	background.Fill(color.RGBA{188, 172, 159, 255})
	screen.DrawImage(background, nil)
	for i, row := range board.GameBoard {
		for j, _ := range row {
			image := ebiten.NewImage(64, 64)
			bgImage := ebiten.NewImage(64, 64)
			bgImage.Fill(color.RGBA{204, 192, 179, 255})
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(offset+i*76), float64(offset+j*76))
			screen.DrawImage(bgImage, op)
			// ebitenutil.DebugPrintAt(image, fmt.Sprintf("%d", board.GameBoard[j][i]), 8, 8)
			value := board.GameBoard[j][i]
			numX, numY := 16, 46
			font := mplusNormalFont
			tileColor := color.RGBA{238, 225, 164, 255}
			if value > 1000 {
				numX = 10
				numY = 32
				font = mplusSmallestFont
				tileColor = color.RGBA{247, 95, 60, 255}
			} else if value > 100 {
				numX = 12
				numY = 38
				font = mplusExtraSmallFont
				tileColor = color.RGBA{247, 124, 95, 255}
			} else if value > 10 {
				numX = 10
				numY = 42
				font = mplusSmallFont
				tileColor = color.RGBA{242, 178, 120, 255}
			}
			image.Fill(tileColor)
			if value > 0 {
				text.Draw(image, fmt.Sprintf("%d", board.GameBoard[j][i]), font, numX, numY, color.White)
				screen.DrawImage(image, op)
			}
		}
	}
}
