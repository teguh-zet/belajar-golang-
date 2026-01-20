package errorhandling

import "fmt"

func mayPanic() {
    panic("a problem")
}

func RunRecoverExample() {

    defer func() {
        if r := recover(); r != nil {

            fmt.Println("Recovered. Error:\n", r)
			

        }
    }()

    mayPanic()
	
    fmt.Println("After mayPanic()")
}