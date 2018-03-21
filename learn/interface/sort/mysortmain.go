package main

import ("fmt"
		"./mysort"
)

func ints1() {
	data := []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	a := mysort.IntArray(data) //conversion to type IntArray
	mysort.Sort(a)
	if !mysort.IsSorted(a) {
		panic("fails")
	}
	fmt.Printf("The sorted array is: %v\n", a)
}
/*
func float1(){
	data := []float64{74.12 , 59.12, 238.12, -784.12, 9845.23, 959.23, 905.34, 0, 0, 42.55, 7586.66, -5467984.78, 7586.99}
	a := mysort.Float64Array(data)
	mysort.Sort(a)
	if !mysort.IsSorted(a){
		panic("fails")
	}
	fmt.Printf("The sorted array is: %v\n", a)
}*/

func main()  {
	ints1()
	//float1()
}