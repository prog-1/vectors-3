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

type Game struct {
	width, height, i1, i2 int
	x1, y1, x2, y2        float64
	img                   *ebiten.Image
}

func NewGame(width, height int) *Game {
	return &Game{
		width:  width,
		height: height,
		x1:     float64(rand.Intn(width)),
		y1:     float64(rand.Intn(height)),
		x2:     float64(rand.Intn(width)),
		y2:     float64(rand.Intn(height)),
		img:    ebiten.NewImage(width, height),
	}
}

func (g *Game) Layout(outWidth, outHeight int) (w, h int) {
	return g.width, g.height
}

func (g *Game) Update() error {
	x, y := float64(rand.Intn(g.width)), float64(rand.Intn(g.height))
	if (g.y2-g.y1)*x-float64(g.width)*y+float64(g.width)*g.y1 > 0 {
		g.i1 = 1
	} else {
		g.i1 = -1
	}
	if float64(g.height)*x+(g.x1-g.x2)*y-float64(g.height)*g.x1 > 0 {
		g.i2 = 1
	} else {
		g.i2 = -1
	}
	if g.i1 == 1 && g.i2 == 1 {
		g.img.Set(int(x), int(y), color.RGBA{0, 255, 0, 255})
	} else if g.i1 == 1 && g.i2 == -1 {
		g.img.Set(int(x), int(y), color.RGBA{255, 0, 0, 255})
	} else if g.i1 == -1 && g.i2 == 1 {
		g.img.Set(int(x), int(y), color.White)
	} else {
		g.img.Set(int(x), int(y), color.RGBA{0xFF, 0xA5, 0, 255})
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.img, nil)
	ebitenutil.DrawLine(screen, 0, g.y1, float64(g.width), g.y2, color.White)
	ebitenutil.DrawLine(screen, g.x1, 0, g.x2, float64(g.height), color.White)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	g := NewGame(screenWidth, screenHeight)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
