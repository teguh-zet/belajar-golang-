package concurrency

import "fmt"

func RunChannels() {
	messages := make(chan string)
	go func() {
		messages <- "ping"
	}()
	msg := <-messages
	fmt.Println(msg)
}

// channel bisa ditampung
func RunChannelBuffered() {
	messages := make(chan string, 2)
	messages <- "buffered"
	messages <- "message"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
