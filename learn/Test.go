package main

import (
	"fmt"
	"math"
)

type Books struct {
	 name string
	 number int
}

type Phone interface {
	call()
}

type NokiaPhone struct {

}

func (nokia NokiaPhone) call(){
	fmt.Println("i 'm nokia")
}


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


	//指针使用
	var ip *int        /* 声明指针变量 */

	ip = &a  /* 指针变量的存储地址 */

	fmt.Printf("a 变量的地址是: %x\n", &a  )

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip )

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip )

	//nil 为java的null 空指针

	//指针的指针
	var ptr *int
	var pptr **int

	a = 3000

	/* 指针 ptr 地址 */
	ptr = &a

	/* 指向指针 ptr 地址 */
	pptr = &ptr
	/* 获取 pptr 的值 */
	fmt.Printf("变量 a = %d\n", a )
	fmt.Printf("指针变量 *ptr = %d\n", *ptr )
	fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)

	//结构体
	var book Books

	book.name = "肖隆韬的书"
	book.number = 100

	//Go 语言切片(Slice) 其实就是动态数组 因为go数组不可变
	//初始化有两种方式 ：=  或者和 数组一样 var []T
	//
	 nums:= []int{0,1,2,3,4,5,6,7,8}
	printSlice(nums)

	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers[1:4] ==", nums[1:4])

	/* 默认下限为 0*/
	fmt.Println("numbers[:3] ==", nums[:3])

	/* 默认上限为 len(s)*/
	fmt.Println("numbers[4:] ==", nums[4:])

	numbers1 := make([]int,0,5)
	printSlice(numbers1)

	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	number2 := nums[:2]
	printSlice(number2)

	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	number3 := nums[2:5]
	printSlice(number3)

	//range 使用
	sum := 0
	for _, num := range nums{
		sum += num
	}
	fmt.Println(sum)

	for i, num := range nums {
		if num == 2 {
			fmt.Println(i)
		}
	}
	for i, c := range "go" {
		fmt.Println(i, c)
	}

	//map 如果不初始化 map，那么就会创建一个 nil map。nil map 不能用来存放键值对
	var currentMap map[string]string = make(map[string]string)
	currentMap["1"] = "1"
	currentMap["2"] = "2"

	fmt.Println(currentMap["1"])

	//强转
	var sum2 int = 17
	var count int = 5
	var mean float32

	mean = float32(sum2)/float32(count)
	fmt.Printf("mean 的值为: %f\n",mean)

	//接口
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()
}

func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
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



