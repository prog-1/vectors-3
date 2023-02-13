package main

import (
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

type Point struct {
	x, y float64
}

type Game struct {
	width, height, i1, i2, i3, i4 int
	a, b, c, d, e, f, g, h        Point
	img                           *ebiten.Image
}

func NewGame(width, height int) *Game {
	return &Game{
		width:  width,
		height: height,
		a:      Point{100, 100},
		b:      Point{300, 150},
		c:      Point{300, 150},
		d:      Point{400, 300},
		e:      Point{400, 300},
		f:      Point{150, 200},
		g:      Point{150, 200},
		h:      Point{100, 100},
		img:    ebiten.NewImage(width, height),
	}
}

func (g *Game) Layout(outWidth, outHeight int) (w, h int) {
	return g.width, g.height
}

func (g *Game) Update() error {
	x, y := float64(rand.Intn(g.width)), float64(rand.Intn(g.height))
	if (g.b.y-g.a.y)*x+(g.a.x-g.b.x)*y+(g.b.x*g.a.y-g.b.y*g.a.x) > 0 {
		g.i1 = 1
	} else {
		g.i1 = -1
	}
	if (g.d.y-g.c.y)*x+(g.c.x-g.d.x)*y+(g.d.x*g.c.y-g.d.y*g.c.x) > 0 {
		g.i2 = 1
	} else {
		g.i2 = -1
	}
	if (g.f.y-g.e.y)*x+(g.e.x-g.f.x)*y+(g.f.x*g.e.y-g.f.y*g.e.x) > 0 {
		g.i3 = 1
	} else {
		g.i3 = -1
	}
	if (g.h.y-g.g.y)*x+(g.g.x-g.h.x)*y+(g.h.x*g.g.y-g.h.y*g.g.x) > 0 {
		g.i4 = 1
	} else {
		g.i4 = -1
	}
	if g.i1 == -1 && g.i2 == -1 && g.i3 == -1 && g.i4 == -1 {
		g.img.Set(int(x), int(y), color.White)
	} else {
		g.img.Set(int(x), int(y), color.RGBA{255, 0, 0, 255})
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.img, nil)
	ebitenutil.DrawLine(screen, g.a.x, g.a.y, g.b.x, g.b.y, color.White)
	ebitenutil.DrawLine(screen, g.c.x, g.c.y, g.d.x, g.d.y, color.White)
	ebitenutil.DrawLine(screen, g.e.x, g.e.y, g.f.x, g.f.y, color.White)
	ebitenutil.DrawLine(screen, g.g.x, g.g.y, g.h.x, g.h.y, color.White)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	g := NewGame(screenWidth, screenHeight)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
