package fundamentals

import "fmt"

func GoVariables() {
	var a = "umur"
	var b, c = "teguh", 21
	fmt.Println(a,b,c)

	var d  =true
	fmt.Println(d)
	var o int
	fmt.Println(o,"  <--- apa bila variable int tidak diisi dan ditampilkan nilainya akan otomatis akan nol bukan null") 
}