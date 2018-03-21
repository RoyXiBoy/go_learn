package main

import "fmt"

//Go 方法是作用在接收者（receiver）上的一个函数，接收者是某种类型的变量。因此方法是一种特殊类型的函数。
//一个类型加上它的方法等价于面向对象中的一个类。
//接收者类型可以是（几乎）任何类型, 不能是一个接口类型，因为接口是一个抽象定义,
//不能是一个指针类型，但是它可以是任何其他允许类型的指针。 (暂时不理解)

/*
定义方法的一般格式如下：

func (recv receiver_type) methodName(parameter_list) (return_value_list) { ... }

如果 recv 是 receiver 的实例，Method1 是它的方法名，那么方法调用遵循传统的 object.name 选择器符号：recv.Method1()。

如果 recv 一个指针，Go 会自动解引用。

如果方法不需要使用 recv 的值，可以用 _ 替换它，比如：

func (_ receiver_type) methodName(parameter_list) (return_value_list) { ... }
 */

type TwoInts struct {
	a int
	b int
}

func main()  {
	 t := TwoInts{1,2}
	fmt.Print(t.AddT())

}

func (t TwoInts) AddT() int{
	return t.a + t.b

}

//方法和函数区别：

/**
函数将变量作为参数：Function1(recv)

方法在变量上被调用：recv.Method1()

在接收者是指针时，方法可以改变接收者的值（或状态），这点函数也可以做到（当参数作为指针传递，即通过引用调用时，函数也可以改变参数的状态）。

不要忘记 Method1 后边的括号 ()，否则会引发编译器错误：method recv.Method1 is not an expression, must be called



接收者必须有一个显式的名字，这个名字必须在方法中被使用。

receiver_type 叫做 （接收者）基本类型，这个类型必须在和方法同样的包中被声明。

在 Go 中，（接收者）类型关联的方法不写在类型结构里面，就像类那样；耦合更加宽松；类型和方法之间的关联由接收者来建立。

方法没有和数据定义（结构体）混在一起：它们是正交的类型；表示（数据）和行为（方法）是独立的。

 */