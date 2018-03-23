package main

import "fmt"

/*
main建议写在最前。 函数 方法 按逻辑往下写。
 */



func f() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}
func main() {
	fmt.Println(f())
}