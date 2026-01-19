package constants

import (
	"fmt"
	"math"
)

const s string = "constant"

func GoConstants() {
	fmt.Println(s)
	 
	const n = 500000000
	const d = 3e20/n
	
	fmt.Println(d)
	fmt.Println(int64(d)) 
	
	
	fmt.Println(math.Sin(n))
	// catatan untuk variable biasa atau di deklar dengan var
	// tidak bisa mengubah tipe data secara otomatis seperti const
	// contoh fmt.Println(math.Sin(n)) <- n yang terlihat int sebelumnya bisa di gunakan di func sin yang mempunyai argument float
	// bentuk math.sin adalah = func Sin(x float64) float64

}