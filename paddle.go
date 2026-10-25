package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var colorBlue = color.RGBA{R: 40, G: 100, B: 255, A: 255}

type Paddle struct {
	X, Y   float64
	Width  float64
	Height float64
	Speed  float64
}

func NewPaddle() *Paddle {
	return &Paddle{
		X:      350,
		Y:      540,
		Width:  100,
		Height: 10,
		Speed:  5,
	}
}

func (p *Paddle) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.X -= p.Speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.X += p.Speed
	}

	if p.X < 0 {
		p.X = 0
	}
	if p.X > 800-p.Width {
		p.X = 800 - p.Width
	}
}

func (p *Paddle) Draw(screen *ebiten.Image) {
	vector.FillRect(screen, float32(p.X), float32(p.Y), float32(p.Width), float32(p.Height), colorBlue, false)
}
