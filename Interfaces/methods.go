package main

// Polymorphism in go
import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

type Path []Point

type Line struct {
	begin, end Point
}

func (l Line) Distance() float64 {
	return math.Hypot(l.end.x-l.begin.x, l.end.y-l.begin.y)
}

func (p Path) Distance() (sum float64) {
	for i := 1; i < len(p); i++ {
		sum += Line{p[i-1], p[i]}.Distance()
	}

	return sum
}

type Distancer interface {
	Distance() float64
}

func PrintDistance(d Distancer) {
	fmt.Println(d.Distance())
}

func main() {
	side := Line{Point{1, 2}, Point{4, 6}}
	perimeter := Path{{1, 1}, {5, 1}, {5, 4}, {1, 1}}

	PrintDistance(side)
	PrintDistance(perimeter)
}
