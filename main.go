package main

import (
	"first-go/arrays"
	"first-go/constants"
	"first-go/decisions"
	"first-go/functions"
	"first-go/helloword"
	"first-go/looping"
	"first-go/maps"
	"first-go/slices"
	"first-go/values"
	"first-go/variables"
)

func main() {
	i := 1
	if i == 1 {
		functions.GoRecursiveFunctions()	
	} else {
		functions.GoClosureFunctions()	
		functions.GoVaradicFuntions()
		functions.GoMultipleReturn()
		functions.GoFunctions()
		maps.GoMaps()
		slices.GoSlices()
		arrays.GoArray()
		decisions.GoSwitch()
		decisions.GoIfelse()
		looping.GoFor()
		constants.GoConstants()
		values.GoValues()
		helloword.GoHelloword()
		variables.GoVariables()

	}

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
