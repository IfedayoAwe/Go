package main

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

func main() {
	fmt.Println(Line{Point{1, 2}, Point{4, 6}}.Distance())
	fmt.Println(Path{{1, 1}, {5, 1}, {5, 4}, {1, 1}}.Distance())
}
