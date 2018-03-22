package main

import "fmt"

//使用两次 make() 函数，第一次分配切片，第二次分配 切片中每个 map 元素
func main(){
	items := make([]map[int]int, 5)//map 类型的切片
	for i := range items{
		items[i] = make(map[int]int, 1)
		items[i][1] = 1
	}
	fmt.Printf("Version A: Value of items: %v\n", items)
}

