package main

import (
	"fmt"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 1000
	screenHeight = 800
	numLines     = 500
)

type Drawing struct {
	lines        []line
	currentWidth float32
	increasing   bool
}

type line struct {
	x1, y1, x2, y2, width float32
	color                 color.RGBA
}

// Operation is a function type that takes two float32 operands and returns a float32 result
type Operation func(float32, float32) float32

// Define operations
var operations = map[string]Operation{
	"+": func(a, b float32) float32 { return a + b },
	"-": func(a, b float32) float32 { return a - b },
	"*": func(a, b float32) float32 { return a * b },
	"/": func(a, b float32) float32 { return a / b },
}

const inc float32 = 1.2

func (d *Drawing) Update() error {
	var activeKeys []ebiten.Key
	keys := inpututil.AppendJustPressedKeys(activeKeys[:0])

	if d.currentWidth >= 200 {
		d.increasing = false
	}

	if d.currentWidth <= 7 && !d.increasing {
		d.increasing = true
	}
	var op string
	if d.increasing {
		op = "*"
	} else {
		op = "/"
	}
	if len(keys) > 0 {
		d.updateLines(op)
	}
	return nil
}

func (d *Drawing) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x99, 0xcc, 0xff, 0xff})

	randomStrokes := make([]float32, numLines)

	for s := range randomStrokes {
		randomStrokes[s] = float32(rand.Int31n(30))
	}
	for _, l := range d.lines {
		vector.StrokeLine(screen, l.x1, l.y1, l.x2, l.y2, l.width, l.color, true)
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("WIDTH: %f", d.currentWidth))
}

func (d *Drawing) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (d *Drawing) updateLines(op string) {

	operation := operations[op]
	d.currentWidth = operation(d.currentWidth, inc)
	for i, l := range d.lines {
		d.lines[i] = line{
			x1:    l.x1,
			y1:    l.y1,
			x2:    screenWidth - (l.x2 + rand.Float32()*screenWidth),
			y2:    screenHeight - (l.y2 + rand.Float32()*screenHeight),
			width: d.currentWidth,
			color: l.color,
		}

	}

}

func main() {
	d := &Drawing{
		lines:        make([]line, numLines),
		currentWidth: 7,
		increasing:   true,
	}

	for i := range d.lines {
		d.lines[i] = line{
			x1:    rand.Float32() * screenWidth,
			y1:    rand.Float32() * screenHeight,
			x2:    rand.Float32() * screenWidth,
			y2:    rand.Float32() * screenHeight,
			width: 7,
			color: randomColor(),
		}
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Line Art Generator")
	if err := ebiten.RunGame(d); err != nil {
		panic(err)
	}

}

func randomColor() color.RGBA {
	return color.RGBA{
		R: uint8(rand.Intn(256)),
		G: uint8(rand.Intn(256)),
		B: uint8(rand.Intn(256)),
		A: 255,
	}
}
