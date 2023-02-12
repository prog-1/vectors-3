package main

import (
	"image/color"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Line struct {
	x1, y1, x2, y2 int
	A, B, C        float64
}

func (l *Line) CalculateABC() {
	l.A = float64(l.y2 - l.y1)
	l.B = float64(l.x1 - l.x2)
	l.C = float64((l.x2-l.x1)*l.y1 - (l.y2-l.y1)*l.x1)

}

func (l *Line) OffSet(x, y float64) float64 {
	return l.A*x + l.B*y + l.C
}

type Game struct {
	l1, l2     Line
	BackBuffer *ebiten.Image
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

func (g *Game) Update() error {
	x, y := float64(rand.Intn(640)), float64(rand.Intn(480))

	// fmt.Println()
	if g.l1.OffSet(x, y) < 0 {
		if g.l2.OffSet(x, y) < 0 {
			ebitenutil.DrawCircle(g.BackBuffer, float64(x), float64(y), 3, color.RGBA{255, 0, 0, 255})
		} else {
			ebitenutil.DrawCircle(g.BackBuffer, float64(x), float64(y), 3, color.RGBA{255, 255, 255, 255})

		}
	} else {
		if g.l2.OffSet(x, y) < 0 {
			ebitenutil.DrawCircle(g.BackBuffer, float64(x), float64(y), 3, color.RGBA{0, 255, 0, 255})
		} else {
			ebitenutil.DrawCircle(g.BackBuffer, float64(x), float64(y), 3, color.RGBA{0xFF, 0xA5, 0, 255})

		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.BackBuffer, &ebiten.DrawImageOptions{})
	DrawLineDDA(screen, float64(g.l1.x1), float64(g.l1.y1), float64(g.l1.x2), float64(g.l1.y2), color.White)
	DrawLineDDA(screen, float64(g.l2.x1), float64(g.l2.y1), float64(g.l2.x2), float64(g.l2.y2), color.White)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	ebiten.SetWindowSize(640, 480)
	// ebiten.SetWindowTitle("Hello, World!")
	g := &Game{Line{rand.Intn(640), 0, rand.Intn(640), 480, 0, 0, 0}, Line{0, rand.Intn(480), 640, rand.Intn(480), 0, 0, 0}, ebiten.NewImage(640, 480)}
	g.l1.CalculateABC()
	g.l2.CalculateABC()
	rand.Seed(time.Now().UnixNano())
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
