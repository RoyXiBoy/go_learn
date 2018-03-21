/*
同步包的使用
 */

package main

import "sync"

type Book struct {
	mu sync.Mutex
	name string
}

func Update(book *Book) {
	book.mu.Lock()

	book.name = "xlt"

	book.mu.Unlock()
}

//相对简单的情况下，通过使用 sync 包可以解决同一时间只能一个线程访问变量或 map 类型数据的问题。
// 如果这种方式导致程序明显变慢或者引起其他问题，我们要重新思考来通过 goroutines 和 channels 来解决问题，这是在 Go 语言中所提倡用来实现并发的技术。

func main()  {


}
