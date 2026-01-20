package collections

import (
	"fmt"
	"slices"
)

func RunSlices() {
	var s []string
	fmt.Println("s = ", s, s == nil, len(s) == 0)
	// len() untuk mendapat panjang slice dan cap() adalah untuk mendapatkan batas maksimum slice
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	//
	fmt.Println("len = ", len(s))

	// fungsi append akan menambah data yang akan diletak diindeks paling belakang tidak peduli data itu ksong atau tidak
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)
	//hasil nya akan berubah dikarenakan telah menambah d e f
	fmt.Println("len = ", len(s))
	//fungsi copy untuk slice. ket copy berfungsi untuk membuat slice kosong dengan panjang yang sama dengan yang disalin
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy : ", c)
	// slice dengan sintaks (low : hight)
	l := s[2:5]
	// 2 : 5 artinya dari index 2 sampai <5
	fmt.Println("sl1:", l)
	//menampilkan isi slice <5
	l = s[:5]
	fmt.Println("sl2:", l)
	// termasuk 2
	l = s[2:]
	fmt.Println("sl3:", l)
	// declar dan inisialisasi slice dalam satu baris
	t := []string{"h", "i", "j", "k"}
	fmt.Println("dcl", t)

	t2 := []string{"g", "h", "i"}
	// menggunakan fungsi Equal dari package "slices"
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	} else {
		fmt.Println("t != t2")
	}

	// slice multi dimensi, tidak seperti array jadi panjang slice bisa bervariasi
	twoD := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = j + i
		}

	}
	fmt.Println("2d: ", twoD)

}
