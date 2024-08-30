package basic

import "fmt"

// 引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。
func swap(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

func funcs() {
	/* 定义局部变量 */
	var a = 100
	var b = 200

	fmt.Printf("交换前，a 的值 : %d\n", a )
	fmt.Printf("交换前，b 的值 : %d\n", b )

	/* 调用 swap() 函数
	* &a 指向 a 指针，a 变量的地址
	* &b 指向 b 指针，b 变量的地址
	*/
	swap(&a, &b)

	fmt.Printf("交换后，a 的值 : %d\n", a )
	fmt.Printf("交换后，b 的值 : %d\n", b )
}