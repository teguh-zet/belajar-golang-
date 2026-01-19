package main

import (
	"first-go/constants"
	"first-go/helloword"
	"first-go/looping"
	"first-go/values"
	"first-go/variables"
	"fmt"
	// "fmt"
	// "math"
	// "time"
)

func main() {
	i := 1
	if i == 1 {
		looping.GoFor()
	} else {
		fmt.Println("============================")
		constants.GoConstants()
		fmt.Println("============================")
		values.GoValues()
		fmt.Println("============================")
		helloword.GoHelloword()
		fmt.Println("============================")
		variables.GoVariables()

	}

	// fmt.Printf("hello : %g world \n", math.Sqrt(10))
	// fmt.Println("time now : \n",time.Now())
	// fmt.Println(add(42, 13))

	// //swap
	//  a,b  :="hello","world"
	// println(swap(a,b))
	// //v2
	// c,d := swap("aku", "ikan")
	// fmt.Println(c,d)

	// fmt.Println(split(11))
}
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
func add(x, y int) int {
	return x + y
}
func swap(x, y string) (string, string) {
	return y, x
}
