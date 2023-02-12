package main

import (
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//---------------------------Declaration--------------------------------

const (
	sW = 640 //screen width
	sH = 480 //screen height
)

type Game struct {
	width, height    int
	a, b, c, d, e, f point
	pixels           []*pixel
	red, green       color.RGBA
}

type point struct {
	x, y float64
}

type pixel struct {
	x, y  float64
	color color.RGBA
}

//---------------------------Update-------------------------------------

func (g *Game) Update() error {
	g.pixels = append(g.pixels, g.randomPixel(g.width, g.height)) //adding new pixel
	return nil
}

//---------------------------Draw-------------------------------------

func (g *Game) Draw(screen *ebiten.Image) {

	for _, p := range g.pixels {
		screen.Set(int(p.x), int(p.y), p.color)
	}

	//drawing hexagon!
	ebitenutil.DrawLine(screen, g.a.x, g.a.y, g.b.x, g.b.y, color.RGBA{255, 255, 255, 255})
	ebitenutil.DrawLine(screen, g.b.x, g.b.y, g.c.x, g.c.y, color.RGBA{255, 255, 255, 255})
	ebitenutil.DrawLine(screen, g.c.x, g.c.y, g.d.x, g.d.y, color.RGBA{255, 255, 255, 255})
	ebitenutil.DrawLine(screen, g.d.x, g.d.y, g.e.x, g.e.y, color.RGBA{255, 255, 255, 255})
	ebitenutil.DrawLine(screen, g.e.x, g.e.y, g.f.x, g.f.y, color.RGBA{255, 255, 255, 255})
	ebitenutil.DrawLine(screen, g.f.x, g.f.y, g.a.x, g.a.y, color.RGBA{255, 255, 255, 255})
}

//-------------------------Functions----------------------------------

func (g *Game) randomPixel(width, height int) *pixel {

	var color color.RGBA

	//random point coordinates
	x := float64(rand.Intn(width))
	y := float64(rand.Intn(height))

	//line normals
	nrm1 := formula(x, y, g.a.x, g.a.y, g.b.x, g.b.y)
	nrm2 := formula(x, y, g.b.x, g.b.y, g.c.x, g.c.y)
	nrm3 := formula(x, y, g.c.x, g.c.y, g.d.x, g.d.y)
	nrm4 := formula(x, y, g.d.x, g.d.y, g.e.x, g.e.y)
	nrm5 := formula(x, y, g.e.x, g.e.y, g.f.x, g.f.y)
	nrm6 := formula(x, y, g.f.x, g.f.y, g.a.x, g.a.y)

	//set the pixel color depending on normals
	if nrm1 > 0 && nrm2 > 0 && nrm3 > 0 && nrm4 > 0 && nrm5 > 0 && nrm6 > 0 {
		color = g.green
	} else {
		color = g.red
	}

	return &pixel{x, y, color} //return pixel with random coordinates and according color
}

//returns result of the normal formula of the line
func formula(x, y, x1, y1, x2, y2 float64) float64 {
	dy := y2 - y1
	dx := x2 - x1
	A1 := -dy
	B1 := dx
	C1 := dy*x1 - dx*y1
	return A1*x + B1*y + C1
}

//---------------------------Main-------------------------------------

func (g *Game) Layout(inWidth, inHeight int) (outWidth, outHeight int) {
	return g.width, g.height
}

func main() {

	rand.Seed(time.Now().UnixNano()) //seed for random

	ebiten.SetWindowSize(sW, sH)
	ebiten.SetWindowTitle("Hexagon!")
	ebiten.SetWindowResizable(true) //enablening window resizes

	//running game
	g := NewGame(sW, sH)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func NewGame(width, height int) *Game {

	//All pre-declared stuff is stored here
	var a, b, c, d, e, f point
	a.x, a.y = (sW / 2), (sH/2)+100
	b.x, b.y = (sW/2)-100, (sH/2)+50
	c.x, c.y = (sW/2)-100, (sH/2)-50
	d.x, d.y = (sW / 2), (sH/2)-100
	e.x, e.y = (sW/2)+100, (sH/2)-50
	f.x, f.y = (sW/2)+100, (sH/2)+50

	red := color.RGBA{255, 100, 100, 255}
	green := color.RGBA{100, 255, 100, 255}

	//creating and returning game instance
	return &Game{sW, sH, a, b, c, d, e, f, nil, red, green}

}
