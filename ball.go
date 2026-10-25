package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var colorWhite = color.RGBA{R: 255, G: 255, B: 255, A: 255}

type Ball struct {
	X, Y   float64
	VX, VY float64
	Radius float64
	Stuck  bool
	Paddle *Paddle
	Hit    bool
}

func NewBall(paddle *Paddle) *Ball {
	return &Ball{
		Radius: 8,
		Stuck:  true,
		Paddle: paddle,
	}
}

func (b *Ball) Update() {
	if b.Stuck {
		b.X = b.Paddle.X + (b.Paddle.Width / 2)
		b.Y = b.Paddle.Y - b.Radius - 2

		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			b.Stuck = false
			b.VX = 4.0
			b.VY = -4.0
		}
		return
	}

	b.X += b.VX
	b.Y += b.VY

	if b.X-b.Radius < 0 {
		b.X = b.Radius
		b.VX = -b.VX
	}

	if b.X+b.Radius > 800 {
		b.X = 800 - b.Radius
		b.VX = -b.VX
	}

	if b.Y-b.Radius < 0 {
		b.Y = b.Radius
		b.VY = -b.VY
	}

	horizontalOverlap := b.X+b.Radius >= b.Paddle.X && b.X-b.Radius <= b.Paddle.X+b.Paddle.Width

	verticalOverlap := b.Y+b.Radius >= b.Paddle.Y && b.Y-b.Radius <= b.Paddle.Y+b.Paddle.Height

	if horizontalOverlap && verticalOverlap && b.VY > 0 {
		b.VY = -b.VY
		b.Y = b.Paddle.Y - b.Radius
	}
}

func (b *Ball) Draw(screen *ebiten.Image) {
	vector.FillCircle(screen, float32(b.X), float32(b.Y), float32(b.Radius), colorWhite, false)
}
