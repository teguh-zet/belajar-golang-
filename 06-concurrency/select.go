package concurrency

import (
	"fmt"
	"time"
)

// Switch: Tidak memblokir eksekusi. switch akan segera mengevaluasi nilai dan menjalankan case yang cocok.
// Select: Bersifat blocking. select akan menunggu (memblokir) goroutine hingga salah satu case (operasi channel) siap dieksekusi.
func RunSelect() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for range 2 {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
