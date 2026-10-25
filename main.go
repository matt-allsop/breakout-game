package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	Paddle      *Paddle
	Ball        *Ball
	Bricks      []*Brick
	IsGameStart bool
	Score       int
	Lives       int
}

func (g *Game) ResetBall() {
	g.Ball.Stuck = true
	g.Ball.VX = 0
	g.Ball.VY = 0

	g.Paddle.X = 350
}

func (g *Game) BrickCollision() {
	for _, brick := range g.Bricks {
		if brick.Hit {
			continue
		}

		if g.Ball.X+g.Ball.Radius >= brick.X &&
			g.Ball.X-g.Ball.Radius <= brick.X+BrickWidth &&
			g.Ball.Y+g.Ball.Radius >= brick.Y &&
			g.Ball.Y-g.Ball.Radius <= brick.Y+BrickHeight {

			brick.Hit = true
			g.Ball.VY = -g.Ball.VY
			g.Score += 10

			break
		}
	}
}

func (g *Game) Update() error {
	g.Paddle.Update()
	g.Ball.Update()
	g.BrickCollision()

	if g.Ball.Y-g.Ball.Radius > 600 {
		g.Lives--
		g.ResetBall()
	}

	if g.Lives < 0 {
		g.ResetGame()
	}

	allHit := true
	for _, brick := range g.Bricks {
		if !brick.Hit {
			allHit = false
			break
		}
	}

	if allHit {
		g.ResetGame()
	}

	return nil
}

func (g *Game) ResetGame() {
	g.Score = 0
	g.Lives = 3
	g.ResetBall()
	g.BuildBricks()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Paddle.Draw(screen)
	g.Ball.Draw(screen)

	for _, brick := range g.Bricks {
		brick.Draw(screen)
	}

	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("LIVES: %d", g.Lives), 100, 20)

	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("SCORE: %d", g.Score), 20, 20)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 800, 600
}

func (g *Game) BuildBricks() {
	var bricks []*Brick
	rows := 8
	cols := 10
	startX := float64(75)
	startY := float64(50)
	spacing := float64(6)

	colors := []color.RGBA{colorRed, colorOrange, colorGreen, colorYellow}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			x := startX + float64(c)*(BrickWidth+spacing)
			y := startY + float64(r)*(BrickHeight+spacing)
			bricks = append(bricks, NewBrick(x, y, colors[r/2]))
		}
	}
	g.Bricks = bricks

}

func main() {
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Breakout Game")

	paddle := NewPaddle()

	game := &Game{
		Paddle: paddle,
		Ball:   NewBall(paddle),
		Lives:  3,
	}

	game.BuildBricks()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
