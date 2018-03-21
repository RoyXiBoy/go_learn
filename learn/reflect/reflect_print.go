package main

import (
	"os"
	"strconv"
)

type Stringer interface {
	String() string
}

type Celsius float64

func (c Celsius) String() string {
	return strconv.FormatFloat(float64(c),'f', 1, 64) + " °C"
}

type Day int

var dayName = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

func (day Day) String() string {
	return dayName[day]
}

//具体地了解 Printf 中的反射,使用了 type-switch 来推导参数类型，并根据类型来输出每个参数的值

func print(args ...interface{}) {
	for i, arg := range args {
		if i > 0 {os.Stdout.WriteString(" ")}
		switch a := arg.(type) { // type switch
			case Stringer:	os.Stdout.WriteString(a.String())
			case int:		os.Stdout.WriteString(strconv.Itoa(a))
			case string:	os.Stdout.WriteString(a)
				// more types
			default:		os.Stdout.WriteString("???")
		}
	}
}

func main() {
	print(Day(1), "was", Celsius(18.36))  // Tuesday was 18.4 °C
	print(55.8888)
}
