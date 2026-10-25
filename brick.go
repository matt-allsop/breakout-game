package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var colorYellow = color.RGBA{R: 255, G: 215, B: 0, A: 255}
var colorGreen = color.RGBA{R: 0, G: 170, B: 0, A: 255}
var colorOrange = color.RGBA{R: 255, G: 136, B: 0, A: 255}
var colorRed = color.RGBA{R: 229, G: 37, B: 33, A: 255}

const (
	BrickWidth  = 60
	BrickHeight = 10
)

type Brick struct {
	X, Y  float64
	Hit   bool
	Color color.RGBA
}

func NewBrick(x, y float64, col color.RGBA) *Brick {
	return &Brick{X: x, Y: y, Hit: false, Color: col}
}

func (b *Brick) Draw(screen *ebiten.Image) {
	if !b.Hit {
		vector.FillRect(screen, float32(b.X), float32(b.Y), BrickWidth, BrickHeight, b.Color, false)
	}
}
