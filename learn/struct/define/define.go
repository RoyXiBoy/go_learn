package main

import "fmt"

//定义
type identifier struct {
	field1 int
	field2 float64
}

type T struct {a, b int}

func main(){
	//new
	t := new(T)
	t.a = 1
	fmt.Println(t.a)
}
