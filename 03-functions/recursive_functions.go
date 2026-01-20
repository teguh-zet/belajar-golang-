package functions

import "fmt"

// pada function ini fact akan manggil terus dirinya sampai
// tercapai kondisi n == 0
// akan menampilkan hasil dari 7*6*5*4*3*2*1
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func RunRecursiveFunctions() {
	fmt.Println(fact(7))
	fmt.Println(fact(2))
}
