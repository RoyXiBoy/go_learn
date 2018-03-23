package main

import "fmt"

func main(){
	fmt.Println("1")
	defer test_defer()
	fmt.Println("3")
	a()
	f()

	trace("xlt")
	untrace("xlt")

	}

//延迟执行 类似 finally
func test_defer(){
	fmt.Println("2")
}


//虽然是延迟执行，但是当时传入什么参数，就是什么参数
func a() {
	fmt.Println("a------------")
	i := 0
	defer fmt.Println(i)
	i++
	return
}

//逆序执行
func f() {
	fmt.Println("f----------------")
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d \n", i)
	}
}


//代码追踪
func trace(s string) { fmt.Println("entering:", s) }
func untrace(s string) { fmt.Println("leaving:", s) }
