package goGeek

import "fmt"

func demo15() {
	// 切片
	s1 := make([]int, 5)
	fmt.Println("Value of s1: ", s1)
	fmt.Println("Length: ", len(s1))
	fmt.Println("Capacity: ", cap(s1))

	fmt.Println("==============================")

	s2 := make([]int, 5, 8)
	fmt.Println("Value of s2: ", s2)
	fmt.Println("Length: ", len(s2))
	fmt.Println("Capacity: ", cap(s2))
}
