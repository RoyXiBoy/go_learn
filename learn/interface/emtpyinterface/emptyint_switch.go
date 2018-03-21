package main

import "fmt"

type specialString string

var whatIsThis specialString = "hello"

func TypeSwitch(any2 interface{}) {
	testFunc := func(any interface{}) {
		switch v := any.(type) {
		case bool:
			fmt.Printf("any %v is a bool type\n", v)
		case int:
			fmt.Printf("any %v is an int type\n", v)
		case float32:
			fmt.Printf("any %v is a float32 type\n", v)
		case string:
			fmt.Printf("any %v is a string type\n", v)
		case specialString:
			fmt.Printf("any %v is a special String!\n", v)
		default:
			fmt.Println("unknown type!")
		}
	}
	testFunc(any2)
}

func main() {
	TypeSwitch(1)
	TypeSwitch("2")
	TypeSwitch(whatIsThis)
}
