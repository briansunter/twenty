package ui

import (
	_ "image/jpeg"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// distance between points a and b.
func distance(xa, ya, xb, yb int) float64 {
	x := math.Abs(float64(xa - xb))
	y := math.Abs(float64(ya - yb))
	return math.Sqrt(x*x + y*y)
}

const (
	screenWidth  = 320
	screenHeight = 240
)

var (
	gophersImage *ebiten.Image
)

type Touch struct {
	originX, originY int
	currX, currY     int
	duration         int
	wasPinch, isPan  bool
}

type Pinch struct {
	id1, id2 ebiten.TouchID
	originH  float64
	prevH    float64
}

type Pan struct {
	id ebiten.TouchID

	prevX, prevY     int
	originX, originY int
}

type Tap struct {
	X, Y int
}

type TouchState struct {
	X, Y float64
	Zoom float64

	TouchIDs []ebiten.TouchID
	Touches  map[ebiten.TouchID]*Touch
	Pinch    *Pinch
	Pan      *Pan
	Taps     []Tap
}

func (g *TouchState) HandleTouches() error {
	// Clear the previous frame's taps.
	g.Taps = g.Taps[:0]

	// What touches have just ended?
	for id, t := range g.Touches {
		if inpututil.IsTouchJustReleased(id) {
			if g.Pinch != nil && (id == g.Pinch.id1 || id == g.Pinch.id2) {
				g.Pinch = nil
			}
			if g.Pan != nil && id == g.Pan.id {
				g.Pan = nil
			}

			// If this one has not been touched long (30 frames can be assumed
			// to be 500ms), or moved far, then it's a tap.
			diff := distance(t.originX, t.originY, t.currX, t.currY)
			if !t.wasPinch && !t.isPan && (t.duration <= 30 || diff < 2) {
				g.Taps = append(g.Taps, Tap{
					X: t.currX,
					Y: t.currY,
				})
			}

			delete(g.Touches, id)
		}
	}

	// What touches are new in this frame?
	g.TouchIDs = inpututil.AppendJustPressedTouchIDs(g.TouchIDs[:0])
	for _, id := range g.TouchIDs {
		x, y := ebiten.TouchPosition(id)
		g.Touches[id] = &Touch{
			originX: x, originY: y,
			currX: x, currY: y,
		}
	}

	g.TouchIDs = ebiten.AppendTouchIDs(g.TouchIDs[:0])

	// Update the current position and durations of any touches that have
	// neither begun nor ended in this frame.
	for _, id := range g.TouchIDs {
		t := g.Touches[id]
		t.duration = inpututil.TouchPressDuration(id)
		t.currX, t.currY = ebiten.TouchPosition(id)
	}

	// Interpret the raw touch data that's been collected into g.touches into
	// gestures like two-finger pinch or single-finger pan.
	switch len(g.Touches) {
	case 2:
		// Potentially the user is making a pinch gesture with two fingers.
		// If the diff between their origins is different to the diff between
		// their currents and if these two are not already a pinch, then this is
		// a new pinch!
		id1, id2 := g.TouchIDs[0], g.TouchIDs[1]
		t1, t2 := g.Touches[id1], g.Touches[id2]
		originDiff := distance(t1.originX, t1.originY, t2.originX, t2.originY)
		currDiff := distance(t1.currX, t1.currY, t2.currX, t2.currY)
		if g.Pinch == nil && g.Pan == nil && math.Abs(originDiff-currDiff) > 3 {
			t1.wasPinch = true
			t2.wasPinch = true
			g.Pinch = &Pinch{
				id1:     id1,
				id2:     id2,
				originH: originDiff,
				prevH:   originDiff,
			}
		}
	case 1:
		// Potentially this is a new pan.
		id := g.TouchIDs[0]
		t := g.Touches[id]
		if !t.wasPinch && g.Pan == nil && g.Pinch == nil {
			diff := math.Abs(distance(t.originX, t.originY, t.currX, t.currY))
			if diff > 1 {
				t.isPan = true
				g.Pan = &Pan{
					id:      id,
					originX: t.originX,
					originY: t.originY,
					prevX:   t.originX,
					prevY:   t.originY,
				}
			}
		}
	}

	// Copy any active pinch gesture's movement to the Game's zoom.
	if g.Pinch != nil {
		x1, y1 := ebiten.TouchPosition(g.Pinch.id1)
		x2, y2 := ebiten.TouchPosition(g.Pinch.id2)
		curr := distance(x1, y1, x2, y2)
		delta := curr - g.Pinch.prevH
		g.Pinch.prevH = curr

		g.Zoom += (delta / 100) * g.Zoom
		if g.Zoom < 0.25 {
			g.Zoom = 0.25
		} else if g.Zoom > 10 {
			g.Zoom = 10
		}
	}

	// Copy any active pan gesture's movement to the Game's x and y pan values.
	if g.Pan != nil {
		currX, currY := ebiten.TouchPosition(g.Pan.id)
		deltaX, deltaY := currX-g.Pan.prevX, currY-g.Pan.prevY

		g.Pan.prevX, g.Pan.prevY = currX, currY

		g.X += float64(deltaX)
		g.Y += float64(deltaY)
	}

	// If the user has tapped, then reset the Game's pan and zoom.
	if len(g.Taps) > 0 {
		g.X = screenWidth / 2
		g.Y = screenHeight / 2
		g.Zoom = 1.0
	}
	return nil
}

func (g *TouchState) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	// Apply zoom.
	op.GeoM.Scale(g.Zoom, g.Zoom)

	// Apply pan.
	op.GeoM.Translate(g.X, g.Y)

	// Center the image (corrected by the current zoom).
	w, h := gophersImage.Size()
	op.GeoM.Translate(float64(-w)/2*g.Zoom, float64(-h)/2*g.Zoom)

	screen.DrawImage(gophersImage, op)

	ebitenutil.DebugPrint(screen, "Use a two finger pinch to zoom, swipe with one finger to pan, or tap to reset the view")
}

func (g *TouchState) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

// func main() {
// 	// Decode an image from the image file's byte slice.
// 	img, _, err := image.Decode(bytes.NewReader(images.Gophers_jpg))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	gophersImage = ebiten.NewImageFromImage(img)

// 	g := &TouchState{
// 		x:    screenWidth / 2,
// 		y:    screenHeight / 2,
// 		zoom: 1.0,

// 		touches: map[ebiten.TouchID]*touch{},
// 	}

// 	ebiten.SetWindowSize(screenWidth, screenHeight)
// 	ebiten.SetWindowTitle("Touch (Ebitengine Demo)")
// 	if err := ebiten.RunGame(g); err != nil {
// 		log.Fatal(err)
// 	}
// }
