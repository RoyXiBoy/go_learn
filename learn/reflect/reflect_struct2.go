package main

import (
"fmt"
"reflect"
)

type T struct {
	A int
	B string
	//c float64
}
//结构中只有被导出字段（首字母大写）才是可设置的
func main(){
	//t := T{34, "hello world", 3.14}
	t := T{34, "hello world"}
	s := reflect.ValueOf(&t).Elem()

	typeOfT := s.Type()

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}

	s.Field(0).SetInt(77)
	s.Field(1).SetString("Sunset Strip")
	//s.Field(2).SetFloat(3.1415)
	fmt.Println("t is now", t)

}