package main

import (
	fundamentals "first-go/01-fundamentals"
	collections "first-go/02-collections"
	functions "first-go/03-functions"
	types "first-go/04-types"
	errorhandling "first-go/05-error-handling"
	concurrency "first-go/06-concurrency"
	system "first-go/07-systems"
)

func main() {
	i := 1
	if i == 1 {
		system.RunJSONExample()
	} else {
		//system
		system.RunFileExample()
		//06-concurrency
		concurrency.RunWaitGroup()
		concurrency.RunWorkerPool()
		concurrency.RunTickers()
		concurrency.RunTimers()
		concurrency.RunRangeOverChannel()
		concurrency.RunClosingChannel()
		concurrency.RunTimeouts()
		concurrency.RunChannelNonBlocking()
		
		concurrency.RunSelect()
		concurrency.RunChannelDirection()
		concurrency.RunChannelSynchronization()
		concurrency.RunChannelBuffered()
		concurrency.RunChannels()
		concurrency.RunRoutine()
		//05-error-handling
		errorhandling.RunRecoverExample()
		errorhandling.RunDeferExample()
		errorhandling.RunPanicExample()
		errorhandling.RunCostumeError()
		errorhandling.RunErrorHandling()
		//04-types
		types.RunRangeIterators()
		types.RunGenerics()
		types.RunEmbedding()
		types.RunEnums()
		types.RunInterface()
		types.RunMethod()
		types.RunPointer()
		types.RunStruct()
		//03-functions
		functions.RunFunctions()
		functions.RunClosureFunctions()
		functions.RunMultipleReturn()
		functions.RunVaradicFuntions()
		functions.RunRecursiveFunctions()
		//02-collections
		collections.RunSortFuncExample()
		collections.RunSortingExample()
		collections.RunArray()
		collections.RunSlices()
		collections.RunMaps()
		collections.RunRange()
		//01-fundamental
		fundamentals.RunString()
		fundamentals.RunSwitch()
		fundamentals.RunIfelse()
		fundamentals.RunLooping()
		fundamentals.RunVariables()
		fundamentals.RunConstants()
		fundamentals.RunValues()
	}

}
