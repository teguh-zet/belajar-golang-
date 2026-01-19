package decisions

import (
	"fmt"
	"time"


)

func GoSwitch() {
	i := 2
	fmt.Println("write ", i, " as")
	switch i{
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
	
	switch time.Now().Weekday(){
	case time.Saturday, time.Sunday:
		fmt.Println("weekend")
	default :
		fmt.Println("weekday")
	}
	// switch case tanpa ekspresi
	t:=time.Now()
	switch  {
	case t.Hour() < 12:
		fmt.Println("sebelum siang")
	default:
		fmt.Println("sesudah siang")
	}
	// menggunalkan variable sebagai kondisi
	whatAmI := func ( i interface{}){
		switch t:= i.(type){
		case bool : fmt.Println("i adalah boolean")
		case int : fmt.Println("i adalah integer")
		default : fmt.Printf("i adalah %T \n", t)
		}
	}
	whatAmI(true)
	whatAmI("teguhhh")
	whatAmI(0.4354534)
	whatAmI(3e10)
	
}