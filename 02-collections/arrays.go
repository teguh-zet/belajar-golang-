package collections

import "fmt"

func RunArray() {
	fmt.Println("contoh array")
	var a [6]int
	fmt.Println("array kosong : ", a)
	a[4] = 100
	
	fmt.Println("output array a tanpa index :", a, "<-- aakan tanpil keseluruhan")
	fmt.Println("output array a index 4 :", a[4])

	fmt.Println("mengetahui panjang array dengan fungsi len : ", len(a))
	// deklarasi sekaligus inisialisasi array
	b := [6]int{1, 2, 3, 4, 5, 11}
	fmt.Println("isi dari array ", b)
	b = [...]int{1, 2, 3, 4, 10, 11}
	fmt.Println("setelah dideklar ulang", b)
	b = [...]int{100, 4: 400, 500}
	fmt.Println("idx:", b)

	//array 2 dimensi
	var twoD [3][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d : ", twoD)
	// inisialisasi ulang array 2d
	twoD = [3][3]int{
		{1, 2, 3},
		{3, 2, 1},
		{3, 2, 1},
	}
	fmt.Println("2d : ", twoD)

}
