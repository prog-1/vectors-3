package main

import (
	"image/color"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	sWidth  = 600
	sHeight = 600
)

type point struct {
	x, y int
}

type Game struct {
	width, height int
	a, b, c, d    point
	img           *ebiten.Image
}

var col = color.RGBA{244, 212, 124, 255}

func DrawLine(img *ebiten.Image, x1, x2, y1, y2 int, col color.Color) {
	if math.Abs(float64(x2-x1)) >= math.Abs(float64(y2-y1)) {
		if x2 < x1 {
			x1, x2 = x2, x1
			y1, y2 = y2, y1
		}
		k := float64(y2-y1) / float64(x2-x1)
		for x, y := x1, float64(y1)+0.5; x <= x2; x, y = x+1, y+k {
			img.Set(x, int(y), col)
		}
	} else {
		if y2 < y1 {
			x1, x2 = x2, x1
			y1, y2 = y2, y1
		}
		k := float64(x2-x1) / float64(y2-y1)
		for x, y := float64(x1)+0.5, y1; y <= y2; x, y = x+k, y+1 {
			img.Set(int(x), y, col)
		}
	}
}

func NewGame(width, height int) *Game {
	return &Game{
		width:  width,
		height: height,
		a:      point{rand.Intn(sWidth), 0},
		b:      point{rand.Intn(sWidth), sHeight},
		c:      point{0, rand.Intn(sHeight)},
		d:      point{sWidth, rand.Intn(sHeight)},
		img:    ebiten.NewImage(sWidth, sHeight),
	}
}

func (g *Game) Layout(outWidth, outHeight int) (w, h int) {
	return g.width, g.height
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) NewCircle() {
	x, y := rand.Intn(sWidth), rand.Intn(sHeight)
	p1 := point{g.b.x - g.a.x, g.b.y - g.a.y}
	p2 := point{g.d.x - g.c.x, g.d.y - g.c.y}
	u1 := (x-g.a.x)*p1.y - (y-g.a.y)*p1.x
	u2 := (x-g.c.x)*p2.y - (y-g.c.y)*p2.x

	if u1 > 0 && u2 > 0 {
		g.img.Set(x, y, color.RGBA{0xDC, 14, 0x3C, 255})
	} else if u1 < 0 && u2 > 0 {
		g.img.Set(x, y, color.RGBA{0x2E, 0xff, 0x2E, 255})
	} else if u1 > 0 && u2 < 0 {
		g.img.Set(x, y, color.RGBA{0xff, 0xff, 0xff, 255})
	} else if u1 < 0 && u2 < 0 {
		g.img.Set(x, y, color.RGBA{0xff, 0xA5, 0, 255})
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.NewCircle()
	screen.DrawImage(g.img, nil)
	DrawLine(screen, int(g.a.x), int(g.b.x), int(g.a.y), int(g.b.y), col)
	DrawLine(screen, int(g.c.x), int(g.d.x), int(g.c.y), int(g.d.y), col)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ebiten.SetWindowSize(sWidth, sHeight)
	if err := ebiten.RunGame(NewGame(sWidth, sHeight)); err != nil {
		log.Fatal(err)
	}
}
