package goGeek

import "fmt"

// 判断是否是切片类型 []string
var container = []string{"zero", "one", "two"}

func demo11() {
	// 类型断言表达式的语法形式是x.(T)
	// x 代表要被判断类型的值
	value, isSlice := interface{}(container).([]string)
	fmt.Println(value)
	fmt.Println("Type is slice: ", isSlice)

	fmt.Println("===========")

	// 字典类型
	container := map[string]string{"name": "Kobe", "age": "40", "job": "player"}
	fmt.Println("Name is:", container["name"])
	value2, isMap := interface{}(container).(map[string]string)
	fmt.Println(value2)
	fmt.Println("Type is map:", isMap)
}
