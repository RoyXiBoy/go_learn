package main

import "fmt"

type IDuck interface {
	Quack()
	Walk()
}

type Bird struct {
	// ...
}

//感觉go实现接口的方法就是结构体要实现接口定义的所有方法才叫实现。 实现后 因为形参是接口（父类），因此结构体（子类）也可以放进去。
func (b *Bird) Quack() {
	fmt.Println("I am quacking!")
}

func (b *Bird) Walk()  {
	fmt.Println("I am walking!")
}

func DuckDance(duck IDuck) {
	for i := 1; i <= 3; i++ {
		duck.Quack()
		duck.Walk()
	}
}

func main() {
	b := new(Bird)
	DuckDance(b)
}