package fundamentals

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
func GoIfelse() {

	fmt.Println("if - else")
	// biasa
	if 7%2 == 0 {
		fmt.Println("7 adalah ganjil")
	}else{
		fmt.Println("7 adalah gentap")
	}

	// digolang juga bisa if tanpa else

	if 8%4 == 0 {
        fmt.Println("8 is divisible by 4")
    }

	// menggunakan operator logika 
	if 8%2 ==0 || 7%2 ==0 {
		fmt.Println("7 dan 8 adalah bilangan genap")
	}else{
		fmt.Println("7 adalah ganjil, sedangkan 8 adalah genap")
	}
	// deklarasi variable didalam condisi 
	if num:= 9;num <0{
		fmt.Println(num, " adalah negatif")
	}else if num < 10{
		fmt.Println(num," bilangan 1 digit")
	}else {
		fmt.Println(num, " lebih dari 1 digit")
	} 

}