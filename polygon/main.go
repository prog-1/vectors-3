package main

import (
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type vector struct {
	x, y int
}

type line struct {
	a, b vector
}

func (l *line) normalFormula(a vector) int {
	x, y := a.x, a.y
	dx, dy := l.b.x-l.a.x, l.b.y-l.a.y
	A, B, C := -dy, dx, dy*l.a.x-dx*l.a.y
	return A*x + B*y + C
}

func (l *line) draw(screen *ebiten.Image) {
	DrawLine(screen, l.a.x, l.a.y, l.b.x, l.b.y, color.RGBA{1, 100, 100, 255})
}

const (
	screenWidth  = 1920
	screenHeight = 1080
)

type game struct {
	l      []line
	p      vector
	buffer *ebiten.Image
}

func (g *game) Layout(outWidth, outHeight int) (w, h int) { return screenWidth, screenHeight }
func (g *game) Update() error {
	g.p = vector{rand.Intn(screenWidth + 1), rand.Intn(screenHeight + 1)}
	return nil
}
func (g *game) Draw(screen *ebiten.Image) {
	//g.p = vector{300, 800}
	for _, l := range g.l {
		l.draw(g.buffer)
	}

	n1 := g.l[0].normalFormula(g.p)
	n2 := g.l[1].normalFormula(g.p)
	n3 := g.l[2].normalFormula(g.p)
	n4 := g.l[3].normalFormula(g.p)

	var clr color.Color
	if n1 > 0 && n2 > 0 && n3 > 0 && n4 > 0 {
		clr = color.RGBA{0, 255, 255, 255}
	} else {
		clr = color.RGBA{255, 165, 0, 255}
	}

	g.buffer.Set(g.p.x, g.p.y, clr)

	screen.DrawImage(g.buffer, &ebiten.DrawImageOptions{})
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ebiten.SetWindowSize(screenWidth, screenHeight)
	g := &game{}
	g.buffer = ebiten.NewImage(screenWidth, screenHeight)

	g.l = make([]line, 4)
	// g.l[0] = line{vector{200, 1000}, vector{200, 200}}                // left
	// g.l[1] = line{vector{200, 200}, vector{1000, 200}}                // top
	// g.l[2] = line{vector{1000, 200}, vector{screenWidth - 200, 1000}} // right
	// g.l[3] = line{vector{screenWidth - 200, 1000}, vector{200, 1000}} // bottom

	g.l[0] = line{vector{100, screenHeight - 50}, vector{140, screenHeight / 2}}                                   // left
	g.l[1] = line{vector{140, screenHeight / 2}, vector{screenWidth/2 + 100, screenHeight / 2}}                    // top
	g.l[2] = line{vector{screenWidth/2 + 100, screenHeight / 2}, vector{screenWidth/2 + 140, screenHeight/2 + 60}} // right
	g.l[3] = line{vector{screenWidth/2 + 140, screenHeight/2 + 60}, vector{100, screenHeight - 50}}                // bottom

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
func DrawLine(img *ebiten.Image, x1, y1, x2, y2 int, c color.Color) {
	// abs(Dy) < abs(dx) | / abs(dx) => abs(Dy)/abs(Dx) < 1 == abs(k) < 1
	if abs(y2-y1) < abs(x2-x1) {
		if x1 > x2 {
			x1, x2 = x2, x1
			y1, y2 = y2, y1
		}
		dx, dy := x2-x1, y2-y1
		dirY := 1
		// Dy < 0 => y2 - y1 < 0 => y1 > y2 => Growing downwards
		if dy < 0 {
			dirY = -1
			dy = -dy // For us to pretend that line is growing upwards
		}
		d := 2*dy - dx
		for x, y := x1, y1; x < x2; x++ {
			img.Set(x, y, c)
			if d >= 0 { // NE
				y += dirY
				d += dy - dx
			} else { // E
				d += dy
			}
		}
	} else {
		if y1 > y2 {
			x1, x2 = x2, x1
			y1, y2 = y2, y1
		}
		dx, dy := x2-x1, y2-y1
		dirX := 1
		if dx < 0 {
			dirX = -1
			dx = -dx
		}
		d := 2*dx - dy
		for x, y := x1, y1; y < y2; y++ {
			img.Set(x, y, c)
			if d >= 0 { // NE
				x += dirX
				d += dx - dy
			} else { // E
				d += dx
			}
		}
	}
}
