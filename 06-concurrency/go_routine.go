package concurrency

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := range 3 {
		fmt.Println(from, ":", i)
	}
}
func RunRoutine() {
	//pemanggilan biasa
	f("direct")
	// pemanggilan go routine dengan keyword go

	//hasilnya tergantung yang pertama selesai
	go f("goroutine")

	go func(msg string, dd int) {
		fmt.Println(msg)
	}("asdasdasdasdadasdasdasdadsadasdas", 2)

	go func(msg string) {
		fmt.Println(msg)
	}("going")
	// Menggunakan time.Sleep untuk menunggu goroutine selesai adalah cara yang buruk di aplikasi nyata (karena kita tidak tahu pasti berapa lama prosesnya berjalan).
	time.Sleep(time.Second)
	fmt.Println("done")

}
