package basic

import "fmt"

func maps() {
	var man = make(map[string]string)

	man["zhao"] = "赵一"
	man["qian"] = "钱二"
	man["sun"] = "孙三"
	man["lee"] = "李四"

	for firstName := range man {
		fmt.Println(man[firstName], " 姓 ", firstName)
	}
}
