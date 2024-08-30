package basic

import "fmt"

type Film struct {
	id       int
	name     string
	director string
	actor    string
	actress  string
}

func book() {
	var film1 = Film{1, "阿飞正传", "王家卫", "张国荣", "张曼玉"}
	fmt.Println("Film1 name: " + film1.name)

	var film2 = Film{name: "2046", director: "王家卫", actor: "梁朝伟", actress: "张曼玉"}
	fmt.Println("Film2 name: " + film2.name)
}
