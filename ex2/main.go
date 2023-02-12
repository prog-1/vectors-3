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
	polygon       []point
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
		width:   width,
		height:  height,
		polygon: []point{{100, 400}, {0, 200}, {100, 100}, {400, 100}, {500, 100}, {700, 300}, {400, 400}},
		img:     ebiten.NewImage(winWidth, winHeight),
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
	b := true
	for i := 1; i < len(g.polygon); i++ {
		v := point{g.polygon[i-1].x - g.polygon[i].x, g.polygon[i-1].y - g.polygon[i].y}
		q := (x-g.polygon[i].x)*v.y - (y-g.polygon[i].y)*v.x
		if q < 0 {
			b = false
			break
		}
	}
	v := point{g.polygon[len(g.polygon)-1].x - g.polygon[0].x, g.polygon[len(g.polygon)-1].y - g.polygon[0].y}
	q := (x-g.polygon[0].x)*v.y - (y-g.polygon[0].y)*v.x
	if q < 0 {
		b = false
	}
	if b {
		g.img.Set(x, y, color.RGBA{0xff, 0, 0, 255})
	} else {
		g.img.Set(x, y, color.RGBA{0xff, 0xff, 0xff, 255})
	}

}
func (g *Game) Draw(screen *ebiten.Image) {
	g.pointGen()
	screen.DrawImage(g.img, nil)
	for i := 1; i < len(g.polygon); i++ {
		DrawLineDDA(screen, float64(g.polygon[i].x), float64(g.polygon[i].y), float64(g.polygon[i-1].x), float64(g.polygon[i-1].y), c)
	}
	DrawLineDDA(screen, float64(g.polygon[len(g.polygon)-1].x), float64(g.polygon[len(g.polygon)-1].y), float64(g.polygon[0].x), float64(g.polygon[0].y), c)
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
