package main
/*
大精度数包的使用
 */

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	im := big.NewInt(3)
	in := im
	io := big.NewInt(2)
	ip := big.NewInt(2)
	ip.Mul(im, in).Add(ip, im).Div(ip, io)
	//im * in =ip   ip = ip + im  ip = ip / io
	fmt.Printf("Big Int: %v\n", ip)


	rm := big.NewRat(math.MaxInt64, 1956)
	rn := big.NewRat(-1956, math.MaxInt64)
	ro := big.NewRat(19, 56)
	rp := big.NewRat(1111, 2222)
	rq := big.NewRat(1, 1)
	rq.Mul(rm, rn).Add(rq, ro).Mul(rq, rp)
	fmt.Printf("Big Rat: %v\n", rq)

}