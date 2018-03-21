package main

import "fmt"

type Shaper interface {
	Area() float32
}

type TopologicalGenus interface {
	Rank() int
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (sq *Square) Rank() int {
	return 1
}

type Rectangle struct {
	length, width float32
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

func (r Rectangle) Rank() int {
	return 2
}

func main(){
	r := Rectangle{5,3}

	q := &Square{5}
	shapes := []Shaper{r, q}
	fmt.Println("Looping through shapes for area ...")

	for n, _ := range shapes{
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shape is: ", shapes[n].Area())
	}

	topgen := []TopologicalGenus{r, q}
	fmt.Println("Looping through topgen for rank ...")
	for n, _ := range topgen {
		fmt.Println("Shape details: ", topgen[n])
		fmt.Println("Topological Genus of this shape is: ", topgen[n].Rank())
	}

	// 接口 -  结构体 - 方法实现  其中方法实现是把接口和结构体关联起来的
	// 接口：方法声明  结构体： 成员    方法实现： 实现

	// 使用： 根据结构体构造 实例。  继承同一个接口的，可以使用接口数组来 调用方法。
}