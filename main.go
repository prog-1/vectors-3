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
	winWidth  = 800
	winHeight = 600
	winTitle  = "lines"
)

type point struct {
	x, y int
}

type Game struct {
	width, height int
	a, b, c, d    point
	img           *ebiten.Image
}

func DrawLineDDA(img *ebiten.Image, x1, y1, x2, y2 float64, c color.Color) {
	if math.Abs(x2-x1) <= math.Abs(y2-y1) {
		if y2 < y1 {
			x1, x2 = x2, x1
			y1, y2 = y2, y1
		}
		k := float64(x2-x1) / float64(y2-y1)
		for x, y := float64(x1)+0.5, y1; y <= y2; x, y = x+k, y+1 {
			img.Set(int(x), int(y), c)
		}
	} else {
		if x2 < x1 {
			x1, x2 = x2, x1
			y1, y2 = y2, y1
		}
		k := float64(y2-y1) / float64(x2-x1)
		for x, y := x1, float64(y1)+0.5; x <= x2; x, y = x+1, y+k {
			img.Set(int(x), int(y), c)
		}
	}

}

var c = color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}

func NewGame(width, height int) *Game {
	return &Game{
		width:  width,
		height: height,
		a:      point{rand.Intn(winWidth), 0},
		b:      point{rand.Intn(winWidth), winHeight},
		c:      point{0, rand.Intn(winHeight)},
		d:      point{winWidth, rand.Intn(winHeight)},
		img:    ebiten.NewImage(winWidth, winHeight),
	}
}

func (g *Game) Layout(outWidth, outHeight int) (w, h int) {
	return g.width, g.height
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) pointGen() {
	x, y := rand.Intn(winWidth), rand.Intn(winHeight)
	v1 := point{g.b.x - g.a.x, g.b.y - g.a.y}
	v2 := point{g.d.x - g.c.x, g.d.y - g.c.y}
	q1 := (x-g.a.x)*v1.y - (y-g.a.y)*v1.x
	q2 := (x-g.c.x)*v2.y - (y-g.c.y)*v2.x

	switch {
	case q1 > 0 && q2 > 0:
		g.img.Set(x, y, color.RGBA{0xff, 0, 0, 255})
	case q1 < 0 && q2 > 0:
		g.img.Set(x, y, color.RGBA{0, 0xff, 0, 255})
	case q1 > 0 && q2 < 0:
		g.img.Set(x, y, color.RGBA{0xff, 0xff, 0xff, 255})
	case q1 < 0 && q2 < 0:
		g.img.Set(x, y, color.RGBA{0xff, 0xA5, 0, 255})
	}

}
func (g *Game) Draw(screen *ebiten.Image) {
	g.pointGen()
	screen.DrawImage(g.img, nil)
	DrawLineDDA(screen, float64(g.a.x), float64(g.a.y), float64(g.b.x), float64(g.b.y), c)
	DrawLineDDA(screen, float64(g.c.x), float64(g.c.y), float64(g.d.x), float64(g.d.y), c)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ebiten.SetWindowTitle(winTitle)
	ebiten.SetWindowSize(winWidth, winHeight)
	ebiten.SetWindowResizable(true)
	if err := ebiten.RunGame(NewGame(winWidth, winHeight)); err != nil {
		log.Fatal(err)
	}
}
