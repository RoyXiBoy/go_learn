package main

import (
	"fmt"
	"reflect"
)
/*
变量的最基本信息就是类型和值：反射包的 Type 用来表示一个 Go 类型，反射包的 Value 为 Go 值提供了反射接口。

两个简单的函数，reflect.TypeOf 和 reflect.ValueOf，返回被检查对象的类型和值。例如，x 被定义为：var x float64 = 3.4，那么 reflect.TypeOf(x) 返回 float64，reflect.ValueOf(x) 返回 <float64 Value>
 */



func main()  {
	var x float64 = 3.4

	fmt.Println(reflect.TypeOf(x))

	fmt.Println(reflect.ValueOf(x))

	v := reflect.ValueOf(x)

	fmt.Println(v.Kind())

	//查询是否能通过反射设置值
	fmt.Println(v.CanSet())
	//当 v := reflect.ValueOf(x) 函数通过传递一个 x 拷贝创建了 v，那么 v 的改变并不能更改原始的 x。

	//要想 v 的更改能作用到 x，那就必须传递 x 的地址 v = reflect.ValueOf(&x)
	v = reflect.ValueOf(&x)

	fmt.Println(v.CanSet())

	//通过 Type() 我们看到 v 现在的类型是 *float64 并且仍然是不可设置的 。要想让其可设置我们需要使用 Elem() 函数，这间接的使用指针：v = v.Elem()

	v = v.Elem()

	fmt.Println(v.CanSet())

	v.SetFloat(3.1415)

	fmt.Println(reflect.ValueOf(x))



}
