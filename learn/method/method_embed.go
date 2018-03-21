package main

import (
	"math"
	"fmt"
)

type Point struct {
	x, y float64
}

func (p *Point)Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

type NamePoint struct {
	Point
	name string
}

/*重写
func (n *NamePoint) Abs() float64 {
	return n.Point.Abs() * 100.
}*/

func main()  {
	n := &NamePoint{Point{3, 4}, "Pythagoras"}
	fmt.Println(n.Abs()) // 打印5
}