package main

import "fmt"

func main() {
	// nums := [5]int{1, 2, 3, 4, 5}

	// for i, item := range nums {
	// 	fmt.Println(i, item)
	// }

	ch := make(chan int, 5)
	ch <- 100
	ch <- 200
	close(ch)

	for val := range ch {
		fmt.Println(val)
	}
}
