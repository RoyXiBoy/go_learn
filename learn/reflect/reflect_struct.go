package main

//有些时候需要反射一个结构类型。NumField() 方法返回结构内的字段数量；通过一个 for 循环用索引取得每个字段的值 Field(i)。
//我们同样能够调用签名在结构上的方法，例如，使用索引 n 来调用：Method(n).Call(nil)。

import ("fmt"
		"reflect"
		)

type NotknownType struct {
	s1, s2, s3 string
	s4, s5, s6 int
}

func (n NotknownType) String() string {
	return n.s1 + " - " + n.s2 + " - " + n.s3
}



func main()  {
	// variable to investigate:
	var secret interface{} = NotknownType{"Ada", "Go", "Oberon", 4, 5, 6}

	value := reflect.ValueOf(secret)

	tye := reflect.TypeOf(secret)

	knd := tye.Kind()

	fmt.Println(value)

	fmt.Println(knd)

	// iterate through the fields of the struct:
	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("Field %d: %v\n", i, value.Field(i))
		// error: panic: reflect.Value.SetString using value obtained using unexported field
		//value.Field(i).SetString("C#")
	}

	// call the first method, which is String():
	results := value.Method(0).Call(nil)
	fmt.Println(results) // [Ada - Go - Oberon]

}