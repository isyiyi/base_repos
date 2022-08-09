package main

import "fmt"

func main() {
	var name string = "golang"
	fmt.Println(name)

	var age int
	age = 17
	fmt.Println(age)

	var (
		address = "重庆"
		school  = "重庆邮电大学"
	)
	fmt.Println(address)
	fmt.Println(school)

	var (
		lab     string
		teacher string
	)
	lab = "计算智能重庆市重点实验室"
	teacher = "周应华"
	fmt.Println(lab)
	fmt.Println(teacher)

	var id, teacher2 = "s210231175", "尚明生"
	fmt.Println(id)
	fmt.Println(teacher2)

	masterTeacher := "史晓雨"
	fmt.Println(masterTeacher)
}
