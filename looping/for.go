package looping

import (
	"fmt"
)

func GoFor() {
	// hanya dengan satu kondisi	
	i := 3
	for i <= 3 {
		fmt.Println(i)
		i++
	}
	// biasa

	for j:=0; j<3; j++{
		fmt.Print(j)
		if j!=2 {
			fmt.Print("-")
		}
	}
	println()

	//with func range ket = range bisa auto increment variable
	 for i := range 3 {
        fmt.Println("range", i)
    }

	// bisa tanpa syarat tetapi akan terus berulang tanpa henti
	for {
		fmt.Println("loooop")
		//di perlukan break untuk memberhentikan unlimited looping
		break
	}

	for n :=  range 6{
		//untuk skip bilangan genp
		if n % 2 == 0{
			continue
		}
		fmt.Println(n)
	}


}