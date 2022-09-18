package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

type Line struct {
	begin, end Point
}

func (l Line) Distance() float64 {
	return math.Hypot(l.end.x-l.begin.x, l.end.y-l.begin.y)
}

func main() {
	fmt.Println(Line{Point{1, 2}, Point{4, 6}}.Distance())
}
