package functions

import "fmt"

// <-- jumlah nilai pembaliknya ditentikan oleh jumlah
// variable dalam argumen kedua
//sedang kan argumen pertama untuk menangkap argumen pemanggil
func values() (int, int, int) {
	return 3, 4, 5
}
func RunMultipleReturn() {
	fmt.Println("multiple func")
	a, b, c := values()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	_, d, _ := values()
	fmt.Println(d)

}
