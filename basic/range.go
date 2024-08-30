package basic

import "fmt"

func ranges() {
	// 当不需要使用该元素的序号时，所以使用空白符"_"省略之
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("After sum: ", sum)

	// 需要使用时
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index: ", i)
		}
	}
}
