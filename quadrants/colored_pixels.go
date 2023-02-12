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
	width, height             int
	a, b, c, d                point
	pixels                    []*pixel
	red, green, white, orange color.RGBA
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

	//Random Lines - draw lines after pixels to draw them above all pixels
	ebitenutil.DrawLine(screen, g.a.x, g.a.y, g.b.x, g.b.y, color.RGBA{255, 255, 255, 255})
	ebitenutil.DrawLine(screen, g.c.x, g.c.y, g.d.x, g.d.y, color.RGBA{255, 255, 255, 255})
}

//-------------------------Functions----------------------------------

func (g *Game) randomPixel(width, height int) *pixel {

	var color color.RGBA

	//random point coordinates
	x := float64(rand.Intn(width))
	y := float64(rand.Intn(height))

	s1 := mathStuff(x, y, g.a.x, g.a.y, g.b.x, g.b.y) //first formula result
	s2 := mathStuff(x, y, g.c.x, g.c.y, g.d.x, g.d.y) //second formula result

	//set the color depending on both formula result signs
	if s1 > 0 && s2 >= 0 {
		color = g.red
	} else if s1 > 0 && s2 <= 0 {
		color = g.green
	} else if s1 < 0 && s2 >= 0 {
		color = g.white
	} else /* s1 < 0 && s2 <= 0 */ {
		color = g.orange
	}

	return &pixel{x, y, color} //return pixel with random coordinates and according color
}

//returns result of the normal formula of the line
func mathStuff(x, y, x1, y1, x2, y2 float64) float64 {
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
	ebiten.SetWindowTitle("Colored Pixels!")
	ebiten.SetWindowResizable(true) //enablening window resizes

	//running game
	g := NewGame(sW, sH)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func NewGame(width, height int) *Game {

	//All pre-declared stuff is stored here
	var a, b, c, d point
	a.x, a.y = float64(rand.Intn(sW)), 0
	b.x, b.y = float64(rand.Intn(sW)), sH
	c.x, c.y = 0, float64(rand.Intn(sH))
	d.x, d.y = sW, float64(rand.Intn(sH))
	red := color.RGBA{255, 100, 100, 255}
	green := color.RGBA{100, 255, 100, 255}
	white := color.RGBA{255, 255, 255, 255}
	orange := color.RGBA{200, 165, 0, 255}

	//creating and returning game instance
	return &Game{sW, sH, a, b, c, d, nil, red, green, white, orange}

}
