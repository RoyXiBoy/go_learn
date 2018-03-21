package main

import (
	"fmt"
	"math"
)

type B struct {
	thing int
}


func (b *B) change() { b.thing = 1 }

func (b B) write() string { return fmt.Sprint(b) }

/*
change()接受一个指向 B 的指针，并改变它内部的成员；
write() 接受通过拷贝接受 B 的值并只输出B的内容。

注意 Go 为我们做了探测工作，我们自己并没有指出是否在指针上调用方法，Go 替我们做了这些事情。
b1 是值而 b2 是指针，方法都支持运行了。
 */
func main() {
	var b1 B // b1是值
	b1.change()
	fmt.Println(b1.write())

	b2 := new(B) // b2是指针
	b2.change()
	fmt.Println(b2.write())
}


type Point3 struct { x, y, z float64 }
// A method on Point3

func (p Point3) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y + p.z*p.z)
}

///用指针而不用值传递是因为 值传递传递的是指针的拷贝，这样的花销比较大。

func main2()  {
	p1 := Point3{1,2,3}
	p2 := &Point3{4,5,6}
	p1.Abs() //值传递
	(*p2).Abs() //指针传递
}

