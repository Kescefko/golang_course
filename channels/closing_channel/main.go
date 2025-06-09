package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	ch <- 10
	ch <- 11

	val, ok := <-ch
	fmt.Println(val, ok)
	close(ch)

	// Panic: close already closed channel, sending to a closed channel
	// close(ch)
	// ch <- 11
}
