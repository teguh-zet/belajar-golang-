package decisions

import "fmt"

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