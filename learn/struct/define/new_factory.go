package main

import (
	"fmt"
	"image/color/palette"
)

//假设定义了如下的 File 结构体类型：

type File struct {
	fd      int     // 文件描述符
	name    string  // 文件名
}

//下面是这个结构体类型对应的工厂方法，它返回一个指向结构体实例的指针：

func NewFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}

	return &File{fd, name} //如果 File 是一个结构体类型，那么表达式 new(File) 和 &File{} 是等价的
}

//然后这样调用它：
func main()  {
	f := NewFile(10, "./test.txt")

	fmt.Print(f.name)
}

