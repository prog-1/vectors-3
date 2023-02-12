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
	a, b, c, d, e, f Point
	img              *ebiten.Image
	width, height    int
}

func NewGame(width, height int) *Game {
	return &Game{
		width:  width,
		height: height,
	}
}

func (g *Game) Update() error {
	return nil
}
func (g *Game) DrawPoints(a, b, c, d, e, f Point) {
	x, y := rand.Intn(screenWidth), rand.Intn(screenHeight)
	v := Point{b.x - a.x, b.y - a.y}
	u := Point{d.x - c.x, d.y - c.y}
	z := Point{f.x - e.x, f.y - e.y}
	q1 := (float64(x)-a.x)*v.y - (float64(y)-a.y)*v.x
	q2 := (float64(x)-c.x)*u.y - (float64(y)-c.y)*u.x
	q3 := (float64(x)-e.x)*z.y - (float64(y)-e.y)*z.x
	if q1 < 0 && q2 < 0 && q3 < 0 {
		g.img.Set(x, y, color.RGBA{255, 0, 0, 1})
	} else {
		g.img.Set(x, y, color.RGBA{0, 255, 3, 1})
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.DrawPoints(g.a, g.b, g.c, g.d, g.e, g.f)
	screen.DrawImage(g.img, nil)
	ebitenutil.DrawLine(g.img, float64(g.a.x), float64(g.a.y), float64(g.b.x), float64(g.b.y), color.RGBA{R: 245, G: 114, B: 227, A: 255})
	ebitenutil.DrawLine(g.img, float64(g.c.x), float64(g.c.y), float64(g.d.x), float64(g.d.y), color.RGBA{R: 195, G: 114, B: 245, A: 255})
	ebitenutil.DrawLine(g.img, float64(g.e.x), float64(g.e.y), float64(g.f.x), float64(g.f.y), color.RGBA{R: 195, G: 114, B: 245, A: 255})
}
func (g *Game) Layout(int, int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	a1 := Point{float64(rand.Intn(520)), float64(rand.Intn(360))}
	rand.Seed(time.Now().UnixNano())
	if err := ebiten.RunGame(&Game{
		a:   Point{a1.x, a1.y},
		b:   Point{a1.x + 100, a1.y},
		c:   Point{a1.x + 100, a1.y},
		d:   Point{a1.x - 80, a1.y + 100},
		e:   Point{a1.x - 80, a1.y + 100},
		f:   Point{a1.x, a1.y},
		img: ebiten.NewImage(screenWidth, screenHeight)}); err != nil {
		log.Fatal(err)
	}
}
