// Jika Timer ibarat "Timer Masak" yang berbunyi
// sekali saat waktu habis, maka Ticker ibarat "Metronom" atau "Detak Jantung"
// yang berbunyi berulang kali
// pada interval waktu tertentu.

package concurrency

import (
	"fmt"
	"time"
)

func RunTickers() {

	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(5600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
