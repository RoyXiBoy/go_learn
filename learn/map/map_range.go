package main

import "fmt"

/*
for key, value := range map1 {

}
 */

func main()  {
	map1 := make(map[int]float32)
	map1[1] = 1.0
	map1[2] = 2.0
	map1[3] = 3.0
	map1[4] = 4.0

	for key, value := range map1{
		fmt.Printf("key: %d , value: %f\n", key, value)
	}
//map 不是按照 key 的顺序排列的，也不是按照 value 的序排列的
}