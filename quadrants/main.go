package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type vector struct {
	x, y int
}

type line struct {
	a, b vector
}

func (l line) normalFormula(a vector) int {
	x, y := a.x, a.y
	dx, dy := l.a.x-l.b.x, l.a.y-l.b.y
	A, B, C := -dy, dx, dy*l.a.x-dx*l.a.y
	return A*x + B*y + C
}

const (
	screenWidth  = 1920
	screenHeight = 1080
)

type game struct {
	l1, l2 line
	p      vector
	buffer *ebiten.Image
}

func (g *game) Layout(outWidth, outHeight int) (w, h int) { return screenWidth, screenHeight }
func (g *game) Update() error {
	g.p = vector{rand.Intn(screenWidth + 1), rand.Intn(screenHeight + 1)}
	return nil
}
func (g *game) Draw(screen *ebiten.Image) {
	DrawLine(g.buffer, g.l1.a.x, g.l1.a.y, g.l1.b.x, g.l1.b.y, color.RGBA{1, 100, 100, 255})
	DrawLine(g.buffer, g.l2.a.x, g.l2.a.y, g.l2.b.x, g.l2.b.y, color.RGBA{100, 100, 255, 255})
	ebitenutil.DebugPrint(g.buffer, fmt.Sprintf("l1: a(%v, %v) b(%v, %v)    l2: a(%v, %v) b(%v, %v)", g.l1.a.x, g.l1.a.y, g.l1.b.x, g.l1.b.y, g.l2.a.x, g.l2.a.y, g.l2.b.x, g.l2.b.y))
	n1, n2 := g.l1.normalFormula(g.p), g.l2.normalFormula(g.p)
	if n1 > 0 && n2 > 0 {
		g.buffer.Set(g.p.x, g.p.y, color.RGBA{0, 255, 0, 255})
	} else if n1 > 0 {
		g.buffer.Set(g.p.x, g.p.y, color.RGBA{255, 0, 0, 255})
	} else if n2 > 0 {
		g.buffer.Set(g.p.x, g.p.y, color.RGBA{255, 255, 255, 255})
	} else {
		g.buffer.Set(g.p.x, g.p.y, color.RGBA{255, 165, 0, 255})
	}
	screen.DrawImage(g.buffer, &ebiten.DrawImageOptions{})
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ebiten.SetWindowSize(screenWidth, screenHeight)
	g := &game{}
	g.buffer = ebiten.NewImage(screenWidth, screenHeight)

	halfScreenWidth, halfScreenHeight := screenWidth/2, screenHeight/2

	d1 := rand.Float32()
	if d1 > 0.5 {
		g.l1.a = vector{rand.Intn(halfScreenWidth), 0}
		g.l1.b = vector{halfScreenWidth + rand.Intn(halfScreenWidth), screenHeight}
	} else {
		g.l1.a = vector{0, rand.Intn(screenHeight)}
		g.l1.b = vector{screenWidth, halfScreenHeight + rand.Intn(halfScreenHeight)}
	}

	d2 := rand.Float32()
	if d2 > 0.5 {
		g.l2.a = vector{halfScreenWidth + rand.Intn(halfScreenWidth), 0}
		g.l2.b = vector{rand.Intn(halfScreenWidth), screenHeight}
	} else {
		g.l2.a = vector{screenWidth, rand.Intn(halfScreenHeight)}
		g.l2.b = vector{0, halfScreenHeight + rand.Intn(halfScreenHeight)}
	}

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
