package main

// Polymorphism in go
// import (
// 	"fmt"
// 	"math"
// )

// type Point struct {
// 	x, y float64
// }

// type Path []Point

// type Line struct {
// 	begin, end Point
// }

// func (l Line) Distance() float64 {
// 	return math.Hypot(l.end.x-l.begin.x, l.end.y-l.begin.y)
// }

// func (p Path) Distance() (sum float64) {
// 	for i := 1; i < len(p); i++ {
// 		sum += Line{p[i-1], p[i]}.Distance()
// 	}

// 	return sum
// }

// type Distancer interface {
// 	Distance() float64
// }

// func PrintDistance(d Distancer) {
// 	fmt.Println(d.Distance())
// }

// func main() {
// 	side := Line{Point{1, 2}, Point{4, 6}}
// 	perimeter := Path{{1, 1}, {5, 1}, {5, 4}, {1, 1}}

// 	PrintDistance(side)
// 	PrintDistance(perimeter)
// }

// OOP in go
// import (
// 	"fmt"
// 	"math"
// )

// type Point struct {
// 	x, y float64
// }

// type Line struct {
// 	begin, end Point
// }

// func (l Line) Distance() float64 {
// 	return math.Hypot(l.end.x-l.begin.x, l.end.y-l.begin.y)
// }

// func (l Line) ScaleBy(f float64) Line {
// 	l.end.x += (f - 1) * (l.end.x - l.begin.x)
// 	l.end.y += (f - 1) * (l.end.y - l.begin.y)

// 	return Line{l.begin, Point{l.end.x, l.end.y}}
// }

// func main() {
// 	side := Line{Point{1, 2}, Point{4, 6}}

// 	s2 := side.ScaleBy(2)
// 	fmt.Println(s2.Distance())
// 	fmt.Println(Line{Point{1, 2}, Point{4, 6}}.ScaleBy(2).Distance())
// }

// import (
// 	"fmt"
// 	"image/color"
// 	"math"
// )

// type Point struct {
// 	X, Y float64
// }

// func (p Point) Distance(q Point) float64 {
// 	return math.Hypot(q.X-p.X, q.Y-p.Y)
// }

// type ColoredPoint struct {
// 	Point
// 	Color color.RGBA
// }

// func main() {
// 	p, q := Point{1, 1}, ColoredPoint{Point{5, 4}, color.RGBA{255, 0, 0, 255}}

// 	l1 := q.Distance(p)
// 	l2 := p.Distance(q.Point) // OK: but p.Distance(q) is NOT ALLOWED

// 	fmt.Println(l1, l2)
// }
