package main

import (
	fundamentals "first-go/01-fundamentals"
	collections "first-go/02-collections"
	functions "first-go/03-functions"
	types "first-go/04-types"
	errorhandling "first-go/05-error-handling"
)

func main() {
	i := 1
	if i == 1 {
		errorhandling.GoErrorHandling()
		} else {
		//04-types
		types.GoRangeIterators()
		types.GoGenerics()
		types.GoEmbedding()
		types.GoEnums()
		types.GoInterface()
		types.GoMethod()
		types.GoPointer()
		types.GoStruct()
		//03-functions
		functions.GoFunctions()
		functions.GoClosureFunctions()
		functions.GoMultipleReturn()
		functions.GoVaradicFuntions()
		functions.GoRecursiveFunctions()
		//02-collections
		collections.GoArray()	
		collections.GoSlices()	
		collections.GoMaps()	
		collections.GoRange()	
		//01-fundamental
		fundamentals.GoString()
		fundamentals.GoSwitch()
		fundamentals.GoIfelse()
		fundamentals.GoLooping()
		fundamentals.GoVariables()	
		fundamentals.GoConstants()
		fundamentals.GoValues()
	}

}
