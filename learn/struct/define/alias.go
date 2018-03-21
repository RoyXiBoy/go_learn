package main
import "fmt"

type number struct { //当为结构体定义了一个 alias 类型时，此结构体类型和它的 alias 类型都有相同的底层类型
	f float32
}

type nr number   // alias type (别名)

func main() {
	a := number{5.0}
	b := nr{5.0}
	// var i float32 = b   // compile-error: cannot use b (type nr) as type float32 in assignment
	// var i = float32(b)  // compile-error: cannot convert b (type nr) to type float32
	// var c number = b    // compile-error: cannot use b (type nr) as type number in assignment
	// needs a conversion:
	var c = number(b)
	fmt.Println(a, b, c)
}