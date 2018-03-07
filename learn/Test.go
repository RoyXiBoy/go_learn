package main

import (
	"fmt"
	"math"
)

func main()  {
	fmt.Print("hello world")

	var test1 int = 1
	var test2 = "2"
	test3 := 3.34

	fmt.Println(test1)
	fmt.Println(test2)
	fmt.Println(test3)

	const test4 int = 4
	const test5 = "5"

	a := 21
	b := 55
	var c int

	c = a + b
	fmt.Printf("第一行 - c 的值为 %d\n", c )
	c = a - b
	fmt.Printf("第二行 - c 的值为 %d\n", c )
	c = a * b
	fmt.Printf("第三行 - c 的值为 %d\n", c )
	c = a / b
	fmt.Printf("第四行 - c 的值为 %d\n", c )
	c = a % b
	fmt.Printf("第五行 - c 的值为 %d\n", c )
	a++
	fmt.Printf("第六行 - a 的值为 %d\n", a )
	a=21   // 为了方便测试，a 这里重新赋值为 21
	a--
	fmt.Printf("第七行 - a 的值为 %d\n", a )


	if a == b  {
		fmt.Printf("第一行 - a 等于 b\n" )
	} else {
		fmt.Printf("第一行 - a 不等于 b\n" )
	}
	if ( a < b ) {
		fmt.Printf("第二行 - a 小于 b\n" )
	} else {
		fmt.Printf("第二行 - a 不小于 b\n" )
	}

	if ( a > b ) {
		fmt.Printf("第三行 - a 大于 b\n" )
	} else {
		fmt.Printf("第三行 - a 不大于 b\n" )
	}
	/* Lets change value of a and b */
	a = 5
	b = 20
	if ( a <= b ) {
		fmt.Printf("第四行 - a 小于等于 b\n" )
	}
	if ( b >= a ) {
		fmt.Printf("第五行 - b 大于等于 a\n" )
	}

	var d *int= &a

	fmt.Println(d)
	fmt.Println(*d)

	if a > 6{
		fmt.Println("a 比 6 大")
	}else {
		fmt.Println("a 比 6 小")
	}

	for i := 0; i < 5;i++  {
		fmt.Printf("%d ",i)
	}

	for a < b{
		a++
	}

	numbers := [6]int{1, 2, 3, 5}

	for i,x := range numbers{
		fmt.Printf("第%d位的数为%d\n", i, x)
	}

	fmt.Printf("a ,b 中较大的是 %d\n", max(a, b))

	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	fmt.Println(getSquareRoot(9))

	var balance = []int{1, 2, 3, 4, 5, 6}

	fmt.Printf("%d\n", balance[5])

	
}

func max(num1 ,num2 int)  int{
	var num int
	if num1 > num2 {
		num = num1
	}else {
		num = num2
	}
	return num
}

func swap(num1 , num2 int) (int, int){
	var num3, num4 int
	num3 = num2
	num4 = num1
	return num3,num4
}



