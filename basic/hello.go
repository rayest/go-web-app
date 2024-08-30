package basic

import "fmt"

var (
	campusName string
	campusCity string
)

func getUserInfo() {
	var name = "rayest"
	fmt.Println("name: " + name)

	city := "上海"
	fmt.Println("city: " + city)

	age, job := "28", "Developer"
	fmt.Println("age: " + age)
	fmt.Println("job: " + job)

	campusCity = "南京"
	campusName = "南京理工大学"
	fmt.Println("campus name: " + campusName)
	fmt.Println("campus city: " + campusCity)
}

func hello() {
	fmt.Println("Hello, World!")
	getUserInfo()

}
