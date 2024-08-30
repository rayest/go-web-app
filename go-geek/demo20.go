package goGeek

import "fmt"

// 通道
func demo20() {
	// 声明通道。容量为 3
	ch1 := make(chan int, 3)

	ch1 <- 2
	ch1 <- 1
	ch1 <- 3

	elem1 := <- ch1
	elem2 := <- ch1
	elem3 := <- ch1

	fmt.Println(elem1)
	fmt.Println(elem2)
	fmt.Println(elem3)
}
