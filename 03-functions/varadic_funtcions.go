package functions

import "fmt"

func sum(nums ...int) {
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println("Total : ", total)
}
func RunVaradicFuntions() {
	// seperti fmt.Println yang bisa memanggil
	// berapa pun data yang ada didalam tumpukan data
	// jika tidak spesifik
	sum(1)
	sum(1, 2, 3)
	sum(1, 2, 3, 4, 5, 6, 7, 8)
}
