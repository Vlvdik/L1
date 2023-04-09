package ex24

import (
	"fmt"
	"math"
)

type Point struct {
	X float64
	Y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{
		X: x,
		Y: y,
	}
}

func (p *Point) Distance(p2 *Point) float64 {
	dX := p.X - p2.X
	dY := p.Y - p2.Y
	hypotenuse := math.Pow(math.Pow(dY, 2)+math.Pow(dX, 2), 0.5)

	return hypotenuse
}

func Test() {
	p1 := NewPoint(1, -1)
	p2 := NewPoint(1, 1)
	p3 := NewPoint(-2, 3)

	fmt.Printf("[main] The difference between p1 and p2: %f\n", p1.Distance(p2))
	fmt.Printf("[main] The difference between p1 and p3: %f\n", p1.Distance(p3))

	fmt.Println("[main] done")
}
